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

type UsersRepository struct{}

var _ repository.UsersRepository = (*UsersRepository)(nil)

func NewUsersRepository() *UsersRepository {
	return &UsersRepository{}
}

func (r *UsersRepository) toDomain(input *models.User) *model.User {
	return &model.User{
		ID:          input.ID,
		UUID:        input.UUID,
		DisplayName: input.DisplayName.Ptr(),
	}
}

func (r *UsersRepository) GetUserByUUID(ctx context.Context, conn repository.DBConnection, uuid string) (*model.User, error) {
	user, err := models.Users(
		models.UserWhere.UUID.EQ(uuid),
	).One(ctx, conn)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Something went wrong when getting user by uuid. uuid=%s", uuid))
	}

	return r.toDomain(user), nil
}

func (r *UsersRepository) Exist(ctx context.Context, conn repository.DBConnection, uuid string) bool {
	exist, err := models.Users(
		models.UserWhere.UUID.EQ(uuid),
	).Exists(ctx, conn)
	if err != nil {
		return false
	}

	return exist
}

func (r *UsersRepository) Insert(ctx context.Context, conn repository.DBConnection, user *model.User) (*model.User, error) {
	u := &models.User{
		UUID:        user.UUID,
		DisplayName: null.StringFrom(*user.DisplayName),
	}
	err := u.Insert(ctx, conn, boil.Infer())
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error inserting user. uuid: %s", user.UUID))
	}
	return r.toDomain(u), nil
}
