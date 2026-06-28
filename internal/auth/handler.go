package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *Service
}

// dependency injection 

func NewHandler(service *Service) *Handler {
	return &Handler{
		service: service,
	}
}


func (h *Handler) Login(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message" : "Login Endpoint reached",
	})
}


func (h *Handler )Logout(c *gin.Context){
	c.JSON(http.StatusOK, gin.H{
		"message" : "logout endpoint reached",
	})
}