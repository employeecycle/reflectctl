package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"text/tabwriter"

	"github.com/employeecycle/reflectctl/sdk/reflect"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// suitesLsCmd represents the suitesLs command
var suitesLsCmd = &cobra.Command{
	Use:   "ls",
	Short: "List test suites",
	Long: `List test suites:
	
	reflectctl suites ls
	`,
	RunE: func(cmd *cobra.Command, args []string) error {
		apiKey := viper.GetViper().GetString("key")

		r := reflect.NewReflect(&reflect.NewReflectInput{
			APIKey: apiKey,
		})

		output, err := r.ListSuites()

		if err != nil {
			return fmt.Errorf("suitesLsCommand: %w", err)
		}

		jsonFormat, err := cmd.Flags().GetBool("json")

		if err != nil {
			return fmt.Errorf("executeTagCmd: %w", err)
		}

		if jsonFormat {
			jsonOutput, err := json.Marshal(output)

			if err != nil {
				return nil
			}

			fmt.Println(string(jsonOutput))
			return nil
		}

		display, err := renderSuites(output.Suites.Data)
		if err != nil {
			return fmt.Errorf("executeTagCmd: %w", err)
		}

		fmt.Println(display)

		return nil
	},
}

func init() {
	suitesCmd.AddCommand(suitesLsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// suitesLsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// suitesLsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func renderSuites(suites []*reflect.SuiteData) (string, error) {
	var buf bytes.Buffer
	w := tabwriter.NewWriter(&buf, 0, 0, 2, ' ', 0)
	fmt.Fprintln(w, "Name\tSuite ID\tcreated")

	for _, suite := range suites {
		line := fmt.Sprintf("%v\t%v\t%v", suite.Name, suite.SuiteID, suite.Created)
		fmt.Fprintln(w, line)
	}

	w.Flush()

	buffResult, err := ioutil.ReadAll(&buf)
	if err != nil {
		return "", fmt.Errorf("renderSuites: %w", err)
	}

	result := string(buffResult)

	return result, nil
}
