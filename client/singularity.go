package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"git.hubteam.com/zklapow/singularity-cli/models"
	"io/ioutil"
	"net/http"
	"fmt"
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
	res := make([]models.RequestParent, 0)
	err := c.getJsonSimple(&res, api_list_requests)
	if err != nil {
		return nil, err
	}

	// Always cache the requests after we load the whole lists
	c.cacheRequestList(res)

	return res, err
}

func (c *SingularityClient) GetRequest(requestId string) (models.RequestParent, error) {
	res := models.RequestParent{}
	err := c.getJsonSimple(&res, fmt.Sprintf(api_get_request, requestId))
	return res, err
}

func (c *SingularityClient) GetActiveTasksFor(requestId string) ([]models.SingularityTaskIdHistory, error) {
	res := make([]models.SingularityTaskIdHistory, 0)
	err := c.getJsonSimple(&res, fmt.Sprintf(api_active_tasks_for_request, requestId))

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

	req, err := http.NewRequest("PUT", c.urlFor(fmt.Sprintf(api_scale_request, requestId)).String(), bytes.NewBuffer(data))
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
	req, err := c.requestFor(c.urlFor(fmt.Sprintf(path, requestId)))
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
