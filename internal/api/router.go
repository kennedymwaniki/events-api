package api

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/kennedymwaniki/events-api/internal/api/rest"
	"github.com/kennedymwaniki/events-api/internal/config"
)



func NewRouter(deps *config.Config) *gin.Engine {
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "PATCH"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		MaxAge:          12 * time.Hour,
	}))

	handler := rest.NewEventHandler(deps)

	api := r.Group("/api")

	api.GET("/event", handler.GetEvents)
	api.GET("/event/:id", handler.GetEvent)
	api.POST("/event", handler.CreateNewEvent)
	api.PUT("/event/:id", handler.UpdateEvent)


	r.GET("/", func(c *gin.Context) {
		c.JSON(200, "Welcome")
	})
	
	return r
}