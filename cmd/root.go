/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "federation-cli",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		// You can bind cobra and viper in a few locations, but PersistencePreRunE on the root command works well
		return initializeConfig(cmd)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().BoolVar(&config.insecure, f_insecure, false, "help for insecure")
	rootCmd.PersistentFlags().BoolVar(&config.mtls, f_mtls, false, "help for mtls")
	rootCmd.PersistentFlags().StringVar(&config.server, f_server, "127.0.0.1", "help for server")
	rootCmd.PersistentFlags().IntVar(&config.port, f_port, 443, "help for port")
	rootCmd.PersistentFlags().StringVar(&config.cert, f_cert, "", "help for cert")
	rootCmd.PersistentFlags().StringVar(&config.key, f_key, "", "help for key")
	rootCmd.PersistentFlags().StringVar(&config.cacert, f_cacert, "", "help for cacert")
	rootCmd.PersistentFlags().StringVar(&config.apiKey, f_apiKey, "", "help for apiKey")
	rootCmd.PersistentFlags().StringVar(&config.clientId, f_clientId, "", "help for clientId")
	rootCmd.PersistentFlags().StringVar(&config.cfgFile, f_cfgFile, "", "help for cfgFile (default is $HOME/.federation-cli.yaml)")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if config.cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(config.cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".federation-cli" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".federation-cli")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
