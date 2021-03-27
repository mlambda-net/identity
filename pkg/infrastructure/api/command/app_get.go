package command

import (
"encoding/json"
"github.com/gorilla/mux"
"github.com/mlambda-net/identity/pkg/application/message"
"net/http"
)

// Get App by Id godoc
// @Summary Get the app
// @Produce json
// @Security ApiKeyAuth
// @Param id path string true "search by id"
// @Success 200 {object} message.AppResults
// @Failure 401 {string} string "Unauthorized"
// @Failure 404 {string} string "Not found"
// @Failure 500 {string} string "Internal error"
// @Router /app/{id} [get]
func (c *control) appGet(w http.ResponseWriter, r *http.Request) {
  token := r.Header.Get("Authorization")
  id := mux.Vars(r)["id"]
  rsp, e := c.app.Token(token).Request(&message.AppFind{Id: id}).Unwrap()
  if e != nil {
    http.Error(w, e.Error(), http.StatusInternalServerError)
  } else {
    switch m := rsp.(type) {
    case *message.AppResult:
      {
        _ = json.NewEncoder(w).Encode(m)
      }
    case *message.AppNotFound:
      {
        http.Error(w, "the app was not found", http.StatusNotFound)
      }
    }
  }
}

