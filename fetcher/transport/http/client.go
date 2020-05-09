package http

import (
	"crypto/tls"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"time"

	"github.com/eugeneradionov/stocks/fetcher/config"
	"github.com/eugeneradionov/stocks/fetcher/logger"
	"go.uber.org/zap"
)

// NewClient http client constructor.
func NewClient() *http.Client {
	return NewClientWithTimeout(time.Duration(config.Get().HTTPClient.Timeout))
}

func NewClientWithTimeout(timeout time.Duration) *http.Client {
	tr := &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			Timeout:   60 * time.Second,
			KeepAlive: 60 * time.Second,
		}).DialContext,
		MaxIdleConns:          1024,
		MaxIdleConnsPerHost:   1024,
		IdleConnTimeout:       150 * time.Second,
		TLSHandshakeTimeout:   150 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
		TLSClientConfig:       &tls.Config{InsecureSkipVerify: false},
	}

	return &http.Client{Timeout: timeout, Transport: tr}
}

// CloseResponseBody closes response's body.
func CloseResponseBody(resp *http.Response) {
	_, err := io.Copy(ioutil.Discard, resp.Body)
	if err != nil {
		logger.Get().Error("discard response body", zap.Error(err))
	}

	err = resp.Body.Close()
	if err != nil {
		logger.Get().Error("close response body", zap.Error(err))
	}
}
