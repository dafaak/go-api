package database

import (
	"context"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"go-api/settings"
)

func New(ctx context.Context, s *settings.Settings) (*sqlx.DB, error) {

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		s.DB.User,
		s.DB.Password,
		s.DB.Host,
		s.DB.Port,
		s.DB.Name,
	)

	db, err := sqlx.ConnectContext(ctx, "mysql", connectionString)

	if err != nil {
		return nil, err
	}

	return db, nil
}
