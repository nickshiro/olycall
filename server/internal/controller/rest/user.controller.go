package rest

import (
	"net/http"
)

// @Summary	Get user by ID
// @Tags		Users
// @Param		user-id	path		string	true	"User ID (UUID)"	format(uuid)
// @Success	200		{object}	SuccessResponse[service.User]
// @Failure	401		{object}	ErrorResponse
// @Failure	404		{object}	ErrorResponse
// @Failure	500		{object}	ErrorResponse
// @Router		/user/{user-id} [get]
func (c Controller) getUser(r *http.Request) handlerResponse {
	userID := c.getUserIDFromCtx(r.Context())

	getUserResp, err := c.service.GetUser(r.Context(), userID)
	if err != nil {
		return c.handleError(err)
	}

	return handlerResponse{
		Body:    getUserResp,
		Status:  http.StatusOK,
		IsError: false,
	}
}

// @Summary	Get current user
// @Tags		Users
// @Success	200	{object}	SuccessResponse[service.User]
// @Failure	401	{object}	ErrorResponse
// @Failure	500	{object}	ErrorResponse
// @Router		/user/me [get]
func (c Controller) me(r *http.Request) handlerResponse {
	accessToken := c.getAccessTokenFromCtx(r.Context())

	getUserResp, err := c.service.GetMe(r.Context(), accessToken)
	if err != nil {
		return c.handleError(err)
	}

	return handlerResponse{
		Body:    getUserResp,
		Status:  http.StatusOK,
		IsError: false,
	}
}
