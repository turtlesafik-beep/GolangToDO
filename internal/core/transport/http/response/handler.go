package core_http_response

import (
	"encoding/json"
	"fmt"
	"net/http"

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

func (h *HTTPResponseHamdler) PanicResponse(p any, msg string) {
	stausCode := http.StatusInternalServerError
	err := fmt.Errorf("unexpected panic: %v", p)

	h.log.Error(msg, zap.Error(err))
	h.rw.WriteHeader(stausCode)

	response := map[string]string{
		"messege": msg,
		"error":   err.Error(),
	}

	if err := json.NewEncoder(h.rw).Encode(response); err != nil {
		h.log.Error("write HTTP response", zap.Error(err))
	}
}
