package main

import (
	"sarkaz_api/database"
	"sarkaz_api/routes"

	"github.com/gin-gonic/gin"
)

func main() {
   
    database.InitDB()

    characterRepo := database.NewCharacterRepository(database.DB)

    r := gin.Default()

    routes.InitializeRoutes(r, characterRepo)

    r.Run(":8080")
}
