package db

import (
	"context"
	"fmt"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/watariRyo/tasktree/server/domain/model"
	"github.com/watariRyo/tasktree/server/domain/repository"
	"github.com/watariRyo/tasktree/server/infra/db/models"
)

type BaseTaskRepository struct{}

var _ repository.BaseTaskRepository = (*BaseTaskRepository)(nil)

func NewBaseTaskRepository() *BaseTaskRepository {
	return &BaseTaskRepository{}
}

func (r *BaseTaskRepository) toDomain(input models.BaseTask) *model.BaseTask {
	return &model.BaseTask{
		ID:        input.ID,
		UserID:    input.UserID,
		Title:     input.Title.Ptr(),
		Content:   input.Content.Ptr(),
		DayOfWeek: model.DayOfWeekIntToString(input.DayOfWeekID),
	}
}

func (r *BaseTaskRepository) GetBaseTaskByUserID(ctx context.Context, conn repository.DBConnection, userID string) (*model.BaseTask, error) {
	return nil, nil
}

func (r *BaseTaskRepository) Insert(ctx context.Context, conn repository.DBConnection, baseTask *model.BaseTask) (*model.BaseTask, error) {
	bt := models.BaseTask{
		UserID:      baseTask.UserID,
		Title:       null.StringFrom(*baseTask.Title),
		Content:     null.StringFrom(*baseTask.Content),
		DayOfWeekID: model.DayOfWeekStringToInt(baseTask.DayOfWeek),
	}
	err := bt.Insert(ctx, conn, boil.Infer())
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error inserting baseTask. userID: %d", baseTask.UserID))
	}
	return r.toDomain(bt), nil
}
