package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Khatchi/go-tweet/internal/config"
	"github.com/Khatchi/go-tweet/pkg/internalsql"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	_, err = internalsql.ConnectMySQL(cfg)
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

	server := fmt.Sprintf("127.0.0.1:%s", cfg.Port)
	r.Run(server)
}
