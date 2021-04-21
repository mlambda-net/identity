package handler

import (
  "github.com/go-session/session"
  "net/http"
  "net/url"
)

func (h *handler) Authorize(w http.ResponseWriter, r *http.Request) {
  setupHeaders(w,r)

  store, err := session.Start(r.Context(), w, r)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }

  var form url.Values
  if v, ok := store.Get("ReturnUri"); ok {
    form = v.(url.Values)
  }
  r.Form = form

    store.Delete("ReturnUri")
    store.Save()

    err = h.srv.HandleAuthorizeRequest(w, r)
    if err != nil {
      http.Error(w, err.Error(), http.StatusBadRequest)
    }
  }


