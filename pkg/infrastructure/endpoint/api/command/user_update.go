package command

import (
	"encoding/json"
	"github.com/mlambda-net/identity/pkg/application/message"
	"github.com/mlambda-net/identity/pkg/infrastructure/endpoint/api/model"
	"net/http"
)

// UpdateUser godoc
// @Summary Updates the user
// @Produce json
// @Security ApiKeyAuth
// @Param data body model.Update true "user update"
// @Success 200
// @Failure 401 {string} string "Unauthorized"
// @Failure 500 {string} string "Internal error"
// @Router /user [put]
func (c *control) updateUser(w http.ResponseWriter, r *http.Request) {
	var user model.Update
  _ = json.NewDecoder(r.Body).Decode(&user)
	token := r.Header.Get("Authorization")

	_, e := c.user.Token(token).Request(&message.Update{
		Email: user.Email,
		Name:     user.Name,
		LastName: user.LastName,
	}).Unwrap()

	if e != nil {
		http.Error(w, e.Error(), http.StatusInternalServerError)
	}
}
