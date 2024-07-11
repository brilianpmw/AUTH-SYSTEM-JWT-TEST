package usecase

import "github.com/brilianpmw/linknau/presentation"

type Repositories struct {
	User presentation.IUser
}

type Usecase struct {
	repository *Repositories
}

func New(repository *Repositories) *Usecase {
	uc := &Usecase{
		repository: repository,
	}
	return uc
}
