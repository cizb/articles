package handler

import (
	"app/internal/entity"
	"app/pkg/validation"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateArticle(c *gin.Context) {
	var req entity.CreateArticle

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := validation.Validate.Struct(req)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error validating body", "details": err.Error()})
		return
	}

	article, err := h.useCase.CreateArticle(&req)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": article})
}
