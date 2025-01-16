package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/give-names/backend/internal/model"
	"github.com/give-names/backend/internal/service/name"
)

// NameHandler 处理名字生成相关的请求
type NameHandler struct {
	nameService name.Service
}

// NewNameHandler 创建新的名字处理器
func NewNameHandler(nameService name.Service) *NameHandler {
	return &NameHandler{
		nameService: nameService,
	}
}

// RegisterRoutes 注册路由
func (h *NameHandler) RegisterRoutes(r *gin.Engine) {
	api := r.Group("/api")
	{
		api.POST("/generate", h.GenerateNames)
	}
}

// GenerateNames 处理名字生成请求
func (h *NameHandler) GenerateNames(c *gin.Context) {
	var req model.NameRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := h.nameService.GenerateNames(c.Request.Context(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}
