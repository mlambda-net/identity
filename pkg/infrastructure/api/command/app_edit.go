package command

import (
  "encoding/json"
  "github.com/mlambda-net/identity/pkg/application/message"
  "net/http"
)

// add the app godoc
// @Summary edit the app
// @Produce json
// @Security ApiKeyAuth
// @Param data body message.AppEdit true "app"
// @Success 200  {object} message.AppId
// @Failure 401 {string} string "Unauthorized"
// @Failure 500 {string} string "Internal error"
// @Router /app [put]
func (c *control) appEdit(w http.ResponseWriter, r *http.Request) {
  var data message.AppEdit
  _ = json.NewDecoder(r.Body).Decode(&data)
  token := r.Header.Get("Authorization")
  id, err := c.app.Token(token).Request(&data).Unwrap()
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
  } else {
    _ = json.NewEncoder(w).Encode(id)
  }
}
