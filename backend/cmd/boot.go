package app

import (
	"backend/config"
	authHndl "backend/internal/handler"
	"backend/internal/repository"
	authSvc "backend/internal/service"
	"log"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

var (
	userRepo    repository.IUserRepository
	autService  authSvc.IAuth
	authHandler authHndl.Handler
)

func Init() *mux.Router {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	config := config.Init()
	db := config.ConnectDB()

	userRepo = repository.NewRepository(db)

	autService = authSvc.NewService(userRepo)

	authHandler = authHndl.NewHandler(db, autService)

	r := newRouter()
	initAppRouter(r)

	return r
}
