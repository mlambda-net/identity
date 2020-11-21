package command

import (
	"encoding/json"
	"github.com/mlambda-net/identity/pkg/application/message"
	"github.com/mlambda-net/identity/pkg/infrastructure/endpoint/api/model"
	"net/http"
)

// AuthUser godoc
// @Summary Login the user
// @Produce json
// @Param data body model.Login true "user register"
// @Success 200
// @Failure 500 {string} string "Internal error"
// @Router /auth [post]
func (c *control) loginUser(w http.ResponseWriter, r *http.Request) {

	var login model.Login
  _ = json.NewDecoder(r.Body).Decode(&login)

	reply, e := c.auth.Request(&message.Authenticate{
		Email:    login.Email,
		Password: login.Password,
	}).Unwrap()

	if e != nil {
		http.Error(w, e.Error(), http.StatusInternalServerError)
	} else {
		rsp := reply.(*message.Token)
    _ = json.NewEncoder(w).Encode(rsp)
	}
}
