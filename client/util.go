package client

import (
	"io/ioutil"
	"encoding/json"
	"net/http"
	"fmt"
)

func (c *SingularityClient) getJson(result interface{}, path string, args... interface{}) error {
	req, err := c.requestFor(path, args...)
	if err != nil {
		return err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, result)
	switch err := err.(type) {
	case *json.UnmarshalTypeError:
		failedData := data[err.Offset-50 : err.Offset+20]
		fmt.Printf("JSON Error at %v\n", string(failedData))
	}

	return err
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
