package main

import (
	"log"
	"net/http"

	"github.com/MaKsY2/shelt/src/services/auth/internal/config"
	"github.com/MaKsY2/shelt/src/services/auth/internal/controller"
	"github.com/MaKsY2/shelt/src/services/auth/internal/repository"
	"github.com/MaKsY2/shelt/src/services/auth/internal/service"
)

func main() {
	cfg := config.Load()

	repo, err := repository.NewRepository(cfg)
	if err != nil {
		log.Fatal(err)
	}

	authSvc := service.NewAuthService(repo, cfg.JWTSecret)
	authHandler := handler.NewHandler(authSvc)

	http.HandleFunc("/register", authHandler.Register)
	http.HandleFunc("/login", authHandler.Login)

	log.Printf("Server listening on port %s", cfg.ServerPort)
	http.ListenAndServe(":"+cfg.ServerPort, nil)
}
