package server

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func RenderIndexPage (ctx *gin.Context) {
	ctx.Redirect(http.StatusTemporaryRedirect, "/login")
}

func RenderLoginPage (ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "login.html", nil)
}

func RenderSignupPage (ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "signup.html", nil)
}

type HomeData struct {
	User User
}

func RenderHomePage (ctx *gin.Context) {
	session := sessions.Default(ctx)
	user, ok := session.Get("user").(User)
	if !ok {
		ctx.Redirect(307, "/login")
		return
	}
	data := &HomeData{ User: user }
	ctx.HTML(http.StatusOK, "home.html", data)
}
