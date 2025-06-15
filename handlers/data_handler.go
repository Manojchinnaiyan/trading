package handlers

import (
	"net/http"
	"trading-platform-backend/services"

	"github.com/gin-gonic/gin"
)

type DataHandler struct {
	dataService *services.DataService
}

func NewDataHandler(dataService *services.DataService) *DataHandler {
	return &DataHandler{
		dataService: dataService,
	}
}

// GET /holdings
func (h *DataHandler) GetHoldings(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	holdings := h.dataService.GetHoldings(userID.(uint))
	c.JSON(http.StatusOK, holdings)
}

// GET /orderbook
func (h *DataHandler) GetOrderbook(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	orderbook := h.dataService.GetOrderbook(userID.(uint))
	c.JSON(http.StatusOK, orderbook)
}

// GET /positions
func (h *DataHandler) GetPositions(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	positions := h.dataService.GetPositions(userID.(uint))
	c.JSON(http.StatusOK, positions)
}
