package usecase

import (
	"context"

	"github.com/brilianpmw/synapsis/presentation"
)

type MockUserUsecase struct {
	dataGetUserDataByUserName  presentation.User
	errorGetUserDataByUserName error
}

func (m *MockUserUsecase) GetUserDataByUserName(ctx context.Context, username string) (presentation.User, error) {
	return m.dataGetUserDataByUserName, m.errorGetUserDataByUserName
}
