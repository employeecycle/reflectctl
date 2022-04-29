package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/employeecycle/reflectctl/sdk/reflect"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// suitesExecuteCmd represents the suitesExecute command
var suitesExecuteCmd = &cobra.Command{
	Use:   "execute",
	Short: "Execute a test suite",
	Long: `Execute a test suite

	reflectctl suites execute

This command accepts overrides in .reflectctl.yaml like this:

	executeSuiteOptions:
		overrides:
			cookies:
			- name: someCookie
				value: someCookieValue
				maxAge: 123
	
	`,
	RunE: func(cmd *cobra.Command, args []string) error {
		var id string

		if len(args) != 0 {
			id = args[0]
		} else {
			id = "all"
		}

		apiKey := viper.GetViper().GetString("key")

		r := reflect.NewReflect(&reflect.NewReflectInput{
			APIKey: apiKey,
		})

		var executeSuiteOptions *reflect.ExecuteSuiteOptions

		if viper.IsSet("executeSuiteOptions") {
			executeSuiteOptions = &reflect.ExecuteSuiteOptions{}
			err := viper.UnmarshalKey("executeSuiteOptions", executeSuiteOptions)
			if err != nil {
				return fmt.Errorf("suitesExecuteCmd unmarshal options: %w", err)
			}
		}

		output, err := r.ExecuteSuite(id, executeSuiteOptions)

		if err != nil {
			return fmt.Errorf("suitesExecuteCmd: %w", err)
		}

		jsonFormat, err := cmd.Flags().GetBool("json")

		if err != nil {
			return fmt.Errorf("suitesExecuteCmd: %w", err)
		}

		if jsonFormat {
			jsonOutput, err := json.Marshal(output)

			if err != nil {
				return nil
			}

			fmt.Println(string(jsonOutput))
			return nil
		}

		fmt.Println(output.ExecutionID)

		return nil
	},
}

func init() {
	suitesCmd.AddCommand(suitesExecuteCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// suitesExecuteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// suitesExecuteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
