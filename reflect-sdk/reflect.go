package reflect

import "fmt"

type Reflect struct {
	ApiKey  string
	BaseUrl string
	Version string
}

type NewReflectInput struct {
	ApiKey string
}

func NewReflect(p *NewReflectInput) *Reflect {
	r := &Reflect{
		ApiKey:  p.ApiKey,
		BaseUrl: "https://api.reflect.run",
		Version: "v1",
	}

	return r
}

func (r *Reflect) Url() string {
	return fmt.Sprintf("%s/%s", r.BaseUrl, r.Version)
}

func httpStatusOk(statusCode int) bool {
	return statusCode >= 200 && statusCode < 300
}
