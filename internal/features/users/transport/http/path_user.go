package users_transport_http

import (
	"fmt"
	"net/http"

	core_logger "github.com/turtlesafik-beep/GolangToDO/internal/core/logger"
	core_http_request "github.com/turtlesafik-beep/GolangToDO/internal/core/transport/http/request"
	core_http_response "github.com/turtlesafik-beep/GolangToDO/internal/core/transport/http/response"
)

type PathUserRequest struct {
	FullName    string `json:"full_name"`
	PhoneNumber string `json:"phone_number"`
}

func (h *UsersHTTPHandler) PathUser(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)

	var request PathUserRequest
	if err := core_http_request.DecodeAndValidate(r, &request); err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to decode and validate http request",
		)

		return
	}

	log.Debug(
		fmt.Sprintf(
			"PathUserRequest fields:\nFullName: '%s'\nPhoneNumber: '%s'",
			request.FullName,
			request.PhoneNumber,
		),
	)

	rw.WriteHeader(http.StatusOK)
}
