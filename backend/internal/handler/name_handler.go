package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/xmualex2023/give-name/internal/ai"
	"github.com/xmualex2023/give-name/internal/service"
)

type NameHandler struct {
	nameService   *service.NameService
	geminiService *ai.GeminiService
}

func NewNameHandler(nameService *service.NameService, geminiService *ai.GeminiService) *NameHandler {
	return &NameHandler{
		nameService:   nameService,
		geminiService: geminiService,
	}
}

func (h *NameHandler) GenerateName(c *gin.Context) {
	var req struct {
		FirstName string `json:"firstName"`
		LastName  string `json:"lastName"`
	}

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// 调用名字生成服务
	names := h.nameService.Generate(req.FirstName, req.LastName)

	c.JSON(http.StatusOK, gin.H{
		"names": names,
	})
} 