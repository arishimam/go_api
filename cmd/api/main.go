package main

import (
	"fmt"
	"net/http"

	//	"github.com/arishimam/goapi/internal/handlers"
	"github.com/go-chi/chi"
	log "github.com/sirupsen/logrus" // aliased as log
)

func main() {
	log.SetReportCaller(true)
	var r *chi.Mux = chi.NewRouter()
	handlers.Handler(r)

	fmt.Println("Starting GO API service...")

	fmt.Println("UP AND RUNNING")

	err := http.ListenAndServe("localhost:8000", r)
	if err != nil {
		log.Error(err)

	}
}
