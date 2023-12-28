package web

import "github.com/gin-gonic/gin"

func RegisterRouter(server *gin.Engine) *gin.Engine {
	RegisterUserRouter(server)
	return server
}

func RegisterUserRouter(server *gin.Engine) {
	//由此可见Hander只是一个用于管理方法的东西,它只要出现在需要它的地方就好了
	u := &UserHander{}
	ug := server.Group("/user")
	ug.GET("profile", u.Profile)
	ug.POST("signin", u.SignIn)
	ug.POST("signup", u.SignUp)
	ug.POST("edit", u.Edit)
}
