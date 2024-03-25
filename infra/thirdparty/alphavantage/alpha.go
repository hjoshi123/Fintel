package alphavantage

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/hjoshi123/fintel/infra/util"
)

type AlphaVantage struct {
	BaseURL string
	Client  *http.Client
	APIKey  string
}

func NewAVClient(ctx context.Context, httpClient *http.Client, baseURL, apiKey string) (*AlphaVantage, error) {
	if apiKey == "" || baseURL == "" {
		return nil, errors.New("API Key/ Base URL is required")
	}

	av := new(AlphaVantage)
	if httpClient == nil {
		av.Client = http.DefaultClient
	} else {
		av.Client = httpClient
	}

	av.APIKey = apiKey
	av.BaseURL = baseURL
	return av, nil
}

func (av *AlphaVantage) NewRequest(ctx context.Context, method string, body any, params map[string]string) (*http.Request, error) {
	switch method {
	case http.MethodGet:
		req, err := http.NewRequestWithContext(ctx, method, av.BaseURL, nil)
		if err != nil {
			util.Log.Debug().Ctx(ctx).Err(err).Msg("Failed to create new alpha vantage request")
			return nil, err
		}

		q := req.URL.Query()
		for k, v := range params {
			q.Add(k, v)
		}
		q.Add("apikey", av.APIKey)
		req.URL.RawQuery = q.Encode()
		return req, nil
	default:
		return nil, errors.New("Invalid method")
	}
}

func (av *AlphaVantage) Do(ctx context.Context, req *http.Request, v any) (*http.Response, error) {
	resp, err := av.Client.Do(req)
	if err != nil {
		util.Log.Error().Ctx(ctx).Err(err).Msg("Failed to do request")
		return nil, err
	}

	switch resp.Header.Get("Content-Type") {
	case "application/json":
		err = json.NewDecoder(resp.Body).Decode(v)
		if err != nil {
			util.Log.Error().Ctx(ctx).Err(err).Msg("Failed to decode response")
			return nil, err
		}
	}

	if resp.StatusCode != http.StatusOK {
		util.Log.Error().Ctx(ctx).Msg("Failed to get response")
		return nil, errors.New("Failed to get response")
	}

	return resp, nil
}
