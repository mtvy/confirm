package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mtvy/confirm/internal/models"
	"github.com/mtvy/confirm/internal/usecase"
)

// MessageHandler handles HTTP requests related to messages.
type MessageHandler struct {
	usecase *usecase.MessageUsecase
}

func NewMessageHandler(msgUC *usecase.MessageUsecase) *MessageHandler {
	return &MessageHandler{usecase: msgUC}
}

// SendMessage handles the request to send a new message.
func (h *MessageHandler) SendMessage(c *gin.Context) {
	var msg models.Request
	if err := c.ShouldBindJSON(&msg); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id, err := h.usecase.SendMessage(&msg)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to send message"})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"status": "sent for approval", "id": id})
}

// ApproveMessage approves a message by its ID.
func (h *MessageHandler) ApproveMessage(c *gin.Context) {
	id := c.Param("id")

	msg, err := h.usecase.ApproveMessage(id)
	if err != nil {
		if msg != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": err.Error()})
			return
		}
		c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "approved", "message": models.GetResponse(msg)})
}

// RejectMessage rejects a message by its ID.
func (h *MessageHandler) RejectMessage(c *gin.Context) {
	id := c.Param("id")

	msg, err := h.usecase.RejectMessage(id)
	if err != nil {
		if msg != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": err.Error()})
			return
		}
		c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "rejected", "message": models.GetResponse(msg)})
}

func (h *MessageHandler) GetMessage(c *gin.Context) {
	id := c.Param("id")

	msg, err := h.usecase.GetMessage(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": models.GetResponse(msg)})
}
