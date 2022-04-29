package cmd

import (
	"testing"
	"time"

	"github.com/employeecycle/reflectctl/sdk/reflect"
	"github.com/stretchr/testify/assert"
)

func TestDisplayRunId(t *testing.T) {
	tests := []struct {
		input    int
		expected string
	}{
		{input: 123, expected: "123"},
		{input: 0, expected: "-"},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, DisplayRunID(test.input))
	}
}

func TestDisplayTime(t *testing.T) {
	inputTime, _ := time.Parse(time.RFC3339, "2021-10-21T19:06:45.0Z")

	tests := []struct {
		input    time.Time
		expected string
	}{
		{input: time.Time{}, expected: "-"},
		{input: inputTime, expected: "2021-10-21 19:06:45 +0000 UTC"},
	}

	for _, test := range tests {
		actual := DisplayTime(test.input)
		assert.Equal(t, test.expected, actual)
	}
}

func TestMillisecondsToTime(t *testing.T) {
	expectedTime, _ := time.Parse(time.RFC3339, "2021-10-21T19:06:45.0Z")

	tests := []struct {
		input    int
		expected time.Time
	}{
		{input: 0, expected: time.Time{}},
		{input: 1634843205975, expected: expectedTime},
	}

	for _, test := range tests {
		actual := MillisecondsToTime(test.input)
		assert.Equal(t, test.expected.Local().String(), actual.Local().String())
	}
}

func TestDisplayDuration(t *testing.T) {
	type input struct {
		start int
		end   int
	}

	tests := []struct {
		input    input
		expected string
	}{
		{input: input{start: 0, end: 0}, expected: "-"},
		{input: input{start: 80000, end: 100000}, expected: "20"},
	}

	for _, test := range tests {
		actual := DisplayDuration(test.input.end, test.input.start)
		assert.Equal(t, test.expected, actual)
	}
}

func TestAreAllTestsComplete(t *testing.T) {
	tests := []struct {
		input    *reflect.GetStatusOutput
		expected bool
	}{
		{
			input: &reflect.GetStatusOutput{
				Tests: []reflect.Test{
					{Status: "succeeded"},
					{Status: "succeeded"},
					{Status: "succeeded"},
				},
			},
			expected: true,
		},
		{
			input: &reflect.GetStatusOutput{
				Tests: []reflect.Test{
					{Status: "failed"},
					{Status: "failed"},
					{Status: "failed"},
				},
			},
			expected: true,
		},
		{
			input: &reflect.GetStatusOutput{
				Tests: []reflect.Test{
					{Status: "failed"},
					{Status: "failed"},
					{Status: "pending"},
				},
			},
			expected: false,
		},
	}

	for _, test := range tests {
		actual := AreAllTestsComplete(test.input)
		assert.Equal(t, test.expected, actual)
	}
}
