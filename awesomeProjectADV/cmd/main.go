package main

import (
	"awesomeProjectADV/internal/config"
	"awesomeProjectADV/internal/handler"
	repo "awesomeProjectADV/internal/repository/mongo"
	"awesomeProjectADV/internal/usecase"

	"github.com/gin-gonic/gin"
)

func main() {
	db := config.ConnectMongo("mongodb://localhost:27017")

	// Movie Service setup
	movieRepo := repo.NewMovieRepository(db)
	movieUC := usecase.NewMovieUsecase(movieRepo)
	movieHandler := handler.NewMovieHandler(movieUC)

	// Session Service setup
	sessionRepo := repo.NewSessionRepository(db)
	sessionUC := usecase.NewSessionUsecase(sessionRepo)
	sessionHandler := handler.NewSessionHandler(sessionUC)

	r := gin.Default()
	movieHandler.RegisterRoutes(r)
	sessionHandler.RegisterRoutes(r)

	// отдавать index.html
	r.StaticFile("/", "./frontend/index.html")

	// отдавать остальные файлы (js/css)
	r.Static("/static", "./frontend")

	r.Run(":8081")
}
