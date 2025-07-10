package main

import (
	"compile-service/routes"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	routes.InitRoutes(r)
	log.Println("🚀 compile-service (Gin) listening on :8089")
	r.Run(":8089")
}
