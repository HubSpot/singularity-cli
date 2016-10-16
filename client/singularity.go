package client

import (
	"encoding/json"
	"fmt"
	"git.hubteam.com/zklapow/singularity-cli/models"
	"io/ioutil"
	"net/http"
)

const (
	api_list_requests            = "/api/requests"
	api_get_request              = "/api/requests/%v"
	api_pause_request            = "/api/requests/request/%v/pause"
	api_unpause_request          = "/api/requests/request/%v/unpause"
	api_active_tasks_for_request = "/api/history/request/%v/tasks/active"
)

type SingularityClient struct {
	baseUri    string
	headers    map[string]string
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
			failedData := data[err.Offset-50 : err.Offset+20]
			fmt.Printf("JSON Error at %v\n", string(failedData))
		}
	}

	// Always cache the requests after we load the whole lists
	c.cacheRequestList(res)

	return res, err
}

func (c *SingularityClient) GetRequest(requestId string) (*models.RequestParent, error) {
	req, err := c.requestFor(api_get_request, requestId)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	res := &models.RequestParent{}
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}

	return res, err
}

func (c *SingularityClient) GetActiveTasksFor(requestId string) ([]models.SingularityTaskIdHistory, error) {
	req, err := c.requestFor(api_active_tasks_for_request, requestId)
	if err != nil {
		return nil, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	res := make([]models.SingularityTaskIdHistory, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}

	return res, err
}

func (c *SingularityClient) PauseRequest(requestId string) (*models.RequestParent, error) {
	return c.pauseInternal(api_pause_request, requestId)
}

func (c *SingularityClient) UnPauseRequest(requestId string) (*models.RequestParent, error) {
	return c.pauseInternal(api_unpause_request, requestId)
}

func (c *SingularityClient) pauseInternal(path, requestId string) (*models.RequestParent, error) {
	req, err := c.requestFor(path, requestId)
	if err != nil {
		return nil, err
	}

	req.Method = "POST"
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	res := &models.RequestParent{}
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (c *SingularityClient) requestFor(path string, a ...interface{}) (*http.Request, error) {
	url := c.urlFor(path, a...)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	for k, v := range c.headers {
		req.Header.Add(k, v)
	}

	return req, nil
}

func (c *SingularityClient) urlFor(path string, a ...interface{}) string {
	return fmt.Sprintf(c.baseUri+path, a...)
}
