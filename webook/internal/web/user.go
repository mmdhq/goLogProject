package web

import "github.com/gin-gonic/gin"

type UserHander struct{}

func (u *UserHander) RegisterUserRouter(server *gin.Engine) {
	ug := server.Group("/user")
	ug.GET("/profile", u.Profile)
	ug.POST("/signup", u.SignUp)
	ug.POST("/signin", u.SignIn)
	ug.POST("/edit", u.Edit)
}

func (u *UserHander) SignUp(ctx *gin.Context) {
	ctx.String(200, "sign up")
}
func (u *UserHander) Edit(ctx *gin.Context) {
	ctx.String(200, "edit")
}
func (u *UserHander) Profile(ctx *gin.Context) {
	ctx.String(200, "profile")
}
func (u *UserHander) SignIn(ctx *gin.Context) {
	ctx.String(200, "sign in")
}
