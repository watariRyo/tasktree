package redis

import (
	"encoding/json"
	"time"

	"github.com/go-redis/redis"
	"github.com/watariRyo/tasktree/server/config"
	"github.com/watariRyo/tasktree/server/domain/model"
	"github.com/watariRyo/tasktree/server/domain/repository"
)

type RedisClient struct {
	client *redis.Client
}

var _ repository.RedisClient = (*RedisClient)(nil)

func NewRedisClient(cfg *config.Redis) (*RedisClient, error) {
	client := redis.NewClient(
		&redis.Options{
			Addr:     cfg.Host + ":" + cfg.Port,
			Password: cfg.Password,
		},
	)
	_, err := client.Ping().Result()
	if err != nil {
		return nil, err
	}

	return &RedisClient{
		client: client,
	}, nil
}

func (rc *RedisClient) SaveSession(sessionID string, sessionData model.SessionData, duration time.Duration) error {
	seralized, err := rc.sessionDataToJSON(&sessionData)
	if err != nil {
		return err
	}
	err = rc.client.Set(sessionID, seralized, duration).Err()
	if err != nil {
		return err
	}

	return nil
}

func (rc *RedisClient) GetSession(sessionID string) (*model.SessionData, error) {
	// sessionData, err := rc.client.Get(sessionID).Result()
	sessionData, err := rc.jsonToSessionData(sessionID)
	if err != nil {
		return nil, err
	}

	return sessionData, nil
}

func (rc *RedisClient) DeleteSession(sessionID string) error {
	err := rc.client.Del(sessionID).Err()
	if err != nil {
		return err
	}

	return nil
}

func (rc *RedisClient) sessionDataToJSON(sessionData *model.SessionData) ([]byte, error) {
	seralized, err := json.Marshal(sessionData)
	if err != nil {
		return nil, err
	}
	return seralized, nil
}

func (rc *RedisClient) jsonToSessionData(sessionID string) (*model.SessionData, error) {
	data, err := rc.client.Get(sessionID).Result()
	if err != nil {
		return nil, err
	}

	sessionData := new(model.SessionData)
	err = json.Unmarshal([]byte(data), sessionData)
	if err != nil {
		return nil, err
	}
	return sessionData, nil
}
