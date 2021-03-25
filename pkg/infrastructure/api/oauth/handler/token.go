package handler

import "net/http"

func (h *handler) Token(w http.ResponseWriter, r *http.Request) {
  err := h.srv.HandleTokenRequest(w, r)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
  }
}
