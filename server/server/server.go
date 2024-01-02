package server

import (
	"context"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/watariRyo/tasktree/server/config"
	"github.com/watariRyo/tasktree/server/domain/repository"
	"github.com/watariRyo/tasktree/server/handler"
)

type Server struct {
	e       *echo.Echo
	ctx     context.Context
	cfg     *config.Config
	repo    *repository.AllRepository
	handler *handler.HandlerImpl
}

func NewServer(ctx context.Context, cfg *config.Config, handler *handler.HandlerImpl, repo *repository.AllRepository) *Server {
	server := &Server{
		echo.New(),
		ctx,
		cfg,
		repo,
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

	server.e.GET("/api/v1/logout", server.handler.Logout)

	// SignUp / LoginをFirebaseで行うので全て認証必要
	// r := server.e.Group("/echo/api")

	// ルート（認証必要（/api/**））
	// r.GET("/", server.handler.TestMethod)
}

func (server *Server) Run() {
	server.e.Logger.Fatal(server.e.Start(":" + server.cfg.Server.Port))
}
