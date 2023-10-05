package main

import (
	"sarkaz_api/database"
	"sarkaz_api/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
   
    database.InitDB()

    db := database.GetDBInstance()
    defer db.Conn.Close()

    r := gin.Default()

    config := cors.DefaultConfig()
    config.AllowOrigins = []string{"http://localhost:4200", "https://www.sarkaz.site"}
    config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}

    r.Use(cors.New(config))

    routes.InitializeRoutes(r, db)

    r.Run(":8080")
}
