package handler

import (
	"context"
	"net/http"

	"api-client/repository"
	"github.com/gin-gonic/gin"
)

type ClientHandler struct {
	ClientRepo *repository.ClientRepository
}

func (h *ClientHandler) GetClients(c *gin.Context) {
	clients, err := h.ClientRepo.GetAllClients(context.Background())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, clients)
}

func (h *ClientHandler) GetClientByID(c *gin.Context) {
	clientId := c.Param("id")
	client, err := h.ClientRepo.GetClientByID(context.Background(), clientId)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Client not found"})
		return
	}
	c.JSON(http.StatusOK, client)
}

func (h *ClientHandler) CreateClient(c *gin.Context) {
	var client map[string]interface{}
	if err := c.BindJSON(&client); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	result, err := h.ClientRepo.CreateClient(context.Background(), client)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, result)
}

func (h *ClientHandler) UpdateClient(c *gin.Context) {
	clientId := c.Param("id")
	var update map[string]interface{}
	if err := c.BindJSON(&update); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	result, err := h.ClientRepo.UpdateClient(context.Background(), clientId, update)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, result)
}

func (h *ClientHandler) DeleteClient(c *gin.Context) {
	clientId := c.Param("id")
	result, err := h.ClientRepo.DeleteClient(context.Background(), clientId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, result)
}
