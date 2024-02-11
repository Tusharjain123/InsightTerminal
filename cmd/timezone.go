package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

var timezoneCmd = &cobra.Command{
	Use:   "timezone",
	Short: "Get the current time in a given timezone",
	Long: `Get the current time in a given timezone.
				  This command takes one argument, the timezone you want to get the current time in.
				  It returns the current time in RFC1123 format.`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		timezone := args[0]
		currentTime, err := getTimeInTimezone(timezone)
		if err != nil {
			log.Fatalln("The timezone string is invalid")
		}
		fmt.Println(currentTime)
	},
}

func init() {
	rootCmd.AddCommand(timezoneCmd)
	timezoneCmd.PersistentFlags().String("date", "", "returns the date in a time zone in a specified format")
}
