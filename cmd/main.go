package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Khatchi/go-tweet/internal/config"
	commentHandler "github.com/Khatchi/go-tweet/internal/handler/comment"
	postHandler "github.com/Khatchi/go-tweet/internal/handler/post"
	userHandler "github.com/Khatchi/go-tweet/internal/handler/user"
	commentRepo "github.com/Khatchi/go-tweet/internal/repository/comment"
	postRepo "github.com/Khatchi/go-tweet/internal/repository/post"
	userRepo "github.com/Khatchi/go-tweet/internal/repository/user"
	commentService "github.com/Khatchi/go-tweet/internal/service/comment"
	postService "github.com/Khatchi/go-tweet/internal/service/post"
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
	postRepo := postRepo.NewPostRepository(db)
	commentRepo := commentRepo.NewRepostory(db)

	userService := userService.NewService(cfg, userRepo)
	postService := postService.NewPostService(cfg, postRepo)
	commentService := commentService.NewCommentService(cfg, commentRepo, postRepo)

	userHandler := userHandler.NewHandler(r, validate, userService)
	postHandler := postHandler.NewHandler(r, validate, postService)
	commentHandler := commentHandler.NewHandler(r, validate, commentService)

	userHandler.RouteList(cfg.SecretJwt)
	postHandler.RouteList(cfg.SecretJwt)
	commentHandler.RouteList(cfg.SecretJwt)

	server := fmt.Sprintf("127.0.0.1:%s", cfg.Port)
	r.Run(server)
}
