package cmd

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/adfinis/baoctx/pkg/targetdir"
	"github.com/adrg/xdg"
	"github.com/spf13/cobra"
)

// profileTokenPath returns the path to the per-profile token cache file.
func profileTokenPath(profileName string) string {
	return filepath.Join(targetdir.TargetHome(), "tokens", profileName)
}

// readProfileToken reads the cached token for the given profile.
func readProfileToken(profileName string) (string, error) {
	data, err := os.ReadFile(profileTokenPath(profileName))
	if os.IsNotExist(err) {
		return "", nil
	}
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(data)), nil
}

// writeProfileToken persists the token to the per-profile cache file (0600).
func writeProfileToken(profileName, token string) error {
	return os.WriteFile(profileTokenPath(profileName), []byte(token+"\n"), 0600)
}

// getOrRefreshOIDCToken validates the per-profile cached token and re-runs the
// OIDC login flow if it is absent or expired.
func getOrRefreshOIDCToken(profileName string, context *OpenBao) (string, error) {
	baseEnv := append(os.Environ(), "BAO_ADDR="+context.Endpoint)

	// Try the cached token first.
	cached, err := readProfileToken(profileName)
	if err != nil {
		return "", fmt.Errorf("reading cached token: %w", err)
	}
	if cached != "" {
		lookupEnv := append(baseEnv, "BAO_TOKEN="+cached)
		lookupCmd := exec.Command("bao", "token", "lookup")
		lookupCmd.Env = lookupEnv
		lookupCmd.Stdout = nil
		lookupCmd.Stderr = nil
		if lookupCmd.Run() == nil {
			return cached, nil
		}
	}

	// Token absent or expired
	fmt.Fprintln(os.Stderr, "OpenBao token missing or expired, starting OIDC login...")

	loginArgs := []string{"login", "-format=json", "-method=oidc"}
	if context.OidcCallbackHost != "" {
		loginArgs = append(loginArgs, "callbackhost="+context.OidcCallbackHost)
	}
	if context.OidcListenAddr != "" {
		loginArgs = append(loginArgs, "listenaddress="+context.OidcListenAddr)
	}
	if context.OidcRole != "" {
		loginArgs = append(loginArgs, "role="+context.OidcRole)
	}
	if context.OidcCallbackMode != "" {
		loginArgs = append(loginArgs, "callbackmode="+context.OidcCallbackMode)
	}

	var stdout bytes.Buffer
	loginCmd := exec.Command("bao", loginArgs...)
	loginCmd.Env = baseEnv
	loginCmd.Stdin = os.Stdin
	loginCmd.Stdout = &stdout
	loginCmd.Stderr = os.Stderr
	if err := loginCmd.Run(); err != nil {
		return "", fmt.Errorf("bao login failed: %w", err)
	}

	var result struct {
		Auth struct {
			ClientToken string `json:"client_token"`
		} `json:"auth"`
	}
	if err := json.Unmarshal(stdout.Bytes(), &result); err != nil {
		return "", fmt.Errorf("parsing bao login output: %w", err)
	}
	token := result.Auth.ClientToken
	if token == "" {
		return "", fmt.Errorf("bao login returned no client_token")
	}

	if err := writeProfileToken(profileName, token); err != nil {
		fmt.Fprintf(os.Stderr, "warning: could not cache token for profile %q: %v\n", profileName, err)
	}
	return token, nil
}

// selectOpenbaoCmd represents the switch command for OpenBao
var selectOpenbaoCmd = &cobra.Command{
	Use:     "select [name]",
	Short:   "select a context profile",
	Long:    `select a context profile to use with the select command.`,
	Example: `baoctx select example`,
	ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		if len(args) != 0 {
			return nil, cobra.ShellCompDirectiveNoFileComp
		}
		var profiles []string
		for name := range c.OpenBao {
			profiles = append(profiles, name)
		}
		return profiles, cobra.ShellCompDirectiveNoFileComp
	},
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return errors.New("requires a name argument")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		profile := args[0]

		if c.OpenBao[args[0]] == nil {
			log.Fatalf("Profile %s not found", profile)
		}

		context := c.OpenBao[args[0]]

		exportCommands := []string{}

		if context.Endpoint != "" {
			exportCommands = append(exportCommands, fmt.Sprintf("export BAO_ADDR=%s", context.Endpoint))
		}

		if context.AuthMethod == "oidc" {
			token, err := getOrRefreshOIDCToken(profile, context)
			if err != nil {
				log.Fatalf("OIDC login failed: %v", err)
			}
			exportCommands = append(exportCommands, fmt.Sprintf("export BAO_TOKEN=%s", token))
		} else if context.Token != "" {
			exportCommands = append(exportCommands, fmt.Sprintf("export BAO_TOKEN=%s", context.Token))
		}

		if context.Namespace != "" {
			exportCommands = append(exportCommands, fmt.Sprintf("export BAO_NAMESPACE=%s", context.Namespace))
		}

		if context.CaCert != "" {
			exportCommands = append(exportCommands, fmt.Sprintf("export BAO_CACERT=%s", context.CaCert))
		}

		if context.Cert != "" {
			exportCommands = append(exportCommands, fmt.Sprintf("export BAO_CLIENT_CERT=%s", context.Cert))
		}

		if context.CaPath != "" {
			exportCommands = append(exportCommands, fmt.Sprintf("export BAO_CAPATH=%s", context.CaPath))
		}

		if context.Key != "" {
			exportCommands = append(exportCommands, fmt.Sprintf("export BAO_CLIENT_KEY=%s", context.Key))
		}

		if context.Format != "" {
			exportCommands = append(exportCommands, fmt.Sprintf("export BAO_FORMAT=%s", context.Format))
		}

		if context.SkipVerify != "" {
			exportCommands = append(exportCommands, fmt.Sprintf("export BAO_SKIP_VERIFY=%s", context.SkipVerify))
		}

		if context.ClientTimeout != "" {
			exportCommands = append(exportCommands, fmt.Sprintf("export BAO_CLIENT_TIMEOUT=%s", context.ClientTimeout))
		}

		if context.ClusterAddr != "" {
			exportCommands = append(exportCommands, fmt.Sprintf("export BAO_CLUSTER_ADDR=%s", context.ClusterAddr))
		}

		if context.LogLevel != "" {
			exportCommands = append(exportCommands, fmt.Sprintf("export BAO_LOG_LEVEL=%s", context.LogLevel))
		}

		if context.MaxRetries != "" {
			exportCommands = append(exportCommands, fmt.Sprintf("export BAO_MAX_RETRIES=%s", context.MaxRetries))
		}

		if context.RedirectAddr != "" {
			exportCommands = append(exportCommands, fmt.Sprintf("export BAO_REDIRECT_ADDR=%s", context.RedirectAddr))
		}

		if context.TlsServerName != "" {
			exportCommands = append(exportCommands, fmt.Sprintf("export BAO_TLS_SERVER_NAME=%s", context.TlsServerName))
		}

		if context.CliNoColour != "" {
			exportCommands = append(exportCommands, fmt.Sprintf("export BAO_CLI_NO_COLOR=%s", context.CliNoColour))
		}

		if context.RateLimit != "" {
			exportCommands = append(exportCommands, fmt.Sprintf("export BAO_RATE_LIMIT=%s", context.RateLimit))
		}

		if context.SvrLookup != "" {
			exportCommands = append(exportCommands, fmt.Sprintf("export BAO_SRV_LOOKUP=%s", context.SvrLookup))
		}

		if context.Mfa != "" {
			exportCommands = append(exportCommands, fmt.Sprintf("export BAO_MFA=%s", context.Mfa))
		}

		if context.HttpProxy != "" {
			exportCommands = append(exportCommands, fmt.Sprintf("export BAO_HTTP_PROXY=%s", context.HttpProxy))
		}

		if context.DisableRedirects != "" {
			exportCommands = append(exportCommands, fmt.Sprintf("export BAO_DISABLE_REDIRECTS=%s", context.DisableRedirects))
		}

		fmt.Println(strings.Join(exportCommands, "; "))

		activeToken := context.Token
		if context.AuthMethod == "oidc" {
			if t, err := readProfileToken(profile); err == nil && t != "" {
				activeToken = t
			}
		}
		if activeToken != "" {
			if context.AuthMethod != "oidc" {
				if err := writeProfileToken(profile, activeToken); err != nil {
					fmt.Fprintf(os.Stderr, "warning: could not cache token for profile %q: %v\n", profile, err)
				}
			}
			tokenFile := filepath.Join(xdg.Home, ".vault-token")
			if err := os.WriteFile(tokenFile, []byte(activeToken), 0600); err != nil {
				fmt.Fprintf(os.Stderr, "warning: could not write ~/.vault-token: %v\n", err)
			}
		}
	},
}
