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
	tz, _ := time.LoadLocation("America/New_York")
	output := &reflect.GetStatusOutput{
		ExecutionID: 123,
		Tests: []reflect.Test{
			{
				TestID:    111,
				Status:    "complete",
				Started:   int(startTime.Unix()) * 1000,
				Completed: int(endTime.Unix()) * 1000,
				RunID:     222,
			},
			{
				TestID:    333,
				Status:    "complete",
				Started:   int(startTime.Unix()) * 1000,
				Completed: int(endTime.Unix()) * 1000,
				RunID:     444,
			},
		},
	}

	tests := []struct {
		input    renderParams
		expected string
	}{
		{
			input: renderParams{
				output:   output,
				format:   Text,
				timezone: tz,
			},
			expected: "\nStatus for execution 123:\n\nTest ID  Status    Started                        Completed                      Duration (s)  Run ID\n111      complete  2021-10-21 15:06:45 -0400 EDT  2021-10-21 15:07:45 -0400 EDT  60            222\n333      complete  2021-10-21 15:06:45 -0400 EDT  2021-10-21 15:07:45 -0400 EDT  60            444\n",
		},
		{
			input: renderParams{
				output:   output,
				format:   JSON,
				timezone: tz,
			},
			expected: "{\"executionId\":123,\"tests\":[{\"testId\":111,\"status\":\"complete\",\"started\":1634843205000,\"completed\":1634843265000,\"runId\":222},{\"testId\":333,\"status\":\"complete\",\"started\":1634843205000,\"completed\":1634843265000,\"runId\":444}]}",
		},
	}

	for _, test := range tests {
		result, err := render(test.input)
		assert.Nil(t, err)
		assert.Equal(t, test.expected, result)
	}
}
