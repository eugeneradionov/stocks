package common

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"

	exterrors "github.com/eugeneradionov/ext-errors"
	"github.com/eugeneradionov/stocks/api/config"
	"github.com/eugeneradionov/stocks/api/logger"
	"github.com/eugeneradionov/stocks/api/validator"
	v "github.com/go-playground/validator/v10"
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

// ProcessRequestBody unmarshal and validate request body and sends errors if any
func ProcessRequestBody(w http.ResponseWriter, r *http.Request, body interface{}) error {
	extErr := UnmarshalRequestBody(r, body)
	if extErr != nil {
		SendExtError(w, extErr)
		return extErr
	}

	httpErrs := ValidateRequestBody(r, body)
	if httpErrs != nil {
		SendExtErrors(w, http.StatusUnprocessableEntity, httpErrs)
		return httpErrs
	}

	return nil
}

// UnmarshalRequestBody unmarshals request body
func UnmarshalRequestBody(r *http.Request, body interface{}) exterrors.ExtError {
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		serverError := exterrors.NewInternalServerErrorError(errors.New("read JSON body"))

		logger.WithCtxValue(r.Context()).Error("Invalid request body", zap.Error(err))
		return serverError
	}
	defer r.Body.Close()

	err = json.Unmarshal(reqBody, body)
	if err != nil {
		serverError := exterrors.NewInternalServerErrorError(errors.New("parse JSON body"))
		logger.WithCtxValue(r.Context()).Error("Invalid JSON request body",
			zap.Error(err), zap.String("Corrupted JSON", string(reqBody)))

		return serverError
	}

	return nil
}

// ValidateRequestBody uses validator to validate request body
func ValidateRequestBody(r *http.Request, body interface{}) exterrors.ExtErrors {
	err := validator.Get().Struct(body)
	if err != nil {
		validationErrors := err.(v.ValidationErrors)
		serverError := validator.FormatErrors(validationErrors)

		logger.WithCtxValue(r.Context()).Info("request body validation failed", zap.Error(err))
		return serverError
	}

	return nil
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
