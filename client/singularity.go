package client

import (
	"bytes"
	"encoding/json"
	"errors"
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
	api_scale_request            = "/api/requests/request/%v/scale"
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

func (c *SingularityClient) ScaleRequest(requestId string, numInstances int) (*models.RequestParent, error) {
	scaleRequest := models.SingularityScaleRequest{
		Instances: numInstances,
	}

	data, err := json.Marshal(scaleRequest)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PUT", c.urlFor(api_scale_request, requestId), bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}

	c.setStandardRequestsHeaders(req)
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode == http.StatusBadRequest {
		return nil, errors.New("Scaled number of instances must not match the current number of instances.")
	}

	data, err = ioutil.ReadAll(resp.Body)
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

	c.setStandardRequestsHeaders(req)

	return req, nil
}

func (c *SingularityClient) setStandardRequestsHeaders(req *http.Request) {
	for k, v := range c.headers {
		req.Header.Add(k, v)
	}
}

func (c *SingularityClient) urlFor(path string, a ...interface{}) string {
	return fmt.Sprintf(c.baseUri+path, a...)
}
