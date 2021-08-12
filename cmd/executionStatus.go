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
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"text/tabwriter"
	"time"

	"github.com/briandowns/spinner"
	"github.com/gosuri/uilive"
	"github.com/jasonblanchard/reflectctl/reflect-sdk"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// executionStatusCmd represents the executionStatus command
var executionStatusCmd = &cobra.Command{
	Use:   "status",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		var id string

		if IsInputFromPipe() {
			buffer, err := ioutil.ReadAll(os.Stdin)
			if err != nil {
				return err
			}

			id = strings.TrimSuffix(string(buffer), "\n")
		} else {
			if len(args) > 0 {
				id = args[0]
			}
		}

		if id == "" {
			return errors.New("Execution ID is required")
		}

		// apiKey, err := cmd.Flags().GetString("key")
		apiKey := viper.GetViper().GetString("key")

		r := reflect.NewReflect(&reflect.NewReflectInput{
			ApiKey: apiKey,
		})

		output, err := r.GetStatus(id)

		if err != nil {
			return err
		}

		jsonFormat, err := cmd.Flags().GetBool("json")

		if err != nil {
			return err
		}

		format := ""

		if jsonFormat {
			format = "json"
		}

		watch, err := cmd.Flags().GetBool("watch")
		if err != nil {
			return err
		}

		if watch {
			writer := uilive.New()
			writer.Start()
			allComplete := false
			s := spinner.New(spinner.CharSets[4], 100*time.Millisecond)
			s.Start()

			for allComplete == false {
				output, _ = r.GetStatus(id)
				text, _ := render(output, format)
				fmt.Fprintf(writer, text)
				time.Sleep(3 * time.Second)
				allComplete = AreAllTestsComplete(output)
			}

			s.Stop()
			return nil
		}

		text, err := render(output, format)

		fmt.Println(text)

		return nil
	},
}

func render(output *reflect.GetStatusOutput, format string) (string, error) {
	result := ""
	var resultErr error

	switch format {
	case "json":
		jsonOutput, err := json.Marshal(output)

		if err != nil {
			resultErr = err
			break
		}

		result = string(jsonOutput)
	default:
		var buf bytes.Buffer
		w := tabwriter.NewWriter(&buf, 0, 0, 2, ' ', 0)

		fmt.Fprintln(w, fmt.Sprintf("\nStatus for execution %v:\n", output.ExecutionID))

		fmt.Fprintln(w, "Test ID\tStatus\tStarted\tCompleted\tDuration (s)\tRun ID")

		for _, test := range output.Tests {
			line := fmt.Sprintf("%v\t%v\t%v\t%v\t%v\t%v", test.TestID, test.Status, DisplayTime(MillisecondsToTime(test.Started)), DisplayTime(MillisecondsToTime(test.Completed)), DisplayDuration(test.Completed, test.Started), DisplayRunId(test.RunID))
			fmt.Fprintln(w, line)
		}

		w.Flush()

		buffResult, err := ioutil.ReadAll(&buf)
		if err != nil {
			resultErr = err
			break
		}

		result = string(buffResult)
	}

	return result, resultErr
}

func init() {
	executionsCmd.AddCommand(executionStatusCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// executionStatusCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// executionStatusCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	executionStatusCmd.Flags().BoolP("watch", "w", false, "Watch live output")
}
