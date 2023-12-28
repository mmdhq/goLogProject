package web

import (
	"fmt"
	regexp "github.com/dlclark/regexp2"
	"github.com/gin-gonic/gin"
)

type UserHander struct {
	EmailRegexp    *regexp.Regexp
	PasswordRegexp *regexp.Regexp
}

func NewUserHander() *UserHander {
	//常量的定义是在编译的时候就确定的，所以不能在定时使用MustCompile
	const (
		EmailRegexPattern    = "^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\\.[a-zA-Z]{2,}$"
		PasswordRegexPattern = "^(?=.*[a-z])(?=.*[A-Z])(?=.*\\d)(?=.*[^a-zA-Z\\d]).{9,}$"
	)
	EmailRegexp := regexp.MustCompile(EmailRegexPattern, 0)
	PasswordRegexp := regexp.MustCompile(PasswordRegexPattern, 0)
	hander := &UserHander{
		EmailRegexp:    EmailRegexp,
		PasswordRegexp: PasswordRegexp,
	}
	return hander
}

func (u *UserHander) RegisterUserRouter(server *gin.Engine) {
	ug := server.Group("/user")
	ug.GET("/profile", u.Profile)
	ug.POST("/signup", u.SignUp)
	ug.POST("/signin", u.SignIn)
	ug.POST("/edit", u.Edit)
}

func (u *UserHander) SignUp(ctx *gin.Context) {
	type UserInfo struct {
		UserEmail      string `json:"useremail"`
		Password       string `json:"password"`
		VerifyPassword string `json:"verifyPassword"`
	}
	var userinfo UserInfo
	if err := ctx.Bind(&userinfo); err != nil {
		fmt.Println("解析失败", err)
		return
	} else {
		fmt.Printf("%+v\n", userinfo)
	}
	if isMatch, _ := u.EmailRegexp.MatchString(userinfo.UserEmail); !isMatch {
		ctx.String(200, "邮箱格式错误")
		return
	}
	if isMatch, _ := u.PasswordRegexp.MatchString(userinfo.Password); !isMatch {
		ctx.String(200, "密码格式错误")
		return
	}
	if userinfo.Password != userinfo.VerifyPassword {
		ctx.String(200, "两次密码不一致")
		return
	}
	ctx.String(200, "注册成功")

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
