package usecase

import (
	"context"

	"github.com/watariRyo/tasktree/server/domain/model"
)

func (u *UseCaseImpl) GetBaseTask(ctx context.Context, ssid string) (*model.BaseTask, error) {
	sessionData, err := u.repo.RedisClient.GetSession(ssid)
	if err != nil {
		return nil, err
	}
	baseTask, err := u.repo.BaseTasksRepository.GetBaseTaskByUserID(ctx, u.repo.DBConnection, sessionData.UserID)
	if err != nil {
		return nil, err
	}
	return baseTask, nil
}
