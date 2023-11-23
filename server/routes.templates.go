package server

import (
	"encoding/json"
	"net/http"
	"fmt"

	"github.com/Azpect3120/ChatApp/database"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func RenderIndexPage(ctx *gin.Context) {
	ctx.Redirect(http.StatusTemporaryRedirect, "/login")
}

func RenderLoginPage(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "login.html", nil)
}

func RenderSignupPage(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "signup.html", nil)
}

type HomeData struct {
	User         User
	MessagesJSON string
}

func RenderHomePage(ctx *gin.Context, db *database.Database) {
	session := sessions.Default(ctx)
	user, ok := session.Get("user").(User)
	if !ok {
		ctx.Redirect(307, "/login")
		return
	}
	messages := db.GetMessages()

	messagesJSON, err := json.Marshal(messages)
	if err != nil {
		fmt.Println(err)
		ctx.Redirect(307, "/login")
		return
	}
	data := &HomeData{User: user, MessagesJSON: string(messagesJSON)}
	ctx.HTML(http.StatusOK, "home.html", data)
}
