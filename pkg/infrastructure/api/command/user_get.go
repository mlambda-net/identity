package command

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/mlambda-net/identity/pkg/application/message"
	"net/http"
)

// Get User godoc
// @Summary Get the user by id
// @Produce json
// @Security ApiKeyAuth
// @Param id path string true "search by id"
// @Success 200 {object} message.Result
// @Failure 401 {string} string "Unauthorized"
// @Failure 500 {string} string "Internal error"
// @Router /user/{id} [get]
func (c *control) getUser(w http.ResponseWriter, r *http.Request) {
  id := mux.Vars(r)["id"]
  token := r.Header.Get("Authorization")

  result, e := c.user.Token(token).Request(&message.GetUser{
    Id: id,
  }).Unwrap()
  if e != nil {
    http.Error(w, e.Error(), http.StatusInternalServerError)
  } else {
    _ = json.NewEncoder(w).Encode(result.(*message.Result))
  }
}
