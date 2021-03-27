package handler

import (
  "fmt"
  "github.com/go-session/session"
  "github.com/mlambda-net/identity/pkg/application/message"
  "net/http"
)

func (h *handler) Login(w http.ResponseWriter, r *http.Request) {

  setupHeaders(w)
  store, err := session.Start(r.Context(), w, r)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }

  if r.Method == "POST" {
    if r.Form == nil {
      if err := r.ParseForm(); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
      }
    }

    user := r.Form.Get("username")
    password := r.Form.Get("password")

    resp, err := h.auth.Request(&message.Authenticate{
      Email:    user,
      Password: password,
    }).Unwrap()

    if err != nil {
      w.Header().Set("Location", "/fail")
      w.WriteHeader(http.StatusFound)
      return
    }

    usr := resp.(*message.User)
    store.Set("LoggedInUserID",usr.Id)
    _ = store.Save()
    h.cache.Set(usr.Id, usr)

    w.Header().Set("Location", "/auth")
    w.WriteHeader(http.StatusFound)
    return
  }
  outputHTML(w,  fmt.Sprintf("%s/login.html", h.conf.OAuth.Web))
}
