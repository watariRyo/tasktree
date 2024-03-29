package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/watariRyo/tasktree/server/domain/repository"
)

type BaseTasksRepositoryTestSuite struct {
	suite.Suite
	btr  BaseTasksRepository
	conn repository.DBConnection
	ctx  context.Context
}

func (s *BaseTasksRepositoryTestSuite) SetupSuite() {
	s.conn = testConnection(s.T())
	s.ctx = context.Background()
	s.btr = BaseTasksRepository{}
}

func (s *BaseTasksRepositoryTestSuite) TearDownTest() {
	truncateTables(s.T())
}

func TestSuiteBaseTasksRepository(t *testing.T) {
	suite.Run(t, new(BaseTasksRepositoryTestSuite))
}

// func (s *BaseTaskRepositoryTestSuite) Test_BaseTaskRepositoryTestSuite_GetBaseTaskByUserID() {
// 	t := s.T()

// 	userID := 1

// }
