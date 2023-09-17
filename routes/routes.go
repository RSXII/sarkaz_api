package routes

import (
	"sarkaz_api/handlers"

	"github.com/gin-gonic/gin"
)

func InitializeRoutes(r *gin.Engine) {
    api := r.Group("/api")
    {
        api.GET("/characters/:name", handlers.GetCharacterHandler)
        api.GET("/characters//simple/:name", handlers.GetSimpleCharacterHandler)
    }
}
