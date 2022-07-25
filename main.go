package main

import (
	"log"
	"net/http"
	"os"

	"github.com/miguelanselmo/my-web-app/controllers"
	"github.com/miguelanselmo/my-web-app/repository"
)

func main() {
	log := log.Logger{}
	log.SetOutput(os.Stdout)
	repo := repository.New(&log)
	ctrl := controllers.New(repo, &log)

	http.HandleFunc("/", ctrl.AuthController)
	http.ListenAndServe(":8080", nil)
}
