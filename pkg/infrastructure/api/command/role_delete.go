package command

import (
  "encoding/json"
  "github.com/mlambda-net/identity/pkg/application/message"
  "net/http"
)

// DeleteRole godoc
// @Summary Delete role
// @Produce json
// @Security ApiKeyAuth
// @Param data body string true "role id"
// @Success 200
// @Failure 401 {string} string "Unauthorized"
// @Failure 500 {string} string "Internal error"
// @Router /role [delete]
func (c *control) roleDelete(w http.ResponseWriter, r *http.Request) {
  var id string
  _ = json.NewDecoder(r.Body).Decode(&id)
  token := r.Header.Get("Authorization")
  _, e := c.role.Token(token).Request(&message.RoleDelete{Id: id}).Unwrap()

  if e != nil {
    http.Error(w, e.Error(), http.StatusInternalServerError)
  }
}
