package server

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/sessions"
)

// Create a new user in the application
func CreateUser(ctx *gin.Context) {
	var newUser NewUser

	err := ctx.ShouldBindJSON(&newUser)
	if err != nil {
		ctx.JSON(200, gin.H{ "status": 400, "error": err.Error() })
		return
	}

	if newUser.Password != newUser.ConfirmPassword {
		ctx.JSON(200, gin.H{ "status": 400, "error": "Passwords must match" })
		return
	}

	payload := []byte(fmt.Sprintf(`{ "applicationID": "%s", "username": "%s", "password": "%s", "data": "{}" }`, auth_app_id, newUser.Username, newUser.Password))

	response, err := http.Post("http://54.176.161.136:8080/users/create", "application/json", bytes.NewBuffer(payload))
	if err != nil {
		ctx.JSON(200, gin.H{ "status": 500, "error": err.Error() })
		return
	}
	
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		ctx.JSON(200, gin.H{ "status": 500, "error": err.Error() })
		return
	}

	var authResponse AuthCreateResponse

	err = json.Unmarshal(body, &authResponse)
	if err != nil {
		ctx.JSON(200, gin.H{ "status": 500, "error": err.Error() })
		return
	}

	if authResponse.Status != 201 {
		ctx.JSON(200, gin.H{ "status": authResponse.Status, "error": authResponse.Error })
		return
	}

	session := sessions.Default(ctx)
	session.Set("user", authResponse.User)
	session.Save()

	ctx.Header("HX-Redirect", "/home/100000")
	ctx.JSON(200, gin.H{ "status": 200 })
}

// Log a user into the application
func VerifyUser (ctx *gin.Context) {
	var user User

	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		ctx.JSON(200, gin.H{ "status": 400, "error": err.Error() })
		return
	}

	payload := []byte(fmt.Sprintf(`{ "applicationID": "%s", "username": "%s", "password": "%s" }`, auth_app_id, user.Username, user.Password))

	response, err := http.Post("http://54.176.161.136:8080/users/verify", "application/json", bytes.NewBuffer(payload))
	if err != nil {
		ctx.JSON(200, gin.H{ "status": 500, "error": err.Error() })
		return
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		ctx.JSON(200, gin.H{ "status": 500, "error": err.Error() })
		return
	}

	var authResponse AuthCreateResponse
	err = json.Unmarshal(body, &authResponse)
	if err != nil {
		ctx.JSON(200, gin.H{ "status": 500, "error": err.Error() })
		return
	}

	if authResponse.Status != 200 {
		ctx.JSON(200, gin.H{ "status": authResponse.Status, "error": authResponse.Error })
		return
	}

	session := sessions.Default(ctx)
	session.Set("user", authResponse.User)
	session.Save()

	ctx.Header("HX-Redirect", "/home/100000")
	ctx.JSON(200, gin.H{ "status": 200 })
}

// Logs a user out
// Sets user key in session to nil
func Logout (ctx *gin.Context) {
	session := sessions.Default(ctx)
	session.Set("user", nil)
	session.Save()

	ctx.Header("HX-Redirect", "/login")
	ctx.JSON(200, gin.H{ "status": 200 })
}
