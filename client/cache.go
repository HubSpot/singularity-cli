package client

import (
	"io/ioutil"
	"path/filepath"
	"crypto/md5"
	"encoding/base64"
	"fmt"
	"os"
	"strings"
	"git.hubteam.com/zklapow/singularity-cli/models"
	"bytes"
)

func (c *SingularityClient) GetCachedRequestList() ([]string, error) {
	tmpfp, err := c.getCacheFilePath()
	if err != nil {
		return nil, err
	}

	file, err := os.Open(tmpfp)
	if err != nil {
		return nil, err
	}

	cachedData, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	return strings.Split(string(cachedData), "\n"), nil
}

func (c *SingularityClient) cacheRequestList(requests []models.RequestParent) error {
	tmpfp, err := c.getCacheFilePath()
	if err != nil {
		return err
	}

	buffer := bytes.NewBuffer([]byte{})
	for _, req := range requests {
		buffer.WriteString(req.Request.Id + "\n")
	}

	return ioutil.WriteFile(tmpfp, buffer.Bytes(), os.ModeTemporary | os.ModePerm)
}

func (c *SingularityClient) getCacheFilePath() (string, error) {
	dir, err := ioutil.TempDir("", "sng")
	if err != nil {
		return "", err
	}

	hasher := md5.New()
	md5sum := hasher.Sum([]byte(c.baseUri))

	b64hash := make([]byte, base64.StdEncoding.EncodedLen(len(md5sum)))
	base64.StdEncoding.Encode(b64hash, []byte(md5sum))

	tmpfp := filepath.Join(dir, string(b64hash))

	return tmpfp, nil
}
