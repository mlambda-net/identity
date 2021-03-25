package command

import (
  "encoding/json"
  "github.com/mlambda-net/identity/pkg/infrastructure/api/model"
  "net/http"
)


// AuthUser godoc
// @Summary Login the user
// @Produce json
// @Param data body model.Login true "user register"
// @Success 200
// @Failure 500 {string} string "Internal error"
// @Router /login [post]
func (c *control) login(w http.ResponseWriter, r *http.Request) {

  var login model.Login
  _ = json.NewDecoder(r.Body).Decode(&login)

  token, err := c.oauth.PasswordCredentialsToken(r.Context(), "yordivad@gmail.com", "123")
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
  }

  _ = json.NewEncoder(w).Encode(token)


}
