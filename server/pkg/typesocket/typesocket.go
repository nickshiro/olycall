package typesocket

import (
	"context"
	"encoding/json"
)

const JSONRPCVer = "2.0"

type RPCEvent struct {
	JSONRPC string `json:"jsonrpc"`
	Event   string `json:"event"`
	Data    any    `json:"data,omitempty"`
}

type RPCRequest struct {
	JSONRPC string          `json:"jsonrpc"`
	ID      json.RawMessage `json:"id,omitempty"`
	Method  string          `json:"method"`
	Params  json.RawMessage `json:"params,omitempty"`
}

type RPCResponse struct {
	JSONRPC string          `json:"jsonrpc"`
	ID      json.RawMessage `json:"id"`
	Result  any             `json:"result,omitempty"`
	Error   *RPCError       `json:"error,omitempty"`
}

type RPCError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

func NewRPCError(code int, msg string) *RPCError {
	return &RPCError{Code: code, Message: msg}
}

type RPCErrorable interface {
	ToRPCError() *RPCError
}

var (
	ErrParse          = NewRPCError(-32700, "Parse error")
	ErrInvalidReq     = NewRPCError(-32600, "Invalid request")
	ErrMethodNotFound = NewRPCError(-32601, "Method not found")
	ErrInvalidParams  = NewRPCError(-32602, "Invalid params")
)

func NewInternalError(msg string) *RPCError {
	return &RPCError{
		Code:    -32603,
		Message: msg,
	}
}

type Server struct {
	handlers map[string]handler
}

func NewServer() *Server {
	return &Server{
		handlers: make(map[string]handler),
	}
}

type handler func(ctx context.Context, id json.RawMessage, rawParams json.RawMessage) *RPCResponse

func Register[Params any, Result any](s *Server, method string, fn func(context.Context, *Params) (Result, *RPCError)) {
	s.handlers[method] = func(ctx context.Context, id json.RawMessage, rawParams json.RawMessage) *RPCResponse {
		var p *Params
		if len(rawParams) > 0 {
			p = new(Params)
			if err := json.Unmarshal(rawParams, p); err != nil {
				return makeErrorResponse(id, ErrInvalidParams)
			}
		}

		result, err := fn(ctx, p)
		if err != nil {
			return makeErrorResponse(id, err)
		}

		return &RPCResponse{
			JSONRPC: JSONRPCVer,
			ID:      id,
			Result:  result,
		}
	}
}

func (s *Server) HandleRequest(ctx context.Context, raw []byte) ([]byte, error) {
	var req RPCRequest
	if err := json.Unmarshal(raw, &req); err != nil {
		return json.Marshal(makeErrorResponse(nil, ErrParse))
	}

	if req.JSONRPC != JSONRPCVer || req.Method == "" || len(req.ID) == 0 {
		return json.Marshal(makeErrorResponse(req.ID, ErrInvalidReq))
	}

	handler, ok := s.handlers[req.Method]
	if !ok {
		return json.Marshal(makeErrorResponse(req.ID, ErrMethodNotFound))
	}

	resp := handler(ctx, req.ID, req.Params)

	return json.Marshal(resp)
}

func MakeEvent(eventName string, data any) []byte {
	resp := RPCEvent{
		JSONRPC: JSONRPCVer,
		Event:   eventName,
		Data:    data,
	}

	b, err := json.Marshal(resp)
	if err != nil {
		panic(err)
	}

	return b
}

func makeErrorResponse(id json.RawMessage, err *RPCError) *RPCResponse {
	return &RPCResponse{
		JSONRPC: JSONRPCVer,
		ID:      id,
		Error:   err,
	}
}
