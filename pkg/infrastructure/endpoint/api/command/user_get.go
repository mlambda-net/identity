package command

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/mlambda-net/identity/pkg/application/message"
	"net/http"
)

// DeleteUser godoc
// @Summary Get the user
// @Produce json
// @Security ApiKeyAuth
// @Param email path string true "search by email"
// @Success 200 {object} model.User
// @Failure 401 {string} string "Unauthorized"
// @Failure 500 {string} string "Internal error"
// @Router /user/{email} [get]
func (c *control) getUser(w http.ResponseWriter, r *http.Request) {

	email := mux.Vars(r)["email"]
	token := r.Header.Get("Authorization")

	result, e := c.user.Token(token).Request(&message.Find{
		Email: email,
	}).Unwrap()

	if e != nil {
		http.Error(w, e.Error(), http.StatusInternalServerError)
	} else {
		usr := result.(*message.Result)
    _ = json.NewEncoder(w).Encode(usr)
	}
}
