package cmd

import (
	"errors"
	"fmt"
	"os"
	"reflect"

	"github.com/adfinis/baoctx/pkg/targetdir"
	"github.com/adrg/xdg"
	"github.com/spf13/cobra"

	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
)

//var cfgFile = ""

var version string = "dev"

// Config struct containing different product profiles
type Config struct {
	OpenBao map[string]*OpenBao `json:"openbao,omitempty" mapstructure:"openbao"`
}

// OpenBao struct with flag parameters
type OpenBao struct {
	Endpoint         string `json:"endpoint,omitempty" mapstructure:"endpoint"`
	Token            string `json:"token,omitempty" mapstructure:"token"`
	CaPath           string `json:"ca_path,omitempty" mapstructure:"ca_path"`
	CaCert           string `json:"ca_cert,omitempty" mapstructure:"ca_cert"`
	Cert             string `json:"cert,omitempty" mapstructure:"cert"`
	Key              string `json:"key,omitempty" mapstructure:"key"`
	Format           string `json:"format,omitempty" mapstructure:"format"`
	Namespace        string `json:"namespace,omitempty" mapstructure:"namespace"`
	SkipVerify       string `json:"skip_verify,omitempty" mapstructure:"skip_verify"`
	ClientTimeout    string `json:"client_timeout,omitempty" mapstructure:"client_timeout"`
	ClusterAddr      string `json:"cluster_addr,omitempty" mapstructure:"cluster_addr"`
	LogLevel         string `json:"log_level,omitempty" mapstructure:"log_level"`
	MaxRetries       string `json:"max_retries,omitempty" mapstructure:"max_retries"`
	RedirectAddr     string `json:"redirect_addr,omitempty" mapstructure:"redirect_addr"`
	TlsServerName    string `json:"tls_server_name,omitempty" mapstructure:"tls_server_name"`
	CliNoColour      string `json:"cli_no_colour,omitempty" mapstructure:"cli_no_colour"`
	RateLimit        string `json:"rate_limit,omitempty" mapstructure:"rate_limit"`
	SvrLookup        string `json:"svr_lookup,omitempty" mapstructure:"svr_lookup"`
	Mfa              string `json:"mfa,omitempty" mapstructure:"mfa"`
	HttpProxy        string `json:"http_proxy,omitempty" mapstructure:"http_proxy"`
	DisableRedirects string `json:"disable_redirects,omitempty" mapstructure:"disable_redirects"`
	AuthMethod       string `json:"auth_method,omitempty" mapstructure:"auth_method"`
	OidcCallbackHost string `json:"oidc_callback_host,omitempty" mapstructure:"oidc_callback_host"`
	OidcListenAddr   string `json:"oidc_listen_addr,omitempty" mapstructure:"oidc_listen_addr"`
	OidcRole         string `json:"oidc_role,omitempty" mapstructure:"oidc_role"`
	OidcCallbackMode string `json:"oidc_callback_mode,omitempty" mapstructure:"oidc_callback_mode"`
}

var (
	c *Config
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "baoctx",
	Short: "A context switcher CLI tool for OpenBao",
	Long: `baoctx allows a user to configure and switch between different contexts for OpenBao by setting tool-specific environment variables.

A context contains connection details for a given target.
Example:
	an openbao-dev context could point to
	https://example-dev-openbao.com:8200 with a token value of s.jidjibndiyuqepjepwo`,
	ValidArgs: []string{
		"create",
		"delete",
		"list",
		"select",
		"update",
		"set-default",
		"config",
		"version",
	},
	Args:    cobra.OnlyValidArgs,
	Version: version,
}

// Root returns the cobra root command.
func Root() *cobra.Command {
	return rootCmd
}

func init() {
	cobra.OnInitialize(initConfig)
	targetdir.TargetHomeCreate()

	rootCmd.AddCommand(openbaoCreateCmd)
	rootCmd.AddCommand(deleteOpenbaoCmd)
	rootCmd.AddCommand(openbaoSetDefaultCmd)
	rootCmd.AddCommand(selectOpenbaoCmd)
	rootCmd.AddCommand(openbaoUpdateCmd)
	rootCmd.AddCommand(listOpenbaoCmd)
	rootCmd.AddCommand(configlCmd)
	rootCmd.AddCommand(versionCmd)

}

// sliceOfMapsToMapHookFunc merges a slice of maps to a map
func sliceOfMapsToMapHookFunc() mapstructure.DecodeHookFunc {
	return func(from reflect.Type, to reflect.Type, data interface{}) (interface{}, error) {
		if from.Kind() == reflect.Slice && from.Elem().Kind() == reflect.Map && (to.Kind() == reflect.Struct || to.Kind() == reflect.Map) {
			source, ok := data.([]map[string]interface{})
			if !ok {
				return data, nil
			}
			if len(source) == 0 {
				return data, nil
			}
			if len(source) == 1 {
				return source[0], nil
			}
			// flatten the slice into one map
			convert := make(map[string]interface{})
			for _, mapItem := range source {
				for key, value := range mapItem {
					convert[key] = value
				}
			}
			return convert, nil
		}
		return data, nil
	}
}

func initConfig() {
	viper.AddConfigPath(xdg.Home)
	viper.AddConfigPath("$HOME/.baoctx")
	viper.SetConfigName("profiles")
	viper.SetConfigType("json")

	// Attempt to read the config file to see if it exists
	if err := viper.ReadInConfig(); err != nil {
		var configFileNotFoundError viper.ConfigFileNotFoundError
		if errors.As(err, &configFileNotFoundError) {
			// Config file not found, use default configuration
			c = &Config{}
		}
	}

	// Config file found and successfully loaded
	configOption := viper.DecodeHook(mapstructure.ComposeDecodeHookFunc(
		sliceOfMapsToMapHookFunc(),
		mapstructure.StringToTimeDurationHookFunc(),
		mapstructure.StringToSliceHookFunc(","),
	))
	if err := viper.Unmarshal(&c, configOption); err != nil {
		fmt.Println("Error unmarshaling config:", err)
		os.Exit(1)
	}

	if c.OpenBao == nil {
		c.OpenBao = map[string]*OpenBao{}
	}

	// Automatically bind environment variables
	viper.AutomaticEnv()

}
