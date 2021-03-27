package command

import (
  "encoding/json"
  "github.com/mlambda-net/identity/pkg/application/message"
  "net/http"
)

// CreateUser godoc
// @Summary Creates the user
// @Produce json
// @Param data body message.Create true "user create"
// @Success 200
// @Failure 500 {string} string "Internal error"
// @Router /user [post]
func (c *control) createUser(w http.ResponseWriter, r *http.Request) {
  var user *message.Create
  _ = json.NewDecoder(r.Body).Decode(&user)
  token := r.Header.Get("Authorization")

  reply, e := c.user.Token(token).Request(user).Unwrap()

  if e != nil {
    http.Error(w, e.Error(), http.StatusInternalServerError)
  } else {
    rsp := reply.(*message.Response)
    id := rsp.Id
    _ = json.NewEncoder(w).Encode(id)
  }
}
