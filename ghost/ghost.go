package ghost

import (
	"github.com/go-resty/resty"
)

const VERSION = "v3"

func New(baseURL, apiKey string) (*Ghost, error) {
	_, err := generateJwt(apiKey)
	if err != nil {
		return nil, err
	}

	client := resty.New()
	client.HostURL = baseURL + "ghost/api/" + VERSION
	client.OnBeforeRequest(func(client *resty.Client, request *resty.Request) error {
		token, err := generateJwt(apiKey)
		if err != nil {
			return err
		}
		request.Header.Set("Authorization", "Ghost "+token)
		return nil
	})
	client.SetError(&errorsResponse{})

	return &Ghost{
		c: client,
	}, nil
}

func (g *Ghost) CreateMember(email string) error {
	req := membersRequest{Members: []member{
		{Email: email, Labels: []string{"API"}},
	}}
	res, err := g.c.R().SetBody(req).Post("admin/members")
	if err != nil {
		return err
	}
	if res.IsError() {
		return beautifyError(res.Error().(*errorsResponse))
	}
	return nil
}
