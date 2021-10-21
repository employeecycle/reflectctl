// Package reflect is an SDK wrapper around the reflect HTTP API
package reflect

import "fmt"

type Reflect struct {
	APIKey  string
	BaseURL string
	Version string
}

type NewReflectInput struct {
	APIKey string
}

func NewReflect(p *NewReflectInput) *Reflect {
	r := &Reflect{
		APIKey:  p.APIKey,
		BaseURL: "https://api.reflect.run",
		Version: "v1",
	}

	return r
}

func (r *Reflect) URL() string {
	return fmt.Sprintf("%s/%s", r.BaseURL, r.Version)
}

func httpStatusOk(statusCode int) bool {
	return statusCode >= 200 && statusCode < 300
}
