package main

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRequestWithoutReadRespBody(t *testing.T) {
	url := "http://www.baidu.com"
	for i := 0; i < 10; i++ {
		resp, err := http.Get(url)
		assert.Nil(t, err)
		resp.Body.Close()
		assert.Equal(t, http.StatusOK, resp.StatusCode)
	}
}
