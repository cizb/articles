package router

import (
	"app/internal/handler"
	"app/internal/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, handler *handler.Handler) {
	r.Use(middleware.CustomHeader())

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "app",
		})
	})

	group := r.Group(handler.Config.HTTP.BaseURL)
	group.POST("/article", handler.CreateArticle)
	group.GET("/article/:id", handler.GetArticleByID)
}
