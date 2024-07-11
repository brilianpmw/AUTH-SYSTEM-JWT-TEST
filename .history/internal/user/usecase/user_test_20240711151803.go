package usecase

import (
	"context"

	"github.com/brilianpmw/synapsis/presentation"
)

type MockUserUsecase struct {
	dataGetUserDataByUserName  presentation.User
	errorGetUserDataByUserName error
}

func (m *MockUserUsecase) GetUserDataByUserName(ctx context.Context, username string) (User, error)
