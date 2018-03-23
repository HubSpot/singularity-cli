package client

import (
	"git.hubteam.com/zklapow/singularity-cli/models"
	"fmt"
	"path/filepath"
	"strconv"
)

const (
	api_sandbox_browse = "/api/sandbox/%v/browse"
	api_sandbox_read = "/api/sandbox/%v/read"
)

func (c *SingularityClient) BrowseSandbox(taskId, path string) (*models.SingularitySandbox, error) {
	res := &models.SingularitySandbox{}
	url := fmt.Sprintf(api_sandbox_browse, taskId)

	if path != "" {
		url += fmt.Sprintf("?path=%v", filepath.Join(taskId, path))
	}

	finalurl := c.urlWithQueryParams(fmt.Sprintf(api_sandbox_browse, taskId), map[string]string{"path": filepath.Join(taskId, path)})

	err := c.getJson(res, finalurl)
	return res, err
}

func (c *SingularityClient) GetFileChunk(taskId, path string) (*models.MesosFileChunk, error) {
	return c.GetFileChunckWithOffset(taskId, path, 0, 0)
}


func (c *SingularityClient) GetFileChunckWithOffset(taskId, path string, offset, length uint64) (*models.MesosFileChunk, error) {
	res := &models.MesosFileChunk{}

	reqPath := fmt.Sprintf(api_sandbox_read, taskId)

	queryParams := map[string]string{}
	queryParams["path"] = filepath.Join(taskId, path)

	if length > 0 {
		queryParams["length"] = strconv.FormatUint(length, 10)
	}

	if offset > 0 {
		queryParams["offset"] = strconv.FormatUint(offset, 10)
	}

	finalurl := c.urlWithQueryParams(reqPath, queryParams)

	err := c.getJson(res, finalurl)
	return res, err
}
