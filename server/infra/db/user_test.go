package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/watariRyo/tasktree/server/domain/repository"
)

type UsersRepositoryTestSuite struct {
	suite.Suite
	ur   UsersRepository
	conn repository.DBConnection
	ctx  context.Context
}

func (s *UsersRepositoryTestSuite) SetupSuite() {
	s.conn = testConnection(s.T())
	s.ctx = context.Background()
	s.ur = UsersRepository{}
}

func (s *UsersRepositoryTestSuite) TearDownTest() {
	truncateTables(s.T())
}

func TestSuiteUsersRepository(t *testing.T) {
	suite.Run(t, new(UsersRepositoryTestSuite))
}

// func (s *BaseTaskRepositoryTestSuite) Test_BaseTaskRepositoryTestSuite_GetBaseTaskByUserID() {
// 	t := s.T()

// 	userID := 1

// }
