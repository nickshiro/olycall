package rest

import (
	"context"
	"maps"
	"net/http"

	"olycall-server/pkg/rest"
)

type SuccessResponse[T any] struct {
	Data T `json:"data"` // Data
} // @name Response

type ErrorResponse struct {
	Error any `json:"error"` // Error details
} // @name ErrorResponse

type response struct {
	Status int `json:"status"`
	Data   any `json:"data,omitempty"`
	Error  any `json:"error,omitempty"`
}

type handlerResponse struct {
	Body    any
	Status  int
	Headers *http.Header
	IsError bool
}

func (c Controller) processHandlerResponse(ctx context.Context, w http.ResponseWriter, handlerResp handlerResponse) {
	if handlerResp.Headers != nil {
		maps.Copy(w.Header(), *handlerResp.Headers)
	}

	var resp response
	if handlerResp.IsError {
		resp = response{
			Status: handlerResp.Status,
			Error:  handlerResp.Body,
			Data:   nil,
		}
	} else {
		resp = response{
			Status: handlerResp.Status,
			Data:   handlerResp.Body,
			Error:  nil,
		}
	}

	if err := rest.WriteJSON(w, handlerResp.Status, resp); err != nil {
		panic(err)
	}

	c.logger.InfoContext(ctx, "response", "status", handlerResp.Status, "body", resp)
}

type handlerFunc func(r *http.Request) handlerResponse

func (c Controller) makeHandler(h handlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		c.processHandlerResponse(r.Context(), w, h(r))
	}
}
