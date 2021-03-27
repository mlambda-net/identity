package command

import (
  "encoding/json"
  "github.com/mlambda-net/identity/pkg/application/message"
  "net/http"
)

// UpdateUser godoc
// @Summary Updates the user
// @Produce json
// @Security ApiKeyAuth
// @Param data body message.Update true "user update"
// @Success 200
// @Failure 401 {string} string "Unauthorized"
// @Failure 500 {string} string "Internal error"
// @Router /user [put]
func (c *control) updateUser(w http.ResponseWriter, r *http.Request) {
  var user message.Update
  _ = json.NewDecoder(r.Body).Decode(&user)
  token := r.Header.Get("Authorization")
  _, e := c.user.Token(token).Request(&user).Unwrap()

  if e != nil {
    http.Error(w, e.Error(), http.StatusInternalServerError)
  }
}
