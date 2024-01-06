package usecase

import (
	"github.com/watariRyo/tasktree/server/config"
	"github.com/watariRyo/tasktree/server/domain/repository"
	"github.com/watariRyo/tasktree/server/time"
)

type UseCaseImpl struct {
	repo  *repository.AllRepository
	cfg   *config.Config
	clock time.Clock
}

type UseCase interface {
	Logout(ssid string) error
}

var _ UseCase = (*UseCaseImpl)(nil)

func NewUseCase(repo *repository.AllRepository, config *config.Config) *UseCaseImpl {
	clock, err := time.NewClockJST()
	if err != nil {
		panic(err)
	}

	u := &UseCaseImpl{
		repo:  repo,
		cfg:   config,
		clock: clock,
	}

	return u
}
