// Copied from https://github.com/carolynvs/stingoftheviper/tree/main
package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

const (
	f_insecure = "insecure"
	f_mtls     = "mtls"
	f_server   = "server"
	f_port     = "port"
	f_cert     = "cert"
	f_key      = "key"
	f_cacert   = "cacert"
	f_apiKey   = "apiKey"
	f_clientId = "clientId"
	f_cfgFile  = "cfgFile"
	f_topic    = "topic"
	f_offset   = "offset"
)

const (
	// The name of our config file, without the file extension because viper supports many different config file languages.
	defaultConfigFilename = "federation-cli"

	// The environment variable prefix of all environment variables bound to our command line flags.
	// For example, --number is bound to STING_NUMBER.
	envPrefix = "FED"

	// Replace hyphenated flag names with camelCase in the config file
	replaceHyphenWithCamelCase = true
)

type commonConfig struct {
	insecure bool
	mtls     bool
	server   string
	port     int
	cert     string
	key      string
	cacert   string
	apiKey   string
	clientId string
	cfgFile  string
}

func (c commonConfig) Validate() error {

	if c.insecure && c.mtls {
		return fmt.Errorf("only one of insecure and mtls flags must be set")
	}

	if c.mtls && (c.key == "" || c.cert == "") {
		return fmt.Errorf("key and cert must be supplied for mtls connections")
	}

	if c.port < 1 {
		return fmt.Errorf("port must be greater than 0")
	}

	return nil
}

var config = commonConfig{}

func initializeConfig(cmd *cobra.Command) error {
	v := viper.New()

	if config.cfgFile == "" {
		v.SetConfigName(defaultConfigFilename)
		v.AddConfigPath(".")
		if err := v.ReadInConfig(); err != nil {
			// It's okay if there isn't a config file
			if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
				return err
			}
		}
	} else {
		v.SetConfigFile(config.cfgFile)
		err := v.ReadInConfig()
		if err != nil {
			return err
		}
	}

	// Set the base name of the config file, without the file extension.
	v.SetConfigName(defaultConfigFilename)

	// Set as many paths as you like where viper should look for the
	// config file. We are only looking in the current working directory.
	v.AddConfigPath(".")

	// Attempt to read the config file, gracefully ignoring errors
	// caused by a config file not being found. Return an error
	// if we cannot parse the config file.
	if err := v.ReadInConfig(); err != nil {
		// It's okay if there isn't a config file
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			fmt.Println(err)
			return err
		}
	}

	// When we bind flags to environment variables expect that the
	// environment variables are prefixed, e.g. a flag like --number
	// binds to an environment variable STING_NUMBER. This helps
	// avoid conflicts.
	v.SetEnvPrefix(envPrefix)

	// Environment variables can't have dashes in them, so bind them to their equivalent
	// keys with underscores, e.g. --favorite-color to STING_FAVORITE_COLOR
	v.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))

	// Bind to environment variables
	// Works great for simple config names, but needs help for names
	// like --favorite-color which we fix in the bindFlags function
	v.AutomaticEnv()

	// Bind the current command's flags to viper
	bindFlags(cmd, v)

	return nil
}

// Bind each cobra flag to its associated viper configuration (config file and environment variable)
func bindFlags(cmd *cobra.Command, v *viper.Viper) {
	cmd.Flags().VisitAll(func(f *pflag.Flag) {
		// Determine the naming convention of the flags when represented in the config file
		configName := f.Name
		// If using camelCase in the config file, replace hyphens with a camelCased string.
		// Since viper does case-insensitive comparisons, we don't need to bother fixing the case, and only need to remove the hyphens.
		if replaceHyphenWithCamelCase {
			configName = strings.ReplaceAll(f.Name, "-", "")
		}

		// Apply the viper config value to the flag when the flag is not set and viper has a value
		if !f.Changed && v.IsSet(configName) {
			val := v.Get(configName)
			cmd.Flags().Set(f.Name, fmt.Sprintf("%v", val))
		}
	})
}
