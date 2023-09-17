// main.go

package main

import (
	"sarkaz_api/database"
	"sarkaz_api/routes"

	"github.com/gin-gonic/gin"
)

func main() {
    database.InitDB()

    r := gin.Default()

    routes.InitializeRoutes(r)

    r.Run(":8080")
}
