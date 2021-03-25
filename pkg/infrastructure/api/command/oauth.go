package command

import (
  "encoding/json"
  "golang.org/x/oauth2"
  "net/http"
)

func (c *control) OAuth2 (w http.ResponseWriter, r *http.Request) {
  _ = r.ParseForm()
  state := r.Form.Get("state")
  if state != "xyz" {
    http.Error(w, "State invalid", http.StatusBadRequest)
    return
  }
  code := r.Form.Get("code")
  if code == "" {
    http.Error(w, "Code not found", http.StatusBadRequest)
    return
  }
  token, err := c.oauth.Exchange(r.Context(), code, oauth2.SetAuthURLParam("code_verifier", "s256example"))
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }

  e := json.NewEncoder(w)
  e.SetIndent("", "  ")
  e.Encode(token)
}
