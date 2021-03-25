package handler

import (
  "io/ioutil"
  "net/http"
)

func outputHTML(w http.ResponseWriter, filename string) {

  file, err := ioutil.ReadFile(filename)
  if err != nil {
    http.Error(w, err.Error(), 500)
    return
  }

  w.Header()
  w.Header().Set("Content-Type", "text/html; charset=utf-8")
  _, _ = w.Write(file)
}
