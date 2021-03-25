package handler

import (
  "fmt"
  "net/http"
)

func (h *handler) Fail(w http.ResponseWriter, r *http.Request) {
  outputHTML(w,  fmt.Sprintf("%s/fail.html", h.conf.OAuth.Web))
}
