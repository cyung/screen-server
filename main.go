package main

import (
  "fmt"
  "log"
  "net/html"
  "github.com/gorilla/mux"
)

func main() {
  router := mux.NewRouter().StrictSlash(true)
  router.HandleFunc("/", index)

  log.Fatal(http.ListenAndServe(":3000", router))
}

func index(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintln(w, "Homepage")
}