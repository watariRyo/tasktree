package main

import (
	"github.com/labstack/echo/v4"
	"github.com/watariRyo/tasktree/server/config"
	"github.com/watariRyo/tasktree/server/domain/repository"
	"github.com/watariRyo/tasktree/server/handler"
	"github.com/watariRyo/tasktree/server/infra/db"
	"github.com/watariRyo/tasktree/server/infra/redis"
	"github.com/watariRyo/tasktree/server/usecase"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		panic(err)
	}
	println(cfg.Db.Host)

	// infra
	dbConnectionManager := db.NewConnectionManager(
		db.Username(cfg.Db.Username),
		db.Password(cfg.Db.Password),
		db.Host(cfg.Db.Host),
		db.Port(cfg.Db.Port),
		db.Schema(cfg.Db.Schema),
		db.DebugMode(cfg.Db.DebugMode),
	)
	conn, err := dbConnectionManager.Open()
	if err != nil {
		panic(err)
	}

	redisClient, err := redis.NewRedisClient(&cfg.Redis)
	if err != nil {
		panic(err)
	}

	allRepository := &repository.AllRepository{
		DBConnection:  conn,
		DBTransaction: db.Transaction,
		RedisClient:   redisClient,
	}

	usecase := usecase.NewUseCase(allRepository, cfg)
	handler := handler.NewHandler(usecase)

	e := echo.New()
	e.GET("/test", handler.TestMethod)
	e.Logger.Fatal(e.Start(":" + cfg.Server.Port))
}
