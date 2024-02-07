package repository

import (
	"context"

	"github.com/watariRyo/tasktree/server/domain/model"
)

type BaseTaskRepository interface {
	GetBaseTaskByUserID(ctx context.Context, conn DBConnection, userID string) (*model.BaseTask, error)
	Insert(ctx context.Context, conn DBConnection, input *model.BaseTask) (*model.BaseTask, error)
}
