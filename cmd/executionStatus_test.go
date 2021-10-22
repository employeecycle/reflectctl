package cmd

import (
	"testing"
	"time"

	"github.com/jasonblanchard/reflectctl/reflect-sdk"
	"github.com/stretchr/testify/assert"
)

func TestRender(t *testing.T) {
	startTime, _ := time.Parse(time.RFC3339, "2021-10-21T19:06:45.0Z")
	endTime, _ := time.Parse(time.RFC3339, "2021-10-21T19:07:45.0Z")
	output := &reflect.GetStatusOutput{
		ExecutionID: 123,
		Tests: []reflect.Test{
			{
				TestID:    111,
				Status:    "complete",
				Started:   int(startTime.Unix()),
				Completed: int(endTime.Unix()),
				RunID:     222,
			},
			{
				TestID:    333,
				Status:    "complete",
				Started:   int(startTime.Unix()),
				Completed: int(endTime.Unix()),
				RunID:     444,
			},
		},
	}

	type input struct {
		output *reflect.GetStatusOutput
		format DisplayFormat
	}

	tests := []struct {
		input    input
		expected string
	}{
		{
			input: input{
				output: output,
				format: Text,
			},
			expected: "\nStatus for execution 123:\n\nTest ID  Status    Started                        Completed                      Duration (s)  Run ID\n111      complete  1970-01-19 17:07:23 -0500 EST  1970-01-19 17:07:23 -0500 EST  0.06          222\n333      complete  1970-01-19 17:07:23 -0500 EST  1970-01-19 17:07:23 -0500 EST  0.06          444\n",
		},
		{
			input: input{
				output: output,
				format: JSON,
			},
			expected: "{\"executionId\":123,\"tests\":[{\"testId\":111,\"status\":\"complete\",\"started\":1634843205,\"completed\":1634843265,\"runId\":222},{\"testId\":333,\"status\":\"complete\",\"started\":1634843205,\"completed\":1634843265,\"runId\":444}]}",
		},
	}

	for _, test := range tests {
		result, err := render(test.input.output, test.input.format)
		assert.Nil(t, err)
		assert.Equal(t, test.expected, result)
	}
}
