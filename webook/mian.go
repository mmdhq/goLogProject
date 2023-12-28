package main

import (
	"github.com/gin-gonic/gin"
	"recreate/webook/internal/web"
)

func main() {
	server := gin.Default()
	server = web.RegisterRouter(server)
	err := server.Run(":8880")
	if err != nil {
		return
	}
}
