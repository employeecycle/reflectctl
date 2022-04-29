package reflect

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

type SuiteData struct {
	Name    string `json:"name"`
	SuiteID string `json:"suiteId"`
	Created uint   `json:"created"`
}

type Suites struct {
	Data []*SuiteData `json:"data"`
}

type ListSuiteOutput struct {
	Suites Suites
}

func (r *Reflect) ListSuites() (*ListSuiteOutput, error) {
	client := &http.Client{}

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/suites", r.URL()), nil)

	if err != nil {
		return nil, fmt.Errorf("ListSuites: %w", err)
	}

	req.Header.Add("X-API-KEY", r.APIKey)

	resp, err := client.Do(req)

	if err != nil {
		return nil, fmt.Errorf("ListSuites: %w", err)
	}

	defer resp.Body.Close()

	if !httpStatusOk(resp.StatusCode) {
		return nil, fmt.Errorf("ListSuites: response status code %s", resp.Status)
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, fmt.Errorf("ListSuites: %w", err)
	}

	output := &ListSuiteOutput{}
	err = json.Unmarshal(body, output)

	if err != nil {
		return nil, fmt.Errorf("ListSuites: %w", err)
	}

	return output, nil
}

type SuiteCookieOverride struct {
	Name     string `mapstructure:"name" json:"name"`
	Value    string `mapstructure:"value" json:"value"`
	Domain   string `mapstructure:"domain" json:"domain"`
	Expires  int32  `mapstructure:"expires" json:"expires"`
	HTTPOnly bool   `mapstructure:"httpOnly" json:"httpOnly"`
	MaxAge   int32  `mapstructure:"maxAge" json:"maxAge"`
	Path     string `mapstructure:"path" json:"path"`
	Secure   bool   `mapstructure:"secure" json:"secure"`
}

type SuiteOverrides struct {
	Cookies []SuiteCookieOverride `mapstructure:"cookies" json:"cookies"`
}

type ExecuteSuiteOutput struct {
	ExecutionID int `json:"executionId"`
}

type ExecuteSuiteOptions struct {
	Overrides *SuiteOverrides `json:"overrides"`
}

func (r *Reflect) ExecuteSuite(id string, options *ExecuteSuiteOptions) (*ExecuteSuiteOutput, error) {
	client := &http.Client{}

	var reqBody io.Reader

	if options == nil {
		reqBody = nil
	} else {
		reqBytes, err := json.Marshal(options)

		if err != nil {
			return nil, fmt.Errorf("ExecuteSuite: %w", err)
		}

		reqBody = bytes.NewBuffer(reqBytes)
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/suites/%s/executions", r.URL(), id), reqBody)

	if err != nil {
		return nil, fmt.Errorf("ExecuteSuite: %w", err)
	}

	req.Header.Add("X-API-KEY", r.APIKey)

	resp, err := client.Do(req)

	if err != nil {
		return nil, fmt.Errorf("ExecuteSuite: %w", err)
	}

	defer resp.Body.Close()

	if !httpStatusOk(resp.StatusCode) {
		return nil, fmt.Errorf("ExecuteSuite: response status code %s", resp.Status)
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, fmt.Errorf("ExecuteSuite: %w", err)
	}

	output := &ExecuteSuiteOutput{}
	err = json.Unmarshal(body, output)

	if err != nil {
		return nil, fmt.Errorf("ExecuteSuite: %w", err)
	}

	return output, nil
}
