/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"telicent.io/federation-cli/pkg/api/v1alpha"
)

// topicsCmd represents the topics command
var topicsCmd = &cobra.Command{
	Use:   "topics",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: runTopics,
}

func init() {
	rootCmd.AddCommand(topicsCmd)
}

func runTopics(cmd *cobra.Command, args []string) {

	err := config.Validate()
	handleErr(err)

	conn, err := getClientConn(config)
	handleErr(err)
	defer conn.Close()

	client := v1alpha.NewFederatorServiceClient(conn)

	api := &v1alpha.API{
		Key:    config.apiKey,
		Client: config.clientId,
	}

	res, err := client.GetKafkaTopics(context.Background(), api)
	handleErr(err)

	fmt.Println(res)
}

func handleErr(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
