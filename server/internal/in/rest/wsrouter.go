package rest

import "net/http"

func (c Controller) primaryWs(w http.ResponseWriter, r *http.Request) {
	c.getAccessToken(r)
}
