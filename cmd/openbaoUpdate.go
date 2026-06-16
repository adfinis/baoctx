package cmd

import (
	"errors"
	"fmt"

	"log"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// updateCmd represents the update command
var openbaoUpdateCmd = &cobra.Command{
	Use:     "update [name]",
	Short:   "Update an existing context",
	Long:    `The update command allows you to modify an existing context.`,
	Example: `baoctx openbao update example --endpoint="https://example2-openbao.com:8200" --token="t.loejwikdjuidfhjdi"`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return errors.New("requires a name argument")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		if c.OpenBao[args[0]] == nil {
			log.Fatal("profile does not exists.  Try using the create command")

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
		}

		c.OpenBao[args[0]] = v
		viper.Set("openbao", c.OpenBao)
		err := viper.WriteConfig()
		if err != nil {
			return
		}

		fmt.Printf("Updated OpenBao profile '%s'\n", args[0])

	},
}

func init() {
	openbaoUpdateCmd.PersistentFlags().StringVar(&openbaoEndpoint, "endpoint", "", "set target endpoint details. e.g https://example-openbao.com:8200")
	openbaoUpdateCmd.PersistentFlags().StringVar(&openbaoToken, "token", "", "set openbao auth token for this context")
	openbaoUpdateCmd.PersistentFlags().StringVar(&openbaoCaPath, "capath", "", "set path to a directory of PEM-encoded CA certificate files on the local disk")
	openbaoUpdateCmd.PersistentFlags().StringVar(&openbaoCaCert, "cacert", "", "set path to a PEM-encoded CA certificate file on the local disk")
	openbaoUpdateCmd.PersistentFlags().StringVar(&openbaoCert, "cert", "", "set path to a PEM-encoded client certificate on the local disk")
	openbaoUpdateCmd.PersistentFlags().StringVar(&openbaoKey, "key", "", "set path to an unencrypted, PEM-encoded private key on disk which corresponds to the matching client certificate")
	openbaoUpdateCmd.PersistentFlags().StringVar(&openbaoFormat, "format", "", `set openbao output (read/status/write) in the specified format. Valid formats are "table", "json", or "yaml"`)
	openbaoUpdateCmd.PersistentFlags().StringVar(&openbaoNamespace, "namespace", "", "set openbao namespace to use for command")

	openbaoUpdateCmd.PersistentFlags().StringVar(&openbaoSkipVerify, "skip-verify", "", "Do not verify OpenBao's presented certificate before communicating with it")
	openbaoUpdateCmd.PersistentFlags().StringVar(&openbaoClientTimeout, "client-timeout", "", "Set the Timeout variable")
	openbaoUpdateCmd.PersistentFlags().StringVar(&openbaoClusterAddr, "cluster-addr", "", "Set the address that should be used for other cluster members to connect to this node when in High Availability mode")
	openbaoUpdateCmd.PersistentFlags().StringVar(&openbaoLogLevel, "log-level", "", "Set the OpenBao server's log level")
	openbaoUpdateCmd.PersistentFlags().StringVar(&openbaoMaxRetries, "max-retries", "", "Set the maximum number of retries when certain error codes are encountered")
	openbaoUpdateCmd.PersistentFlags().StringVar(&openbaoRedirectAddr, "redirect-addr", "", "Set the address that should be used when clients are redirected to this node when in High Availability mode")
	openbaoUpdateCmd.PersistentFlags().StringVar(&openbaoTlsServerName, "tls-server-name", "", "Set the name to use as the SNI host when connecting via TLS")
	openbaoUpdateCmd.PersistentFlags().StringVar(&openbaoCliNoColour, "cli-no-colour", "", "If provided, OpenBao output will not include ANSI color escape sequence characters")
	openbaoUpdateCmd.PersistentFlags().StringVar(&openbaoRateLimit, "rate-limit", "", "Set the rate at which the openbao command sends requests to OpenBao")
	openbaoUpdateCmd.PersistentFlags().StringVar(&openbaoSvrLookup, "svr-lookup", "", "Enables the client to lookup the host through DNS SRV look up")
	openbaoUpdateCmd.PersistentFlags().StringVar(&openbaoMfa, "mfa", "", "Set the MFA credentials in the format mfa_method_name[:key[=value]]")
	openbaoUpdateCmd.PersistentFlags().StringVar(&openbaoHttpProxy, "http-proxy", "", "Set the HTTP or HTTPS proxy location which should be used by all requests to access OpenBao")
	openbaoUpdateCmd.PersistentFlags().StringVar(&openbaoDisableRedirects, "disable-redirects", "", "Prevents the OpenBao client from following redirects")

}
