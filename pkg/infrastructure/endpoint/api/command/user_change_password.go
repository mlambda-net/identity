package command

import (
	"encoding/json"
	"github.com/mlambda-net/identity/pkg/application/message"
	"github.com/mlambda-net/identity/pkg/infrastructure/endpoint/api/model"
	"net/http"
)

// ChangePassword godoc
// @Summary Change the user password
// @Produce json
// @Security ApiKeyAuth
// @Param data body model.ChangePassword true "change password"
// @Success 200
// @Failure 401 {string} string "Unauthorized"
// @Failure 500 {string} string "Internal error"
// @Router /user/change_password [post]
func (c *control) changePassword(w http.ResponseWriter, r *http.Request) {

	var register model.ChangePassword

	json.NewDecoder(r.Body).Decode(&register)
	token := r.Header.Get("Authorization")

	reply, e := c.user.Token(token).Request(&message.ChangePassword{
		Email:       register.Email,
		Password:    register.NewPassword,
		OldPassword: register.OldPassword,
	}).Unwrap()

	if e != nil {
		http.Error(w, e.Error(), http.StatusInternalServerError)
	} else {
		rsp := reply.(*message.Response)
		id := rsp.Id
		json.NewEncoder(w).Encode(id)
	}
}
