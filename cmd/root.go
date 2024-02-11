package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "terminal_chat",
	Short: "Your used to application to chat with chat GPT",
	Long:  `Chat with Open AI from terminal need not to switch to browser to chat with GPT-3.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
