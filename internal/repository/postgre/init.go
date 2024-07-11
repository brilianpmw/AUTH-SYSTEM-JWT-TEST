package postgre

import (
	"github.com/brilianpmw/linknau/internal/pkg/database"
)

// Postgre : Main struct of this repository
type Postgre struct {
	Database *database.Database
}

func New() (*Postgre, error) {

	db, err := database.ConnectDB()

	if err != nil {
		return nil, err
	}

	// Return database struct object
	return &Postgre{
		Database: db,
	}, nil
}
