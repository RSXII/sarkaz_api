package routes

import (
	"net/http"
	"sarkaz_api/auth"
	"sarkaz_api/database"
	"sarkaz_api/handlers"

	"github.com/gin-gonic/gin"
)

func InitializeRoutes(r *gin.Engine, db *database.Database) {
    userRepository := db.GetUserRepository()
    characterRepository := db.GetCharacterRepository()


    api := r.Group("/")
    {
        api.GET("protected", auth.AuthMiddleware(), func(c *gin.Context) {
            c.JSON(http.StatusOK, gin.H{"message": "This is a protected route"})
        })
        api.POST("auth/signup", handlers.SignUpHandler(userRepository))
        api.GET("characters/:name", handlers.GetCharacterHandler(characterRepository))
        api.GET("characters/simple/:name", handlers.GetSimpleCharacterHandler(characterRepository))
        api.GET("characters/rarity/:stars", handlers.GetCharactersByRarityHandler(characterRepository))
    }
}
