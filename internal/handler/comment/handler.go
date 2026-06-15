package comment

import (
	"github.com/Khatchi/go-tweet/internal/middleware"
	"github.com/Khatchi/go-tweet/internal/service/comment"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type Handler struct {
	api            *gin.Engine
	validate       *validator.Validate
	commentService comment.CommentService
}

func NewHandler(api *gin.Engine, validate *validator.Validate, commentService comment.CommentService) *Handler {
	return &Handler{
		api:            api,
		validate:       validate,
		commentService: commentService,
	}
}

func (h *Handler) RouteList(secretKey string) {
	routhAuth := h.api.Group("/comment")
	routhAuth.Use(middleware.AuthMiddleware(secretKey))
	routhAuth.POST("/", h.CreateComment)

}
