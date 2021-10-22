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
	Short: "View the status of a test execution",
	Long: `View the status of test execution with execution ID 123:

	reflectctl executions status 123

This will print a table with the execution status showing the test ID, status, timestamps, duration and run ID.

Poll for live updates until all tests are complete with the -w flag:

	reflectctl executions status 123 -w
`,
	RunE: func(cmd *cobra.Command, args []string) error {
		var id string

		if IsInputFromPipe() {
			buffer, err := ioutil.ReadAll(os.Stdin)
			if err != nil {
				return fmt.Errorf("executionStatusCmd: %w", err)
			}

			id = strings.TrimSuffix(string(buffer), "\n")
		} else {
			if len(args) > 0 {
				id = args[0]
			}
		}

		if id == "" {
			return errors.New("execution ID is required")
		}

		apiKey := viper.GetViper().GetString("key")

		r := reflect.NewReflect(&reflect.NewReflectInput{
			APIKey: apiKey,
		})

		output, err := r.GetStatus(id)

		if err != nil {
			return fmt.Errorf("executionStatusCmd: %w", err)
		}

		isJSONFormat, err := cmd.Flags().GetBool("json")

		if err != nil {
			return fmt.Errorf("executionStatusCmd: %w", err)
		}

		var format DisplayFormat

		if isJSONFormat {
			format = JSON
		}

		watch, err := cmd.Flags().GetBool("watch")
		if err != nil {
			return fmt.Errorf("executionStatusCmd: %w", err)
		}

		if watch {
			renderWithWatcher(r, id, format)
			return nil
		}

		err = renderOnce(output, format)

		if err != nil {
			return fmt.Errorf("executionStatusCmd: %w", err)
		}

		return nil
	},
}

func renderOnce(output *reflect.GetStatusOutput, format DisplayFormat) error {
	text, err := render(renderParams{output: output, format: format})
	if err != nil {
		return fmt.Errorf("renderOnce: %w", err)
	}

	fmt.Println(text)
	return nil
}

func renderWithWatcher(r *reflect.Reflect, id string, format DisplayFormat) {
	writer := uilive.New()
	writer.Start()
	allComplete := false
	s := spinner.New(spinner.CharSets[4], 100*time.Millisecond)
	s.Start()

	for !allComplete {
		output, _ := r.GetStatus(id)
		text, _ := render(renderParams{output: output, format: format})
		fmt.Fprint(writer, text)
		time.Sleep(3 * time.Second)
		allComplete = AreAllTestsComplete(output)
	}

	s.Stop()
}

type renderParams struct {
	output   *reflect.GetStatusOutput
	format   DisplayFormat
	timezone *time.Location
}

func render(p renderParams) (string, error) {
	result := ""
	var resultErr error

	if p.timezone == nil {
		t := time.Now()
		p.timezone = t.Location()
	}

	switch p.format {
	case JSON:
		jsonOutput, err := json.Marshal(p.output)

		if err != nil {
			resultErr = err
			break
		}

		result = string(jsonOutput)
	default:
		var buf bytes.Buffer
		w := tabwriter.NewWriter(&buf, 0, 0, 2, ' ', 0)

		fmt.Fprintf(w, "\nStatus for execution %v:\n\n", p.output.ExecutionID)

		fmt.Fprintln(w, "Test ID\tStatus\tStarted\tCompleted\tDuration (s)\tRun ID")

		for _, test := range p.output.Tests {
			line := fmt.Sprintf("%v\t%v\t%v\t%v\t%v\t%v", test.TestID, test.Status, DisplayTime(MillisecondsToTime(test.Started).In(p.timezone)), DisplayTime(MillisecondsToTime(test.Completed).In(p.timezone)), DisplayDuration(test.Completed, test.Started), DisplayRunID(test.RunID))
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
