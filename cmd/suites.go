package cmd

import (
	"github.com/spf13/cobra"
)

// suitesCmd represents the suites command
var suitesCmd = &cobra.Command{
	Use:   "suites",
	Short: "Work with test suites",
}

func init() {
	rootCmd.AddCommand(suitesCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// suitesCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// suitesCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
