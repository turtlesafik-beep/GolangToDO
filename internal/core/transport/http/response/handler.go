package core_http_response

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	core_errors "github.com/turtlesafik-beep/GolangToDO/internal/core/errors"
	core_logger "github.com/turtlesafik-beep/GolangToDO/internal/core/logger"
	"go.uber.org/zap"
)

type HTTPResponseHamdler struct {
	log *core_logger.Logger
	rw  http.ResponseWriter
}

func NewHTTPResponseHandler(log *core_logger.Logger, rw http.ResponseWriter) *HTTPResponseHamdler {
	return &HTTPResponseHamdler{
		log: log,
		rw:  rw,
	}
}

func (h *HTTPResponseHamdler) JSONResponse(
	responseBody any,
	statusCode int,
) {
	h.rw.WriteHeader(statusCode)

	if err := json.NewEncoder(h.rw).Encode(responseBody); err != nil {
		h.log.Error("write http response", zap.Error(err))
	}
}

func (h *HTTPResponseHamdler) ErrorResponse(err error, msg string) {
	var (
		statusCode int
		logFunc    func(string, ...zap.Field)
	)

	switch {
	case errors.Is(err, core_errors.ErrInavildArgument):
		statusCode = http.StatusBadRequest
		logFunc = h.log.Warn

	case errors.Is(err, core_errors.ErrNotFound):
		statusCode = http.StatusNotFound
		logFunc = h.log.Debug

	case errors.Is(err, core_errors.ErrConflict):
		statusCode = http.StatusConflict
		logFunc = h.log.Warn

	default:
		statusCode = http.StatusInternalServerError
		logFunc = h.log.Error
	}

	logFunc(msg, zap.Error(err))

	h.errorResponse(
		statusCode,
		err,
		msg,
	)
}

func (h *HTTPResponseHamdler) PanicResponse(p any, msg string) {
	statusCode := http.StatusInternalServerError
	err := fmt.Errorf("unexpected panic: %v", p)

	h.log.Error(msg, zap.Error(err))
	h.errorResponse(
		statusCode,
		err,
		msg,
	)
}

func (h *HTTPResponseHamdler) errorResponse(
	statusCode int,
	err error,
	msg string,
) {
	response := map[string]string{
		"messege": msg,
		"error":   err.Error(),
	}

	h.JSONResponse(
		response,
		statusCode,
	)
}
