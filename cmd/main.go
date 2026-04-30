package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Khatchi/go-tweet/internal/config"
	userHandler "github.com/Khatchi/go-tweet/internal/handler/user"
	userRepo "github.com/Khatchi/go-tweet/internal/repository/user"
	userService "github.com/Khatchi/go-tweet/internal/service/user"
	"github.com/Khatchi/go-tweet/pkg/internalsql"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func main() {
	r := gin.Default()
	validate := validator.New()

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	db, err := internalsql.ConnectMySQL(cfg)
	if err != nil {
		log.Fatal(err)
	}

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.GET("/check-health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "it works...",
		})
	})

	userRepo := userRepo.NewRepository(db)
	userService := userService.NewService(cfg, userRepo)
	userHandler := userHandler.NewHandler(r, validate, userService)
	userHandler.RouteList()

	server := fmt.Sprintf("127.0.0.1:%s", cfg.Port)
	r.Run(server)
}
