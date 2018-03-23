package client

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func (c *SingularityClient) getJsonSimple(result interface{}, reqPath string, args ...interface{}) error {
	finalReqUrl, err := url.Parse(c.baseUri + reqPath)
	if err != nil {
		return err
	}

	return c.getJson(result, finalReqUrl)
}

func (c *SingularityClient) getJson(result interface{}, reqUrl *url.URL) error {
	req, err := c.requestFor(reqUrl)
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

func (c *SingularityClient) requestFor(reqUrl *url.URL) (*http.Request, error) {
	req, err := http.NewRequest("GET", reqUrl.String(), nil)
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

func (c *SingularityClient) urlFor(path string) *url.URL {
	return c.urlWithQueryParams(path, map[string]string{})
}

func (c *SingularityClient) urlWithQueryParams(path string, queryParams map[string]string) *url.URL {
	finalurl, _ := url.Parse(c.baseUri + path)

	values := url.Values{}
	for k, v := range queryParams {
		values.Add(k, v)
	}

	finalurl.RawQuery = values.Encode()

	return finalurl
}
