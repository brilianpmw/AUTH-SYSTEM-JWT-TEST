package usecase

import "github.com/brilianpmw/synapsis/presentation"

type MockUserUsecase struct {
	dataGetUserDataByUserName  presentation.User
	errorGetUserDataByUserName error
}
