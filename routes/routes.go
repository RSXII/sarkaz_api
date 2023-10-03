package routes

import (
	"net/http"
	"sarkaz_api/auth"
	"sarkaz_api/database"
	"sarkaz_api/handlers"

	"github.com/gin-gonic/gin"
)

func InitializeRoutes(r *gin.Engine, characterRepo *database.CharacterRepository) {
    api := r.Group("/")
    {
        api.GET("protected", auth.AuthMiddleware(), func(c *gin.Context) {
            c.JSON(http.StatusOK, gin.H{"message": "This is a protected route"})
        })
        api.GET("characters/:name", handlers.GetCharacterHandler(characterRepo))
        api.GET("characters/simple/:name", handlers.GetSimpleCharacterHandler(characterRepo))
        api.GET("characters/rarity/:stars", handlers.GetCharactersByRarityHandler(characterRepo))
    }
}
