package command

import (
	"encoding/json"
	"github.com/mlambda-net/identity/pkg/application/message"
	"net/http"
)

// DeleteUser godoc
// @Summary Delete the user
// @Produce json
// @Security ApiKeyAuth
// @Param data body int64 true "user id"
// @Success 200
// @Failure 401 {string} string "Unauthorized"
// @Failure 500 {string} string "Internal error"
// @Router /user [delete]
func (c *control) deleteUser(w http.ResponseWriter, r *http.Request) {
	var id int64
  _ = json.NewDecoder(r.Body).Decode(&id)
	token := r.Header.Get("Authorization")
	_, e := c.user.Token(token).Request(&message.Delete{Id: id}).Unwrap()

	if e != nil {
		http.Error(w, e.Error(), http.StatusInternalServerError)
	}
}
