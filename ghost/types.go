package ghost

import "github.com/go-resty/resty"

type Ghost struct {
	c *resty.Client
}

type membersRequest struct {
	Members []member `json:"members"`
}

type member struct {
	Email  string   `json:"email"`
	Labels []string `json:"labels"`
}

type errorsResponse struct {
	Errors []struct {
		Context string `json:"context"`
	} `json:"errors"`
}
