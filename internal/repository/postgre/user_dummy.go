package postgre

import (
	"context"

	"github.com/brilianpmw/synapsis/presentation"
)

func (db *Postgre) GetUserDataByUserName(ctx context.Context, username string) (presentation.User, error) {

	return presentation.User{
		Username: "Brilian",
		Gender:   "Man",
		Password: "hashedpw",
	}, nil
}
