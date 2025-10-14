package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/kennedymwaniki/events-api/internal/domain"

	// "github.com/kennedymwaniki/events-api/internal/usecase/event"
	"github.com/kennedymwaniki/events-api/internal/config"
)


type EventHandler struct {
	deps *config.Config
}


func NewEventHandler(deps *config.Config) *EventHandler {
	return &EventHandler{
		deps: deps,
	}
}

func (e *EventHandler) CreateNewEvent(c *gin.Context) {
	var event domain.Event

	if err := c.ShouldBindJSON(&event); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err := e.deps.Event.CreateEvent(c.Request.Context(), &event)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "Event created successfully"})
}


func (e *EventHandler) GetEvents(c *gin.Context) {
	result, err := e.deps.Event.GetAllEvents(c.Request.Context())
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, result)
}


func (e *EventHandler) GetEvent(c *gin.Context) {
	id := c.Param("id")
	event, err := e.deps.Event.GetEvent(c.Request.Context(), id)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, event)
}


func (e *EventHandler) UpdateEvent(c *gin.Context) {
	var event domain.Event

	if err := c.ShouldBindJSON(&event); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	id := c.Param("id")
	if err := e.deps.Event.UpdateEvent(c.Request.Context(), &event, id); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "Event updated successfully"})
}
