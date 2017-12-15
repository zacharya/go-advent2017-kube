package main

import (
  "net/http"
  log "github.com/sirupsen/logrus"
  "github.com/zacharya/go-advent2017-kube/handlers"
  "os"
)

func main() {
  log.Print("Starting the service...")

  port := os.Getenv("PORT")
  if port == "" {
    log.Fatal("Port is not set.")
  }

  r := handlers.Router()

  log.Print("The service is ready to listen and serve.")
	log.Fatal(http.ListenAndServe(":"+port, r))
}
