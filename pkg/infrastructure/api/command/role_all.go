package command

import (
  "encoding/json"
  "github.com/mlambda-net/identity/pkg/application/message"
  "net/http"
)

// Get Roles godoc
// @Summary Get Roles
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {array} message.RoleResult
// @Failure 401 {string} string "Unauthorized"
// @Failure 500 {string} string "Internal error"
// @Router /role [get]
func (c *control) roleAll(w http.ResponseWriter, r *http.Request) {
  token := r.Header.Get("Authorization")
  rsp, e := c.role.Token(token).Request(&message.RoleSearch{ Filter: ""}).Unwrap()
  if e != nil {
    http.Error(w, e.Error(), http.StatusInternalServerError)
  } else {
    data := rsp.(*message.RoleResults)
    _ = json.NewEncoder(w).Encode(data.Results)
  }
}
