package common

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	exterrors "github.com/eugeneradionov/ext-errors"
	"github.com/eugeneradionov/stocks/api/config"
	"github.com/eugeneradionov/stocks/api/logger"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

// SendResponse - encode response to json and send it.
func SendResponse(w http.ResponseWriter, statusCode int, respBody interface{}) {
	binRespBody, err := json.Marshal(respBody)
	if err != nil {
		logger.Get().Error("failed to marshal response body to json", zap.Error(err))
		statusCode = http.StatusInternalServerError
	}

	SendRawResponse(w, statusCode, binRespBody)
}

// SendRawResponse sends any raw ([]byte) response.
func SendRawResponse(w http.ResponseWriter, statusCode int, binBody []byte) {
	w.Header().Set("Content-Type", "application/json;charset=utf-8")

	w.WriteHeader(statusCode)
	_, err := w.Write(binBody)
	if err != nil {
		logger.Get().Error("failed to write response body", zap.Error(err))
	}
}

// SendExtError sends extError with code selected based on the ErrorCode.
func SendExtError(w http.ResponseWriter, extError exterrors.ExtError) {
	errs := exterrors.NewExtErrors()
	errs.Add(extError)

	SendExtErrors(w, extError.HTTPCode(), errs)
}

// SendExtError sends extErrors.
func SendExtErrors(w http.ResponseWriter, code int, httpErrors exterrors.ExtErrors) {
	SendResponse(w, code, httpErrors)
}

// GetLimitAndOffset returns limit and offset from query parameters.
func GetLimitAndOffset(urlQuery url.Values) (_, _ int, extErr exterrors.ExtError) {
	limitStr := urlQuery.Get("limit")
	limit, err := strconv.ParseInt(limitStr, 0, 64)
	if err != nil {
		extErr = exterrors.NewBadRequestError(errors.Wrap(err, "failed to parse limit from url"))
		return
	}

	if limit <= 0 {
		extErr = exterrors.NewBadRequestError(errors.New("'limit' must be greater than 0"))
		return
	}

	if limit > config.Get().PaginationMaxLimit {
		extErr = exterrors.NewBadRequestError(
			fmt.Errorf("'limit' must be less than %d", config.Get().PaginationMaxLimit),
		)
		return
	}

	offsetStr := urlQuery.Get("offset")
	offset, err := strconv.ParseInt(offsetStr, 0, 64)
	if err != nil {
		extErr = exterrors.NewBadRequestError(errors.Wrap(err, "failed to parse offset from url"))
		return
	}

	if offset < 0 {
		extErr = exterrors.NewBadRequestError(errors.New("'offset' must be greater or equal to 0"))
		return
	}

	return int(limit), int(offset), nil
}
