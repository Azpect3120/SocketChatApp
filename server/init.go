package server

import (
	"encoding/gob"
	"path/filepath"

	"github.com/Azpect3120/ChatApp/database"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func CreateServer(mode string) *Server {
	gin.SetMode(mode)

	server := &Server{
		Router: gin.Default(),
		Config: cors.DefaultConfig(),
		Database: nil,
	}
	return server
}

func (s *Server) Default (databaseConnectionURL string) error {
    s.SetOrigins("*")
    s.SetupPages()
    s.SetupRoutes()
    s.SetupStaticAssets()

	s.SetupDatabase(database.CreateDatabase(databaseConnectionURL))
	return s.Database.Connect()
}

func (s *Server) SetOrigins(origins ...string) {
	s.Config.AllowOrigins = append(s.Config.AllowOrigins, origins...)
	s.Router.Use(cors.New(s.Config))
}

func (s *Server) SetupDatabase (database *database.Database) {
	s.Database = database
}

func (s *Server) SetupPages() {
	path, _ := filepath.Abs(filepath.Join("public", "templates", "*.html"))
	s.Router.LoadHTMLGlob(path)
}

func (s *Server) SetupStaticAssets() {
	path, _ := filepath.Abs(filepath.Join("public", "static"))
	s.Router.Static("/static", path) 
}

func (s *Server) SetupSession (secret string) {
	store := cookie.NewStore([]byte(secret))
	store.Options(sessions.Options{
		MaxAge: 86400,
		SameSite: 1,
	})
	s.Router.Use(sessions.Sessions("mysession", store))
	gob.Register(User{})
}

func (s *Server) SetupRoutes() {
	s.Router.GET("/", RenderIndexPage)
	s.Router.GET("/login", RenderLoginPage)
	s.Router.GET("/signup", RenderSignupPage)
	s.Router.GET("/home", func(ctx *gin.Context) { RenderHomePage(ctx, s.Database)} )

	s.Router.POST("/users", CreateUser)
	s.Router.POST("/login", VerifyUser)
	s.Router.GET("/logout", Logout)

	s.Router.GET("/ws", func(ctx *gin.Context) { OpenSocket(ctx, s.Database) })
}

func (s *Server) Run(port string) error {
	return s.Router.Run(":" + port)
}
