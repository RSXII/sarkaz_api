package main

import (
	"sarkaz_api/database"
	"sarkaz_api/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
   
    database.InitDB()

    characterRepo := database.NewCharacterRepository(database.DB)

    r := gin.Default()

    config := cors.DefaultConfig()
    config.AllowOrigins = []string{"http://localhost:4200"}
    config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}

    r.Use(cors.New(config))

    routes.InitializeRoutes(r, characterRepo)

    r.Run(":8080")
}
