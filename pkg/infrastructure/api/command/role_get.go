package command

import (
  "encoding/json"
  "github.com/gorilla/mux"
  "github.com/mlambda-net/identity/pkg/application/message"
  "net/http"
)

// Get Role by Id godoc
// @Summary Get the role
// @Produce json
// @Security ApiKeyAuth
// @Param id path string true "search by id"
// @Success 200 {object} message.RoleResult
// @Failure 401 {string} string "Unauthorized"
// @Failure 404 {string} string "Not found"
// @Failure 500 {string} string "Internal error"
// @Router /role/{id} [get]
func (c *control) roleGet(w http.ResponseWriter, r *http.Request) {
  token := r.Header.Get("Authorization")
  id := mux.Vars(r)["id"]
  rsp, e := c.role.Token(token).Request(&message.RoleGet{Id: id}).Unwrap()
  if e != nil {
    http.Error(w, e.Error(), http.StatusInternalServerError)
  } else {
    switch m := rsp.(type) {
    case *message.RoleResult:
      {
        _ = json.NewEncoder(w).Encode(m)
      }
    case *message.RoleNotFound:
      {
        http.Error(w, "the role was not found", http.StatusNotFound)
      }
    }
  }
}

