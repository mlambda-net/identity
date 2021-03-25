package command

import (
  "crypto/sha256"
  "encoding/base64"
  "golang.org/x/oauth2"
  "net/http"
)

func (c *control) root(w http.ResponseWriter, r *http.Request) {
  key := sha256.Sum256([]byte("s256example"))
  u := c.oauth.AuthCodeURL("xyz",
    oauth2.SetAuthURLParam("code_challenge", base64.URLEncoding.EncodeToString(key[:])),
    oauth2.SetAuthURLParam("code_challenge_method", "S256"))
  http.Redirect(w, r, u, http.StatusFound)
}
