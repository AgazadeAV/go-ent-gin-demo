package main

//go:generate go run -mod=mod entgo.io/ent/cmd/ent generate ./ent/schema

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	"go-ent-gin-demo/ent"
)

type CreateUserInput struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

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

	router := gin.Default()

	router.POST("/users", func(c *gin.Context) {
		var input CreateUserInput
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		user, err := client.User.
			Create().
			SetName(input.Name).
			SetAge(input.Age).
			Save(ctx)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, user)
	})

	router.GET("/users", func(c *gin.Context) {
		users, err := client.User.Query().All(ctx)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, users)
	})

	router.Run(":8080")
}
