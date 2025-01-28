package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/mtvy/confirm/internal/handler"
	"github.com/mtvy/confirm/internal/repository"
	"github.com/mtvy/confirm/internal/storage"
	"github.com/mtvy/confirm/internal/usecase"
)

func main() {
	r := gin.Default()

	conn, err := storage.Init("postgresql://postgres:postgres@postgres:5432/postgres")
	if err != nil {
		log.Fatal(err)
	}
	storage.RunMigrations(conn, "./migrations")
	if err != nil {
		log.Fatal(err)
	}

	repo := repository.NewMessageRepo(conn)
	uc := usecase.NewMessageUsecase(repo)
	handle := handler.NewMessageHandler(uc)

	r.POST("/send", handle.SendMessage)
	r.POST("/approve/:id", handle.ApproveMessage)
	r.POST("/reject/:id", handle.RejectMessage)
	r.GET("/:id", handle.GetMessage)

	log.Println("Starting server on :8080")
	if err := r.Run("0.0.0.0:8080"); err != nil {
		log.Fatal(err)
	}
}
