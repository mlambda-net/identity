package command

import (
  "encoding/json"
  "github.com/mlambda-net/identity/pkg/application/message"
  "net/http"
)

// Get Users godoc
// @Summary Get Applications
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {array} message.AppResult
// @Failure 401 {string} string "Unauthorized"
// @Failure 500 {string} string "Internal error"
// @Router /app [get]
func (c *control) appGetAll(w http.ResponseWriter, r *http.Request) {
  token := r.Header.Get("Authorization")
  rsp, e := c.app.Token(token).Request(&message.AppSearch{}).Unwrap()
  if e != nil {
    http.Error(w, e.Error(), http.StatusInternalServerError)
  } else {
    data := rsp.(*message.AppResults)
    _ = json.NewEncoder(w).Encode(data.Results)
  }
}
