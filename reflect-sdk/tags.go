package reflect

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

type CookieOverride struct {
	Name     string `mapstructure:"name" json:"name"`
	Value    string `mapstructure:"value" json:"value"`
	Domain   string `mapstructure:"domain" json:"domain"`
	Expires  int32  `mapstructure:"expires" json:"expires"`
	HTTPOnly bool   `mapstructure:"httpOnly" json:"httpOnly"`
	MaxAge   int32  `mapstructure:"maxAge" json:"maxAge"`
	Path     string `mapstructure:"path" json:"path"`
	Secure   bool   `mapstructure:"secure" json:"secure"`
}

type Overrides struct {
	Cookies []CookieOverride `mapstructure:"cookies" json:"cookies"`
}

type CreateTagExecutionOutput struct {
	ExecutionID  int `json:"executionId"`
	NTestsQueued int `json:"NTestsQueued"`
}

type TestExecutionOptions struct {
	Overrides *Overrides `json:"overrides"`
}

func (r *Reflect) CreateTagExecution(tag string, options *TestExecutionOptions) (*CreateTagExecutionOutput, error) {
	client := &http.Client{}

	var reqBody io.Reader

	if options == nil {
		reqBody = nil
	} else {
		reqBytes, err := json.Marshal(options)

		if err != nil {
			return nil, fmt.Errorf("CreateTagExecution: %w", err)
		}

		reqBody = bytes.NewBuffer(reqBytes)
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/tags/%s/executions", r.URL(), tag), reqBody)

	if err != nil {
		return nil, fmt.Errorf("CreateTagExecution: %w", err)
	}

	req.Header.Add("X-API-KEY", r.APIKey)

	resp, err := client.Do(req)

	if err != nil {
		return nil, fmt.Errorf("CreateTagExecution: %w", err)
	}

	defer resp.Body.Close()

	if !httpStatusOk(resp.StatusCode) {
		return nil, errors.New(resp.Status)
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, fmt.Errorf("CreateTagExecution: %w", err)
	}

	output := &CreateTagExecutionOutput{}
	err = json.Unmarshal(body, output)

	if err != nil {
		return nil, fmt.Errorf("CreateTagExecution: %w", err)
	}

	return output, nil
}
