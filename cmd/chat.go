package cmd

import (
	"github.com/spf13/cobra"
)

var chatCmd = &cobra.Command{
	Use:   "chat",
	Short: "Talk to our AI model",
	Long:  `You can type your message and our AI model will respond to you.`,
	Run: func(cmd *cobra.Command, args []string) {
		message := cmd.Flag("m").Value.String()
		getResponse(message)
	},
}

func init() {
	rootCmd.AddCommand(chatCmd)
	chatCmd.Flags().String("m", "", "The message you want to send to the AI model")
}
