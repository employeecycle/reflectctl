/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/jasonblanchard/reflectctl/reflect-sdk"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// executeCmd represents the execute command
var executeCmd = &cobra.Command{
	Use:   "execute",
	Short: "Execute all tests or tests associated with a tag.",
	Long: `Execute all tests:

	reflectctl execute

Execute all tests tagged with "regression":

	reflectctl execute tag regression

This command returns a test ID which you can use to view the test status:

	reflectctl executions status [test ID]

This command accepts overrides in .reflectctl.yaml like this:

	testExecutionOptions:
		overrides:
			cookies:
			- name: someCookie
				value: someCookieValue
				maxAge: 123
`,
	RunE: func(cmd *cobra.Command, args []string) error {
		apiKey := viper.GetViper().GetString("key")

		r := reflect.NewReflect(&reflect.NewReflectInput{
			ApiKey: apiKey,
		})

		var testExecutionOptions *reflect.TestExecutionOptions

		if viper.IsSet("testExecutionOptions") {
			testExecutionOptions = &reflect.TestExecutionOptions{}
			err := viper.UnmarshalKey("testExecutionOptions", testExecutionOptions)
			if err != nil {
				return fmt.Errorf("executeTagCmd unmarshal options: %w", err)
			}
		}

		output, err := r.CreateTagExecution("all", testExecutionOptions)

		if err != nil {
			return fmt.Errorf("executeTagCmd: %w", err)
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

		fmt.Println(output.ExecutionID)

		return nil
	},
}

func init() {
	rootCmd.AddCommand(executeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// executeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// executeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
