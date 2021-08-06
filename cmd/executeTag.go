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

// executeTagCmd represents the executeTag command
var executeTagCmd = &cobra.Command{
	Use:   "tag",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		var tag string

		if len(args) != 0 {
			tag = args[0]
		} else {
			tag = "all"
		}

		apiKey := viper.GetViper().GetString("key")

		r := reflect.NewReflect(&reflect.NewReflectInput{
			ApiKey: apiKey,
		})

		output, err := r.CreateTagExecution(tag)

		if err != nil {
			return err
		}

		jsonFormat, err := cmd.Flags().GetBool("json")

		if err != nil {
			return err
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
	executeCmd.AddCommand(executeTagCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// executeTagCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// executeTagCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
