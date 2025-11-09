package cmd

import (
	"fmt"
	"os"
	"ticket/cmd/rabbitmq"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "help",
	Short: "help provides help for any command",
	Long:  `help provides help for any command`,
}

func Execute() {
	rabbitmq.Send()
	//	rabbitmq.Receive()
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
