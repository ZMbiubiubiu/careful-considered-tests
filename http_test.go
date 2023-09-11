package main

import (
	"context"
	"io"
	"log"
	"net/http"
	"net/http/httptrace"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRequestConnection(t *testing.T) {
	url := "http://www.baidu.com"
	t.Run("request simple -- without read resp", func(t *testing.T) {
		for i := 0; i < 10; i++ {
			resp, err := http.Get(url)
			assert.Nil(t, err)
			resp.Body.Close()
			assert.Equal(t, http.StatusOK, resp.StatusCode)
		}
	})

	t.Run("request with trace -- without read resp", func(t *testing.T) {
		req, err := createHTTPGetRequestWithTrace(context.TODO(), url)
		assert.Nil(t, err)
		for i := 0; i < 10; i++ {
			resp, err := http.DefaultClient.Do(req)
			assert.Nil(t, err)
			resp.Body.Close()
			assert.Equal(t, http.StatusOK, resp.StatusCode)
		}
	})

	t.Run("request with trace --  read resp", func(t *testing.T) {
		req, err := createHTTPGetRequestWithTrace(context.TODO(), url)
		assert.Nil(t, err)
		for i := 0; i < 10; i++ {
			resp, err := http.DefaultClient.Do(req)
			assert.Nil(t, err)
			io.ReadAll(resp.Body)
			resp.Body.Close()
			assert.Equal(t, http.StatusOK, resp.StatusCode)
		}
	})

}

func createHTTPGetRequestWithTrace(ctx context.Context, url string) (*http.Request, error) {
	trace := &httptrace.ClientTrace{
		// httptrace.GotConnInfo's filed Reused represents the
		// connection has been previously used for another HTTP request
		GotConn: func(info httptrace.GotConnInfo) {
			log.Printf("Got Conn :%+v\n", info)
		},
	}
	ctxTrace := httptrace.WithClientTrace(ctx, trace)

	req, err := http.NewRequestWithContext(ctxTrace, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	return req, err
}
