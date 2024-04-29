package finhistory

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/hjoshi123/fintel/infra/constants"
)

type Client struct {
	BaseURL     string
	BearerToken string
	client      *http.Client
}

func NewClient(baseURL, bearerToken string, cl *http.Client) (*Client, error) {
	if baseURL == "" {
		return nil, constants.ErrorNoBaseURL
	}

	if bearerToken == "" {
		return nil, constants.ErrorNoToken
	}

	if cl == nil {
		cl = http.DefaultClient
	}

	return &Client{
		BaseURL:     baseURL,
		BearerToken: bearerToken,
		client:      cl,
	}, nil
}

func (c *Client) Request(ctx context.Context, method, path string, params map[string]string, body any) (*http.Request, error) {
	rawURL := fmt.Sprintf("%s%s", c.BaseURL, path)

	parsedURL, err := url.Parse(rawURL)
	if err != nil {
		return nil, err
	}

	if len(params) > 0 && method == http.MethodGet {
		q := parsedURL.Query()
		for key, val := range params {
			q.Set(key, val)
		}

		parsedURL.RawQuery = q.Encode()
	} else if len(params) > 0 {
		return nil, errors.New("params cannot be greater than 0 without GET request1")
	}

	buf := &bytes.Buffer{}
	if body != nil {
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequestWithContext(ctx, method, parsedURL.String(), buf)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")

	return req, nil
}

func (c *Client) Do(req *http.Request, v interface{}) (*http.Response, error) {
	resp, err := c.client.Do(req)

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if v != nil {
		err = json.NewDecoder(resp.Body).Decode(v)
	}

	return resp, err
}
