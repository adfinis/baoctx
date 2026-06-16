package cmd

import (
	"errors"
	"fmt"

	"log"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var openbaoCreateCmd = &cobra.Command{
	Use:     "create [name]",
	Short:   "create command creates a context profile",
	Long:    `create a context profile with the create command.`,
	Example: `baoctx create example --endpoint="https://example-openbao.com:8200" --token="s.giqoewbnmdjalkjk"`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return errors.New("requires a name argument")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {

		if c.OpenBao[args[0]] != nil {
			log.Fatal("profile already exists")

		}
		v := &OpenBao{
			Endpoint:         openbaoEndpoint,
			Token:            openbaoToken,
			CaPath:           openbaoCaPath,
			CaCert:           openbaoCaCert,
			Cert:             openbaoCert,
			Key:              openbaoKey,
			Format:           openbaoFormat,
			Namespace:        openbaoNamespace,
			SkipVerify:       openbaoSkipVerify,
			ClientTimeout:    openbaoClientTimeout,
			ClusterAddr:      openbaoClusterAddr,
			LogLevel:         openbaoLogLevel,
			MaxRetries:       openbaoMaxRetries,
			RedirectAddr:     openbaoRedirectAddr,
			TlsServerName:    openbaoTlsServerName,
			CliNoColour:      openbaoCliNoColour,
			RateLimit:        openbaoRateLimit,
			SvrLookup:        openbaoSvrLookup,
			Mfa:              openbaoMfa,
			HttpProxy:        openbaoHttpProxy,
			DisableRedirects: openbaoDisableRedirects,
			AuthMethod:       openbaoAuthMethod,
			OidcCallbackHost: openbaoOidcCallbackHost,
			OidcListenAddr:   openbaoOidcListenAddr,
			OidcRole:         openbaoOidcRole,
			OidcCallbackMode: openbaoOidcCallbackMode,
		}

		c.OpenBao[args[0]] = v

		viper.Set("openbao", c.OpenBao)
		err := viper.WriteConfig()
		if err != nil {
			return
		}
		fmt.Printf("Created OpenBao profile '%s'\n", args[0])

	},
}

func init() {

	openbaoCreateCmd.PersistentFlags().StringVar(&openbaoEndpoint, "endpoint", "", "set target endpoint details. e.g https://example-openbao.com:8200")
	openbaoCreateCmd.PersistentFlags().StringVar(&openbaoToken, "token", "", "set openbao auth token for this context")
	openbaoCreateCmd.PersistentFlags().StringVar(&openbaoCaPath, "capath", "", "set path to a directory of PEM-encoded CA certificate files on the local disk")
	openbaoCreateCmd.PersistentFlags().StringVar(&openbaoCaCert, "cacert", "", "set path to a PEM-encoded CA certificate file on the local disk")
	openbaoCreateCmd.PersistentFlags().StringVar(&openbaoCert, "cert", "", "set path to a PEM-encoded client certificate on the local disk")
	openbaoCreateCmd.PersistentFlags().StringVar(&openbaoKey, "key", "", "set path to an unencrypted, PEM-encoded private key on disk which corresponds to the matching client certificate")
	openbaoCreateCmd.PersistentFlags().StringVar(&openbaoFormat, "format", "", `set openbao output (read/status/write) in the specified format. Valid formats are "table", "json", or "yaml"`)
	openbaoCreateCmd.PersistentFlags().StringVar(&openbaoNamespace, "namespace", "", "set openbao namespace to use for command")
	openbaoCreateCmd.PersistentFlags().StringVar(&openbaoSkipVerify, "skip-verify", "", "Do not verify OpenBao's presented certificate before communicating with it")
	openbaoCreateCmd.PersistentFlags().StringVar(&openbaoClientTimeout, "client-timeout", "", "Set the Timeout variable")
	openbaoCreateCmd.PersistentFlags().StringVar(&openbaoClusterAddr, "cluster-addr", "", "Set the address that should be used for other cluster members to connect to this node when in High Availability mode")
	openbaoCreateCmd.PersistentFlags().StringVar(&openbaoLogLevel, "log-level", "", "Set the OpenBao server's log level")
	openbaoCreateCmd.PersistentFlags().StringVar(&openbaoMaxRetries, "max-retries", "", "Set the maximum number of retries when certain error codes are encountered")
	openbaoCreateCmd.PersistentFlags().StringVar(&openbaoRedirectAddr, "redirect-addr", "", "Set the address that should be used when clients are redirected to this node when in High Availability mode")
	openbaoCreateCmd.PersistentFlags().StringVar(&openbaoTlsServerName, "tls-server-name", "", "Set the name to use as the SNI host when connecting via TLS")
	openbaoCreateCmd.PersistentFlags().StringVar(&openbaoCliNoColour, "cli-no-colour", "", "If provided, OpenBao output will not include ANSI color escape sequence characters")
	openbaoCreateCmd.PersistentFlags().StringVar(&openbaoRateLimit, "rate-limit", "", "Set the rate at which the openbao command sends requests to OpenBao")
	openbaoCreateCmd.PersistentFlags().StringVar(&openbaoSvrLookup, "svr-lookup", "", "Enables the client to lookup the host through DNS SRV look up")
	openbaoCreateCmd.PersistentFlags().StringVar(&openbaoMfa, "mfa", "", "Set the MFA credentials in the format mfa_method_name[:key[=value]]")
	openbaoCreateCmd.PersistentFlags().StringVar(&openbaoHttpProxy, "http-proxy", "", "Set the HTTP or HTTPS proxy location which should be used by all requests to access OpenBao")
	openbaoCreateCmd.PersistentFlags().StringVar(&openbaoDisableRedirects, "disable-redirects", "", "Prevents the OpenBao client from following redirects")
	openbaoCreateCmd.PersistentFlags().StringVar(&openbaoAuthMethod, "auth-method", "", `Authentication method for this context (e.g. "oidc"). Leave empty for static token.`)
	openbaoCreateCmd.PersistentFlags().StringVar(&openbaoOidcCallbackHost, "oidc-callback-host", "localhost", "Hostname for the OIDC redirect callback")
	openbaoCreateCmd.PersistentFlags().StringVar(&openbaoOidcListenAddr, "oidc-listen-addr", "127.0.0.1", "Listen address for the OIDC redirect callback")
	openbaoCreateCmd.PersistentFlags().StringVar(&openbaoOidcRole, "oidc-role", "", "OIDC role to request")
	openbaoCreateCmd.PersistentFlags().StringVar(&openbaoOidcCallbackMode, "oidc-callback-mode", "", `OIDC callback mode (e.g. "device" for device-flow)`)

	openbaoCreateCmd.MarkPersistentFlagRequired( //nolint:errcheck
		"endpoint",
	)

}
