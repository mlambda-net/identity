package command

import (
  "encoding/json"
  "github.com/mlambda-net/identity/pkg/application/message"
  "net/http"
  "strings"
)

// Get Users godoc
// @Summary Get users
// @Produce json
// @Param filter query string true "search by filter"
// @Security ApiKeyAuth
// @Success 200 {array} message.Result
// @Failure 401 {string} string "Unauthorized"
// @Failure 500 {string} string "Internal error"
// @Router /user [get]
func (c *control) getUsers(w http.ResponseWriter, r *http.Request) {
  token := r.Header.Get("Authorization")
  filter := r.URL.Query().Get("filter")
  rsp, e := c.user.Token(token).Request(&message.All{Filter: filter}).Unwrap()



  if e != nil {
    http.Error(w, e.Error(), http.StatusInternalServerError)
  } else {
    usr := rsp.(*message.Results)
    _ = json.NewEncoder(w).Encode(usr.Results)
  }
}


func getBearer( bearer string ) string  {
  items := strings.Split(bearer, " ")

  if len(items) == 2 {

    auth := items[1]
    return auth
  }
  return ""
}
