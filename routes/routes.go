package routes

import (
	"sarkaz_api/database"
	"sarkaz_api/handlers"

	"github.com/gin-gonic/gin"
)

func InitializeRoutes(r *gin.Engine, characterRepo *database.CharacterRepository) {
    api := r.Group("/api")
    {
        api.GET("/characters/:name", handlers.GetCharacterHandler(characterRepo))
        api.GET("/characters//simple/:name", handlers.GetSimpleCharacterHandler(characterRepo))
        api.GET("/characters/rating/:stars", handlers.GetCharactersByRatingHandler(characterRepo))
    }
}
