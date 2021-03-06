package reflect

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Test struct {
	TestID    int    `json:"testId"`
	Status    string `json:"status"`
	Started   int    `json:"started"`
	Completed int    `json:"completed"`
	RunID     int    `json:"runId"`
}

type GetStatusOutput struct {
	ExecutionID int    `json:"executionId"`
	Tests       []Test `json:"tests"`
}

func (r *Reflect) GetStatus(id string) (*GetStatusOutput, error) {
	client := &http.Client{}

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/executions/%s", r.URL(), id), nil)

	if err != nil {
		return nil, fmt.Errorf("GetStatus: %w", err)
	}

	req.Header.Add("X-API-KEY", r.APIKey)

	resp, err := client.Do(req)

	if err != nil {
		return nil, fmt.Errorf("GetStatus: %w", err)
	}

	defer resp.Body.Close()

	if !httpStatusOk(resp.StatusCode) {
		return nil, fmt.Errorf("GetStatus: response status code %s", resp.Status)
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, fmt.Errorf("GetStatus: %w", err)
	}

	output := &GetStatusOutput{}
	err = json.Unmarshal(body, output)

	if err != nil {
		return nil, fmt.Errorf("GetStatus: %w", err)
	}

	return output, nil
}
