package reflect

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

type CreateTagExecutionOutput struct {
	ExecutionID  int `json:"executionId"`
	NTestsQueued int `json:"NTestsQueued"`
}

func (r *Reflect) CreateTagExecution(tag string) (*CreateTagExecutionOutput, error) {
	client := &http.Client{}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/tags/%s/executions", r.Url(), tag), nil)

	if err != nil {
		return nil, err
	}

	req.Header.Add("X-API-KEY", r.ApiKey)

	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if !httpStatusOk(resp.StatusCode) {
		return nil, errors.New(resp.Status)
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	output := &CreateTagExecutionOutput{}
	err = json.Unmarshal(body, output)

	if err != nil {
		return nil, err
	}

	return output, nil
}
