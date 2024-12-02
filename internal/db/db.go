package db

import (
	"context"
	"fmt"

	"github.com/akptools/go-services-template/internal/util"
)

type Database interface {
	Disconnect(ctx context.Context) error
	// Add other CRUD Methods
}

func NewDatabase(dbType util.DB, uri string) (Database, error) {
	if dbType == util.MongoDB {
		return NewMongoDB(uri)
	} else if dbType == util.PostgresDB {
		return NewPostgresDB(uri)
	}
	return nil, fmt.Errorf("unrecognized db type")
}
