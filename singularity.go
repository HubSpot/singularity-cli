package main

import (
	"net/http"
	"fmt"
	"encoding/json"
	"io/ioutil"
	"git.hubteam.com/zklapow/singularity-cli/models"
)

const (
	api_list_requests = "/api/requests"
)

type SingularityClient struct {
	baseUri string
	headers map[string]string
	httpClient *http.Client
}

func NewSingularityClient(baseUri string, headers map[string]string) *SingularityClient {
	return &SingularityClient{baseUri: baseUri, headers: headers, httpClient: http.DefaultClient}
}

func (c *SingularityClient) ListAllRequests() ([]models.RequestParent, error) {
	req, err := c.requestFor(api_list_requests)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	res := make([]models.RequestParent, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		switch err := err.(type) {
		case *json.UnmarshalTypeError:
			failedData := data[err.Offset-50:err.Offset+20]
			fmt.Printf("JSON Error at %v\n", string(failedData))
		}
	}

	return res, err
}

func (c *SingularityClient) requestFor(path string, a... interface{}) (*http.Request, error) {
	url := c.urlFor(api_list_requests)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	for k, v := range c.headers {
		req.Header.Add(k, v)
	}

	return req, nil
}

func (c *SingularityClient) urlFor(path string, a... interface{}) string {
	return fmt.Sprintf(c.baseUri + path, a...);
}
