package command

import (
  "encoding/json"
  "github.com/mlambda-net/identity/pkg/application/message"
  "net/http"
)

// register godoc
// @Summary Register the user
// @Produce json
// @Param data body message.Register true "user register"
// @Success 200
// @Failure 500 {string} string "Internal error"
// @Router /register [post]
func (c *control) register(w http.ResponseWriter, r *http.Request) {
  var register *message.Register
  _ = json.NewDecoder(r.Body).Decode(&register)

	reply, e := c.auth.Request(register).Unwrap()

	if e != nil {
		http.Error(w, e.Error(), http.StatusInternalServerError)
	} else {
		rsp := reply.(*message.Response)
		id := rsp.Id
    _ = json.NewEncoder(w).Encode(id)
	}
}
