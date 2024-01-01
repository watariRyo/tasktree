package server

import (
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/watariRyo/tasktree/server/config"
	"github.com/watariRyo/tasktree/server/handler"
)

type Server struct {
	e       *echo.Echo
	cfg     *config.Config
	handler *handler.HandlerImpl
}

func NewServer(cfg *config.Config, handler *handler.HandlerImpl) *Server {
	server := &Server{
		echo.New(),
		cfg,
		handler,
	}

	// Routingとmiddleware
	server.setRouting()
	server.setCors()
	server.setFirebase()

	return server
}

func (server *Server) setRouting() {
	// middleware設定
	server.e.Use(middleware.Logger())
	server.e.Use(middleware.Recover())
	server.e.Use(middleware.Gzip())

	// セッション
	server.e.Use(session.Middleware(sessions.NewCookieStore([]byte(server.cfg.Jwt.Key))))

	// ルート（認証不要）
	// server.e.POST("/echo/signUp", server.signUpHandler.SignUp)

	// ルート（認証必要）
	// r := server.e.Group("/echo/api")

	// ルート（認証必要（/api/**））
	// r.GET("/", server.handler.TestMethod)
}

func (server *Server) Run() {
	server.e.Logger.Fatal(server.e.Start(":" + server.cfg.Server.Port))
}
