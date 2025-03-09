package main

import (
	"log"
	"net/http"

	"github.com/MaKsY2/shelt/auth/internal/config"
	"github.com/MaKsY2/shelt/auth/internal/controller"
	"github.com/MaKsY2/shelt/auth/internal/repository"
	"github.com/MaKsY2/shelt/auth/internal/service"
	"github.com/gorilla/mux"
)

func main() {
	cfg := config.Load()

	repo, err := repository.NewPostgresRepo(cfg)
	if err != nil {
		log.Fatal(err)
	}

	authService := service.NewAuthService(repo, cfg.JWTSecret)
	authController := controller.NewAuthController(authService)

	r := mux.NewRouter()

	r.HandleFunc("/register", authController.Register).Methods("POST")
	r.HandleFunc("/login", authController.Login).Methods("POST")

	log.Println("Auth service started on port", cfg.ServerPort)
	http.ListenAndServe(":"+cfg.ServerPort, r)
}
