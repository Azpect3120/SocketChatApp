package server

import (
	"github.com/Azpect3120/ChatApp/database"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

const auth_app_id string = "90ae7848-7469-4d41-8f30-e37ff96641f3"

type Server struct {
	Router   *gin.Engine
	Config   cors.Config
	Database *database.Database
}

type NewUser struct {
	Username        string `json:"username"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirmPassword"`
}

type User struct {
	ID            string `json:"ID"`
	Username      string `json:"username"`
	Password      string `json:"password"`
	ApplicationID string `json:"applicationID"`
	Data          string `json:"data"`
}

type AuthCreateResponse struct {
	Status int `json:"status"`
	User   User   `json:"user"`
	Error string `json:"error"`
}
