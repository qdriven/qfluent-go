package http

import (
	"github.com/imroc/req/v3"
)

type HttpClient struct {
	Client  *req.Client
	Options map[string]string
}

var DefaultClient = NewHttpClient()

func NewHttpClient() *HttpClient {
	return &HttpClient{
		Client:  req.DefaultClient(),
		Options: nil,
	}
}

func (h HttpClient) Get(url string) (string, string, *req.Response) {
	resp := h.Client.R(). // Use R() to create a request.
				MustGet(url)
	return resp.Status, resp.String(), resp
}
