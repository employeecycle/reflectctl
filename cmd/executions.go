package cmd

import (
	"github.com/spf13/cobra"
)

// executionsCmd represents the executions command
var executionsCmd = &cobra.Command{
	Use:   "executions",
	Short: "View executions. See subcommands for usage.",
}

func init() {
	rootCmd.AddCommand(executionsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// executionsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// executionsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
