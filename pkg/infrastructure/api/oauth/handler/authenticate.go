package handler

import (
  "fmt"
  "github.com/go-session/session"
  "net/http"
)

func (h *handler) Authenticate(w http.ResponseWriter, r *http.Request) {
  store, err := session.Start(r.Context(), w, r)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }

  if _, ok := store.Get("LoggedInUserID"); !ok {
    w.Header().Set("Location", "/login")
    w.WriteHeader(http.StatusFound)
    return
  }

  outputHTML(w, fmt.Sprintf("%s/auth.html", h.conf.OAuth.Web))
}
