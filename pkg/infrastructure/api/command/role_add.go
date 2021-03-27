package command

import (
  "encoding/json"
  "github.com/mlambda-net/identity/pkg/application/message"
  "net/http"
)

// add the role godoc
// @Summary add the role
// @Produce json
// @Security ApiKeyAuth
// @Param data body message.RoleAdd true "role"
// @Success 200  {object} message.RoleId
// @Failure 401 {string} string "Unauthorized"
// @Failure 500 {string} string "Internal error"
// @Router /role [post]
func (c *control) roleAdd(w http.ResponseWriter, r *http.Request) {
  var data message.RoleAdd
  _ = json.NewDecoder(r.Body).Decode(&data)
  token := r.Header.Get("Authorization")
  id, err := c.role.Token(token).Request(&data).Unwrap()
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
  } else {
    _ = json.NewEncoder(w).Encode(id)
  }
}
