package rest

import (
	"context"
	"log/slog"
	"net/http"
	"strconv"
	"time"

	"olycall-server/pkg/ctxlogger"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

func (c Controller) requestIDMw(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		ctx = ctxlogger.AppendCtx(ctx, slog.String("request_id", time.Now().String()))
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (c Controller) requestLoggingMw(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		c.logger.InfoContext(r.Context(), "request",
			"method", r.Method,
			"url", r.URL.String(),
			"remote_addr", r.RemoteAddr,
			"headers", r.Header,
			"body", r.Body,
		)
		next.ServeHTTP(w, r)
	})
}

const NotFoundErrMsg = "not found"

type contextKey string

var (
	accessTokenCtxKey contextKey = "access-token"
	userIDCtxKey      contextKey = "user-id"
)

func (c Controller) newIntURLParamMw(next http.Handler, key string, ctxKey contextKey) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if res := func() *int32 {
			str := chi.URLParam(r, key)
			if str == "" {
				return nil
			}

			parsedi64, err := strconv.ParseInt(str, 10, 32)
			if err != nil {
				return nil
			}

			parsed := int32(parsedi64)
			return &parsed
		}(); res != nil {
			ctx := context.WithValue(r.Context(), ctxKey, *res)
			next.ServeHTTP(w, r.WithContext(ctx))
			return
		}

		c.processHandlerResponse(r.Context(), w, handlerResponse{
			Body:    NotFoundErrMsg,
			Status:  http.StatusNotFound,
			IsError: true,
		})
	})
}

func (c Controller) newUUIDURLParamMw(next http.Handler, key string, ctxKey contextKey) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if res := func() *uuid.UUID {
			str := chi.URLParam(r, key)
			if str == "" {
				return nil
			}

			parsed, err := uuid.Parse(str)
			if err != nil {
				return nil
			}

			return &parsed
		}(); res != nil {
			ctx := context.WithValue(r.Context(), ctxKey, *res)
			next.ServeHTTP(w, r.WithContext(ctx))
			return
		}

		c.processHandlerResponse(r.Context(), w, handlerResponse{
			Body:    NotFoundErrMsg,
			Status:  http.StatusNotFound,
			IsError: true,
		})
	})
}

func (c Controller) userMw(next http.Handler) http.Handler {
	return c.newUUIDURLParamMw(next, "user-id", userIDCtxKey)
}

func (c Controller) getUserIDFromCtx(ctx context.Context) uuid.UUID {
	value, _ := ctx.Value(userIDCtxKey).(uuid.UUID)
	return value
}

func (c Controller) getAccessTokenFromCtx(ctx context.Context) string {
	value, _ := ctx.Value(accessTokenCtxKey).(string)
	return value
}

func (c Controller) accessTokenMw(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		accessTokenCookie, err := r.Cookie(accessTokenCookieName)
		if err != nil {
			c.processHandlerResponse(r.Context(), w, handlerResponse{
				Body:    err.Error(),
				Status:  http.StatusUnauthorized,
				IsError: true,
			})
			return
		}

		ctx := context.WithValue(r.Context(), accessTokenCtxKey, accessTokenCookie.Value)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
