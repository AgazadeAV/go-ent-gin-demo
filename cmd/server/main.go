package main

//go:generate go run -mod=mod entgo.io/ent/cmd/ent generate ./ent/schema

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	"go-ent-gin-demo/ent"
	"go-ent-gin-demo/internal/user"

	"github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "go-ent-gin-demo/docs"
)

// @title Users API
// @version 1.0
// @description This is a simple REST API with Gin and Ent
// @host localhost:8080
// @BasePath /
func main() {
	client, err := ent.Open("sqlite3", "file:ent.db?mode=rwc&cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer client.Close()

	ctx := context.Background()
	if err := client.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	repo := user.NewRepository(client)
	service := user.NewService(repo)
	handler := user.NewHandler(service)

	router := gin.Default()
	handler.RegisterRoutes(router)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run(":8080")
}
