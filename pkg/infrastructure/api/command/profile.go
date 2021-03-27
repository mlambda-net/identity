package command

import (
  "encoding/json"
  "github.com/mlambda-net/identity/pkg/infrastructure/api/model"
  "github.com/mlambda-net/net/pkg/security"
  "net/http"
)

// the profile godoc
// @Summary the profile
// @Produce json
// @Security ApiKeyAuth
// @Success 200  {object} model.Profile
// @Failure 401 {string} string "Unauthorized"
// @Failure 500 {string} string "Internal error"
// @Router /profile [get]
func (c *control) profile(w http.ResponseWriter, r *http.Request) {
  auth := r.Header.Get("Authorization")
  identity, err := security.NewIdentity(auth)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
  } else {
    user := identity.GetUser()
    _ = json.NewEncoder(w).Encode(&model.Profile{
      Email: user.Email,
      Name:  user.Name,
    })
  }
}
