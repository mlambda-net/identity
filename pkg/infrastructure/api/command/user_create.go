package command

import (
  "encoding/json"
  "github.com/mlambda-net/identity/pkg/application/message"
  "github.com/mlambda-net/identity/pkg/infrastructure/api/model"
  "net/http"
)

// CreateUser godoc
// @Summary Creates the user
// @Produce json
// @Param data body model.Register true "user register"
// @Success 200
// @Failure 500 {string} string "Internal error"
// @Router /user [post]
func (c *control) createUser(w http.ResponseWriter, r *http.Request) {

	var register model.Register
  _ = json.NewDecoder(r.Body).Decode(&register)

	reply, e := c.auth.Request(&message.Create{
		Name:     register.Name,
		LastName: register.LastName,
		Email:    register.Login,
		Password: register.Password,
	}).Unwrap()

	if e != nil {
		http.Error(w, e.Error(), http.StatusInternalServerError)
	} else {
		rsp := reply.(*message.Response)
		id := rsp.Id
    _ = json.NewEncoder(w).Encode(id)
	}
}
