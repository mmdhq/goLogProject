package main

import (
	"github.com/gin-gonic/gin"
	"recreate/webook/internal/web"
)

func main() {
	server := gin.Default()
	u := &web.UserHander{}
	u.RegisterUserRouter(server)
	server.Run(":8880")
}
