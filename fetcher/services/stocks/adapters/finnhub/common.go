package finnhub

import (
	"fmt"
	"io/ioutil"
	"net/http"

	exterrors "github.com/eugeneradionov/ext-errors"
	"github.com/eugeneradionov/stocks/fetcher/logger"
	httpCli "github.com/eugeneradionov/stocks/fetcher/transport/http"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

func (f Finnhub) sendRequest(path string, params map[string]string) (_ []byte, extErr exterrors.ExtError) {
	var reqURL = fmt.Sprintf("%s/%s", f.cfg.BaseURL, path)

	req, err := http.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, exterrors.NewInternalServerErrorError(errors.Wrap(err, "new finnhub request"))
	}

	f.prepareReqParams(req, params)

	logger.Get().Info("finnhub send request", zap.String("path", path), zap.Any("params", params))

	resp, err := f.httpClient.Do(req) // nolint:bodyclose
	if err != nil {
		return nil, exterrors.NewInternalServerErrorError(errors.Wrap(err, "send http request to finnhub"))
	}
	defer httpCli.CloseResponseBody(resp) // discard and close response body

	switch resp.StatusCode {
	case http.StatusUnauthorized:
		return nil, exterrors.NewInternalServerErrorError(errors.New("finnhub unauthorized"))
	case http.StatusForbidden:
		return nil, exterrors.NewInternalServerErrorError(errors.New("finnhub forbidden"))
	case http.StatusNotFound:
		return nil, exterrors.NewInternalServerErrorError(errors.New("finnhub not found"))
	case http.StatusPaymentRequired:
		return nil, exterrors.NewInternalServerErrorError(errors.New("finnhub payment required"))
	case http.StatusTooManyRequests:
		return nil, exterrors.NewInternalServerErrorError(errors.New("finnhub too many requests"))
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, exterrors.NewInternalServerErrorError(errors.Wrap(err, "finnhub read response body"))
	}

	return data, nil
}

func (f Finnhub) prepareReqParams(req *http.Request, params map[string]string) {
	var query = req.URL.Query()

	query.Add("token", f.cfg.Token)

	for key, value := range params {
		query.Add(key, value)
	}

	req.URL.RawQuery = query.Encode()
}
