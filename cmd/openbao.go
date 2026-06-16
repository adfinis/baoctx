package cmd

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/adfinis/baoctx/pkg/targetdir"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var (
	openbaoEndpoint         string
	openbaoToken            string
	openbaoCaPath           string
	openbaoCaCert           string
	openbaoCert             string
	openbaoKey              string
	openbaoFormat           string
	openbaoNamespace        string
	openbaoSkipVerify       string
	openbaoClientTimeout    string
	openbaoClusterAddr      string
	openbaoLogLevel         string
	openbaoMaxRetries       string
	openbaoRedirectAddr     string
	openbaoTlsServerName    string
	openbaoCliNoColour      string
	openbaoRateLimit        string
	openbaoSvrLookup        string
	openbaoMfa              string
	openbaoHttpProxy        string
	openbaoDisableRedirects string
	openbaoAuthMethod       string
	openbaoOidcCallbackHost string
	openbaoOidcListenAddr   string
	openbaoOidcRole         string
	openbaoOidcCallbackMode string
)

var openbaoCmd = &cobra.Command{
	Use:   "openbao",
	Short: "Manage OpenBao context profiles ",
	Long:  `Manage OpenBao context profiles.`,
	ValidArgs: []string{
		"create",
		"delete",
		"list",
		"select",
		"update",
		"set-default",
	},
	//Args:                  cobra.OnlyValidArgs,
	DisableFlagsInUseLine: true,
}

var openbaoSetDefaultCmd = &cobra.Command{
	Use:                   "set-default",
	Short:                 "set a default context profile for OpenBao ",
	Long:                  `set a default context profile for OpenBao.`,
	DisableFlagsInUseLine: true,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return errors.New("requires a name argument")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {

		if c.OpenBao[args[0]] == nil {
			log.Fatalf("Profile %s not found", args[0])
		}

		context := c.OpenBao[args[0]]

		var exportCommandStr []string

		var shellCommandEndpoint string
		endpoint := context.Endpoint
		if endpoint != "" {
			shellCommandEndpoint = fmt.Sprintf("export BAO_ADDR=%s", endpoint)
			exportCommandStr = append(exportCommandStr, shellCommandEndpoint)
		}

		var shellCommandToken string
		token := context.Token
		if token != "" && context.AuthMethod != "oidc" {
			shellCommandToken = fmt.Sprintf("export BAO_TOKEN=%s", token)
			exportCommandStr = append(exportCommandStr, shellCommandToken)
		}

		var shellCommandCaCert string
		caCert := context.CaCert
		if caCert != "" {
			shellCommandCaCert = fmt.Sprintf("export BAO_CACERT=%s", caCert)
			exportCommandStr = append(exportCommandStr, shellCommandCaCert)
		}

		var shellCommandCert string
		cert := context.Cert
		if cert != "" {
			shellCommandCert = fmt.Sprintf("export BAO_CLIENT_CERT=%s", cert)
			exportCommandStr = append(exportCommandStr, shellCommandCert)
		}

		var shellCommandCaPath string
		caPath := context.CaPath
		if caPath != "" {
			shellCommandCaPath = fmt.Sprintf("export BAO_CAPATH=%s", caPath)
			exportCommandStr = append(exportCommandStr, shellCommandCaPath)
		}

		var shellCommandKey string
		key := context.Key
		if key != "" {
			shellCommandKey = fmt.Sprintf("export BAO_CLIENT_KEY=%s", key)
			exportCommandStr = append(exportCommandStr, shellCommandKey)
		}

		var shellCommandNameSpace string
		namespace := context.Namespace
		if namespace != "" {
			shellCommandNameSpace = fmt.Sprintf("export BAO_NAMESPACE=%s", namespace)
			exportCommandStr = append(exportCommandStr, shellCommandNameSpace)
		}

		var shellCommandSkipVerify string
		skipVerify := context.SkipVerify
		if skipVerify != "" {
			shellCommandSkipVerify = fmt.Sprintf("export BAO_SKIP_VERIFY=%s", skipVerify)
			exportCommandStr = append(exportCommandStr, shellCommandSkipVerify)
		}

		var shellClientTimeout string
		timeout := context.ClientTimeout
		if timeout != "" {
			shellClientTimeout = fmt.Sprintf("export BAO_CLIENT_TIMEOUT=%s", timeout)
			exportCommandStr = append(exportCommandStr, shellClientTimeout)
		}

		var shellClusterAddr string
		clusterAddr := context.ClusterAddr
		if clusterAddr != "" {
			shellClusterAddr = fmt.Sprintf("export BAO_CLUSTER_ADDR=%s", clusterAddr)
			exportCommandStr = append(exportCommandStr, shellClusterAddr)
		}

		var shellCommandLogLevel string
		logLevel := context.LogLevel
		if logLevel != "" {
			shellCommandLogLevel = fmt.Sprintf("export BAO_LOG_LEVEL=%s", logLevel)
			exportCommandStr = append(exportCommandStr, shellCommandLogLevel)
		}

		var shellCommandMaxRetries string
		maxRetries := context.MaxRetries
		if maxRetries != "" {
			shellCommandMaxRetries = fmt.Sprintf("export BAO_MAX_RETRIES=%s", maxRetries)
			exportCommandStr = append(exportCommandStr, shellCommandMaxRetries)
		}

		var shellCommandRedirectAddr string
		redirectAddr := context.RedirectAddr
		if redirectAddr != "" {
			shellCommandRedirectAddr = fmt.Sprintf("export BAO_REDIRECT_ADDR=%s", redirectAddr)
			exportCommandStr = append(exportCommandStr, shellCommandRedirectAddr)
		}

		var shellCommandServerName string
		serverName := context.TlsServerName
		if serverName != "" {
			shellCommandServerName = fmt.Sprintf("export BAO_TLS_SERVER_NAME=%s", serverName)
			exportCommandStr = append(exportCommandStr, shellCommandServerName)
		}

		var shellCommandCliNoColour string
		cliNoColour := context.CliNoColour
		if cliNoColour != "" {
			shellCommandCliNoColour = fmt.Sprintf("export BAO_CLI_NO_COLOR=%s", cliNoColour)
			exportCommandStr = append(exportCommandStr, shellCommandCliNoColour)
		}

		var shellCommandRateLimit string
		rateLimit := context.RateLimit
		if rateLimit != "" {
			shellCommandRateLimit = fmt.Sprintf("export BAO_RATE_LIMIT=%s", rateLimit)
			exportCommandStr = append(exportCommandStr, shellCommandRateLimit)
		}

		var shellCommandSvrLookup string
		svrLookup := context.SvrLookup
		if svrLookup != "" {
			shellCommandSvrLookup = fmt.Sprintf("export BAO_SRV_LOOKUP=%s", svrLookup)
			exportCommandStr = append(exportCommandStr, shellCommandSvrLookup)
		}

		var shellCommandMfa string
		mfa := context.Mfa
		if mfa != "" {
			shellCommandMfa = fmt.Sprintf("export BAO_MFA=%s", mfa)
			exportCommandStr = append(exportCommandStr, shellCommandMfa)
		}

		var shellCommandHttpProxy string
		httpProxy := context.HttpProxy
		if httpProxy != "" {
			shellCommandHttpProxy = fmt.Sprintf("export BAO_HTTP_PROXY=%s", httpProxy)
			exportCommandStr = append(exportCommandStr, shellCommandHttpProxy)
		}

		var shellCommandDisableRedirects string
		disableRedirects := context.DisableRedirects
		if disableRedirects != "" {
			shellCommandDisableRedirects = fmt.Sprintf("export BAO_DISABLE_REDIRECTS=%s", disableRedirects)
			exportCommandStr = append(exportCommandStr, shellCommandDisableRedirects)
		}

		commandStr := strings.Join(exportCommandStr, "; ")

		defaultScript := `
#!/bin/bash
` + commandStr

		absolutePath := targetdir.TargetHome() + "/defaults/openbao.sh"

		file, err := os.OpenFile(absolutePath, os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatal(err)
		}
		defer func(file *os.File) {
			err := file.Close()
			if err != nil {
				log.Fatal(err)
			}
		}(file)
		_, err = file.WriteString(defaultScript)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("default profile set")
	},
}

func init() {
	openbaoCmd.AddCommand(openbaoCreateCmd)
	openbaoCmd.AddCommand(deleteOpenbaoCmd)
	openbaoCmd.AddCommand(openbaoSetDefaultCmd)
	openbaoCmd.AddCommand(selectOpenbaoCmd)
	openbaoCmd.AddCommand(openbaoUpdateCmd)
	openbaoCmd.AddCommand(listOpenbaoCmd)

}
