package main

import (
	"example.com/gin-project/db"
	"example.com/gin-project/routes"
	"github.com/gin-gonic/gin"
)

func main() {
		db.InitDb()
		server := gin.Default()
		routes.RegisterRoutes(server)

		server.Run(":8080")
}

