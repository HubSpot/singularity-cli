package client

import (
	"bytes"
	"crypto/md5"
	"encoding/base64"
	"git.hubteam.com/zklapow/singularity-cli/models"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"fmt"
	"github.com/mitchellh/go-homedir"
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

	return ioutil.WriteFile(tmpfp, buffer.Bytes(), os.ModeTemporary|os.ModePerm)
}

func (c *SingularityClient) getCacheFilePath() (string, error) {
	dir, err := homedir.Expand("~/.sng_cache")
	if err != nil {
		return "", err
	}

	if _, err := os.Stat(dir); os.IsNotExist(err) {
		os.Mkdir(dir, os.ModePerm)
	}

	hasher := md5.New()
	md5sum := hasher.Sum([]byte(c.baseUri))

	b64hash := make([]byte, base64.StdEncoding.EncodedLen(len(md5sum)))
	base64.StdEncoding.Encode(b64hash, []byte(md5sum))

	tmpfp := filepath.Join(dir, string(b64hash))

	return tmpfp, nil
}
