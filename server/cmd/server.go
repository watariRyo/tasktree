package main

import (
	"context"

	"github.com/watariRyo/tasktree/server/config"
	"github.com/watariRyo/tasktree/server/domain/repository"
	"github.com/watariRyo/tasktree/server/handler"
	"github.com/watariRyo/tasktree/server/infra/db"
	"github.com/watariRyo/tasktree/server/infra/firebase"
	"github.com/watariRyo/tasktree/server/infra/redis"
	"github.com/watariRyo/tasktree/server/server"
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

	ctx := context.Background()
	firebaseApp, err := firebase.InitFirebaseApp(ctx, cfg.Server)

	allRepository := &repository.AllRepository{
		DBConnection:       conn,
		DBTransaction:      db.Transaction,
		RedisClient:        redisClient,
		FirebaseApp:        firebaseApp,
		BaseTaskRepository: db.NewBaseTaskRepository(),
	}

	usecase := usecase.NewUseCase(allRepository, cfg)
	handler := handler.NewHandler(usecase)

	server := server.NewServer(ctx, cfg, handler, allRepository)
	server.Run()
}
