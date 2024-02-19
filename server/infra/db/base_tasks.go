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

type BaseTasksRepository struct{}

var _ repository.BaseTasksRepository = (*BaseTasksRepository)(nil)

func NewBaseTasksRepository() *BaseTasksRepository {
	return &BaseTasksRepository{}
}

func (r *BaseTasksRepository) toDomain(input *models.BaseTask) *model.BaseTask {
	return &model.BaseTask{
		ID:        input.ID,
		UserID:    input.UserID,
		Title:     input.Title.Ptr(),
		Content:   input.Content.Ptr(),
		DayOfWeek: model.DayOfWeekIntToString(input.DayOfWeekID),
	}
}

func (r *BaseTasksRepository) GetBaseTaskByUserID(ctx context.Context, conn repository.DBConnection, userID int) (*model.BaseTask, error) {
	baseTask, err := models.BaseTasks(
		models.BaseTaskWhere.UserID.EQ(userID),
	).One(ctx, conn)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Something went wrong when getting baseTask by userID. userID=%d", userID))
	}

	return r.toDomain(baseTask), nil
}

func (r *BaseTasksRepository) Insert(ctx context.Context, conn repository.DBConnection, baseTask *model.BaseTask) (*model.BaseTask, error) {
	bt := &models.BaseTask{
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
