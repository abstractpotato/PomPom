package main

import (
  "os"
  "net/http"
  "pom/modules"
  "log"
  _ "embed"
)

//go:embed src/index.html
var html []byte

//go:embed src/style.css
var css []byte

//go:embed src/script.js
var js []byte

//go:embed src/favicon.ico
var fav []byte

//go:embed src/logo.svg
var logo []byte

//go:embed src/ring.wav
var ring []byte

func main() {
  port := ":6060"

  log.Printf("Starting Simple Webserver on port %s", port)
  router := http.NewServeMux()
  router.HandleFunc("/", Show(html, "html"))
  router.HandleFunc("/style.css", Show(css, "css"))
  router.HandleFunc("/script.js", Show(js, "js"))
  router.HandleFunc("/favicon.ico", Show(fav, "ico"))
  router.HandleFunc("/logo.svg", Show(logo, "svg"))
  router.HandleFunc("/ring.wav", Show(ring, "wav"))

  server := http.Server{ Addr: port, Handler: modules.Logging(router), }
  server.ListenAndServe()
}

func Display(file_name, file_type string) http.HandlerFunc {
  return func (w http.ResponseWriter, r *http.Request) {

    switch file_type {
      case "css": w.Header().Set("Content-Type", "text/css")
      case "js":  w.Header().Set("Content-Type", "text/javascript")
    }

    content, err := os.ReadFile(file_name)
    if err != nil {
      w.Write([]byte("Error 500"))
      return
    }
    w.Write(content)
  }
}

func Show(file []byte, file_type string) http.HandlerFunc {
  return func (w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Cache-Control", "max-age=3600")
    switch file_type {
      case "css": w.Header().Set("Content-Type", "text/css")
      case "js":  w.Header().Set("Content-Type", "text/javascript")
    }
    w.Write(file)
  }
}
