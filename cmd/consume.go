/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"fmt"
	"io"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"telicent.io/federation-cli/pkg/api/v1alpha"
)

// consumeCmd represents the consume command
var consumeCmd = &cobra.Command{
	Use:   "consume",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: runConsume,
}

func init() {
	rootCmd.AddCommand(consumeCmd)
	consumeCmd.Flags().String(f_topic, "", "the topic to consume")
	consumeCmd.Flags().Int64(f_offset, 0, "the offset to consume from")
}

func runConsume(cmd *cobra.Command, args []string) {

	err := config.Validate()
	handleErr(err)

	conn, err := getClientConn(config)
	handleErr(err)
	defer conn.Close()

	client := v1alpha.NewFederatorServiceClient(conn)

	tr, err := topicRequest(config, cmd.Flags())
	handleErr(err)

	stream, err := client.GetKafkaConsumer(context.Background(), tr)
	handleErr(err)

	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}

		prettyPrint(msg)
	}
}

func topicRequest(c commonConfig, ff *pflag.FlagSet) (*v1alpha.TopicRequest, error) {
	tr := &v1alpha.TopicRequest{
		APIKey: c.apiKey,
		Client: c.clientId,
	}

	topic, err := ff.GetString(f_topic)
	handleErr(err)

	offset, err := ff.GetInt64(f_offset)
	handleErr(err)

	tr.Topic = topic
	tr.Offset = offset

	return tr, nil
}

func prettyPrint(msg *v1alpha.KafkaByteBatch) {

	fmt.Println("-----BEGIN MESSAGE-----")
	fmt.Printf("Message: offset(%d)\n", msg.Offset)
	fmt.Println("Headers:")
	for _, h := range msg.Shared {
		fmt.Printf("\t%s: %s\n", h.Key, h.Value)
	}
	fmt.Println("Body:")
	fmt.Println()
	fmt.Println(string(msg.Value))
	fmt.Println()
	fmt.Println("-----END MESSAGE-----")
	fmt.Println()
}
