package rest

import (
	"net/http"
	"olycall-server/internal/core"
	"olycall-server/pkg/rest"
)

// @Summary	Get user by ID
// @Tags	Users
// @Param	user-id	path		string	true	"User ID (UUID)"	format(uuid)
// @Success	200		{object}	SuccessResponse[domain.User]
// @Failure	401		{object}	ErrorResponse
// @Failure	404		{object}	ErrorResponse
// @Failure	500		{object}	ErrorResponse
// @Router	/user/{user-id} [get]
func (c Controller) getUser(r *http.Request) handlerResponse {
	userID := c.getUserIDFromCtx(r.Context())

	getUserResp, err := c.service.GetUser(r.Context(), userID)
	if err != nil {
		return c.handleError(err)
	}

	return handlerResponse{
		Body:    getUserResp,
		Status:  http.StatusOK,
		Headers: nil,
	}
}

// @Summary	Get current user
// @Tags	Users
// @Success	200	{object}	SuccessResponse[domain.User]
// @Failure	401	{object}	ErrorResponse
// @Failure	500	{object}	ErrorResponse
// @Router	/user/me [get]
func (c Controller) getMe(r *http.Request) handlerResponse {
	accessToken := c.getAccessTokenFromCtx(r.Context())

	getUserResp, err := c.service.GetMe(r.Context(), accessToken)
	if err != nil {
		return c.handleError(err)
	}

	return handlerResponse{
		Body:    getUserResp,
		Status:  http.StatusOK,
		Headers: nil,
	}
}

type putMeBody struct {
	Username string `json:"username" validate:"required"`
} // @name PutMeBody

// @Summary	Update current user
// @Tags	Users
// @Param	body	body	putMeBody	true	"Request body"
// @Success	200	{object}	SuccessResponse[domain.User]
// @Failure	400	{object}	ErrorResponse
// @Failure	401	{object}	ErrorResponse
// @Failure	500	{object}	ErrorResponse
// @Router	/user/me [put]
func (c Controller) putMe(r *http.Request) handlerResponse {
	var body putMeBody
	if err := rest.ReadJSON(r, &body); err != nil {
		return handlerResponse{
			Body:    err,
			Status:  http.StatusBadRequest,
			Headers: nil,
		}
	}

	accessToken := c.getAccessTokenFromCtx(r.Context())

	updateMeResp, err := c.service.UpdateMe(r.Context(), &core.UpdateMeParams{
		AccessToken: accessToken,
		Username:    body.Username,
	})
	if err != nil {
		return c.handleError(err)
	}

	return handlerResponse{
		Body:    updateMeResp,
		Status:  http.StatusOK,
		Headers: nil,
	}
}
