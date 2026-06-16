package cmd

import (
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/spf13/cobra"
)

// selectOpenbaoCmd represents the switch command for OpenBao
var selectOpenbaoCmd = &cobra.Command{
	Use:     "select [name]",
	Short:   "select a context profile",
	Long:    `select a context profile to use with the select command.`,
	Example: `baoctx openbao select example`,
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

		if context.Token != "" {
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
	},
}
