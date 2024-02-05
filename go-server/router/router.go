package router

import (
	"go-server/middleware"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/api/all", middleware.GetAllContests).Methods("GET", "OPTIONS")
	return router
}
