package main

import (
	"log"
	"os"

	"github.com/Azpect3120/ChatApp/server"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

const session_secret string = "890rhjguiqnt89wbghjbt891ngoi89tgb89thb12398hqg89ahbtguillb"

func main () {
    // Update to gin.ReleaseMode for prod
    server := server.CreateServer(gin.DebugMode)

    // Remove for prod
    if err := godotenv.Load(); err != nil {
        log.Fatalln(err)
    }

    server.SetupSession(session_secret)

    // Hardcode url for prod
    if err := server.Default(os.Getenv("db_url")); err != nil {
        log.Fatalln(err)
    }

    server.Run("3000")
}
