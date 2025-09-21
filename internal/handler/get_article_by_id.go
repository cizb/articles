package handler

import (
	"app/internal/entity"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func (h *Handler) GetArticleByID(c *gin.Context) {
	id := c.Param("id")

	article, err := h.useCase.GetArticleByID(&entity.GetArticleByID{
		ID: id,
	})

	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "Not found"})
		return
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": article})
}
