package rest

type SuccessResponse[T any] struct {
	Data T `json:"data"` // Data
} // @name Response

type ErrorResponse struct {
	Error any `json:"error"` // Error details
} // @name ErrorResponse
