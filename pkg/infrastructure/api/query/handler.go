package query

import (
  "encoding/json"
  "github.com/mlambda-net/identity/pkg/infrastructure/api/model"

  "net/http"
)

// Query godoc
// @Summary Query the user
// @Produce json
// @Security ApiKeyAuth
// @Param data body model.Select true "query"
// @Success 200
// @Failure 401 {string} string "Unauthorized"
// @Failure 500 {string} string "Internal error"
// @Router /graphql [post]
func (c *control) handler(w http.ResponseWriter, r *http.Request) {
  token := r.Header.Get("Authorization")
  var s model.Select
  _ = json.NewDecoder(r.Body).Decode(&s)
  result, err := c.exec(s.Query,token)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
  }
  _ = json.NewEncoder(w).Encode(result)
}

