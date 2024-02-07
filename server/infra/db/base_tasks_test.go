package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/watariRyo/tasktree/server/domain/repository"
)

type BaseTaskRepositoryTestSuite struct {
	suite.Suite
	btr  BaseTaskRepository
	conn repository.DBConnection
	ctx  context.Context
}

func (s *BaseTaskRepositoryTestSuite) SetupSuite() {
	s.conn = testConnection(s.T())
	s.ctx = context.Background()
	s.btr = BaseTaskRepository{}
}

func (s *BaseTaskRepositoryTestSuite) TearDownTest() {
	truncateTables(s.T())
}

func TestSuiteBaseTaskRepository(t *testing.T) {
	suite.Run(t, new(BaseTaskRepositoryTestSuite))
}

// func (s *BaseTaskRepositoryTestSuite) Test_BaseTaskRepositoryTestSuite_GetBaseTaskByUserID() {
// 	t := s.T()

// 	userID := 1

// }
