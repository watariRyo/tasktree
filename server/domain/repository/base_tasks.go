package repository

import (
	"context"

	"github.com/watariRyo/tasktree/server/domain/model"
)

type BaseTasksRepository interface {
	GetBaseTaskByUserID(ctx context.Context, conn DBConnection, userID int) (*model.BaseTask, error)
	Insert(ctx context.Context, conn DBConnection, input *model.BaseTask) (*model.BaseTask, error)
}
