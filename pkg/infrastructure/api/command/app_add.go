package command

import (
  "encoding/json"
  "github.com/mlambda-net/identity/pkg/application/message"
  "net/http"
)

// add the app godoc
// @Summary add the app
// @Produce json
// @Security ApiKeyAuth
// @Param data body message.AppAdd true "app"
// @Success 200  {object} message.AppId
// @Failure 401 {string} string "Unauthorized"
// @Failure 500 {string} string "Internal error"
// @Router /app [post]
func (c *control) appAdd(w http.ResponseWriter, r *http.Request) {
  var data message.AppAdd
  _ = json.NewDecoder(r.Body).Decode(&data)
  token := r.Header.Get("Authorization")
  id, err := c.app.Token(token).Request(&data).Unwrap()
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
  } else {
    _ = json.NewEncoder(w).Encode(id)
  }
}
