package main

import (
  "os"
  "net/http"
  "pom/modules"
  "log"
)

const html = "./src/index.html"
const css  = "./src/style.css"
const js   = "./src/script.js"
const fav  = "./src/favicon.ico"
const ring = "./src/ring.wav"

func main() {
  port := ":8000"
  log.Printf("Starting Simple Webserver on port %s", port)
  router := http.NewServeMux()
  router.HandleFunc("/", Display(html, "html"))
  router.HandleFunc("/style.css", Display(css, "css"))
  router.HandleFunc("/script.js", Display(js, "js"))
  router.HandleFunc("/favicon.ico", Display(fav, "ico"))
  router.HandleFunc("/ring.wav", Display(ring, "wav"))
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
