package cmd

import (
	"fmt"
	"time"

	"github.com/jasonblanchard/reflectctl/reflect-sdk"
)

func DisplayRunId(id int) string {
	if id == 0 {
		return "-"
	}

	return fmt.Sprintf("%v", id)
}

func DisplayTime(t time.Time) string {
	// 1969-12-31 19:00:00 -0500 EST
	if t.IsZero() {
		return "-"
	}

	return fmt.Sprintf("%v", t)
}

func MillisecondsToTime(t int) time.Time {
	if t == 0 {
		return time.Time{}
	}
	return time.Unix(int64(t/1000), 0)
}

func DisplayDuration(start int, end int) string {
	if end == 0 || start == 0 {
		return "-"
	}

	duration := float32(start-end) / float32(1000)

	return fmt.Sprintf("%v", duration)
}

func AreAllTestsComplete(output *reflect.GetStatusOutput) bool {
	numComplete := 0

	for _, test := range output.Tests {
		if (test.Status == "succeeded") || (test.Status == "failed") {
			numComplete++
		}
	}

	return numComplete == len(output.Tests)
}
