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

type handlerResponse struct {
	Body    any
	Status  int
	Headers *http.Header
}

func (c Controller) processHandlerResponse(ctx context.Context, w http.ResponseWriter, handlerResp handlerResponse) {
	if handlerResp.Headers != nil {
		maps.Copy(w.Header(), *handlerResp.Headers)
	}

	var resp any
	if e, ok := handlerResp.Body.(error); ok {
		resp = ErrorResponse{
			Error: e.Error(),
		}

		c.logger.InfoContext(ctx, "error response", "status", handlerResp.Status, "body", e)
	} else {
		resp = SuccessResponse[any]{
			Data: handlerResp.Body,
		}

		c.logger.InfoContext(ctx, "response", "status", handlerResp.Status, "body", handlerResp.Body)
	}

	if err := rest.WriteJSON(w, handlerResp.Status, resp); err != nil {
		c.logger.InfoContext(ctx, "write json error", "error", err.Error(), "status", handlerResp.Status, "body", resp)

		return
	}
}

type handlerFunc func(r *http.Request) handlerResponse

func (c Controller) makeHandler(h handlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		c.processHandlerResponse(r.Context(), w, h(r))
	}
}
