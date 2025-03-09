package app

import (
	"github.com/gorilla/mux"
)

func newRouter() *mux.Router {
	r := mux.NewRouter()
	r.Use(logging)

	return r
}

func initAppRouter(r *mux.Router) {
	apiGroup := r.PathPrefix("/api").Subrouter()

	apiGroup.HandleFunc("/auth/register", authHandler.Register).Methods("POST")
	apiGroup.HandleFunc("/auth/login", authHandler.Login).Methods("POST")
}
