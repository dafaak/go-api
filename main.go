package main

import (
	"context"
	"github.com/jmoiron/sqlx"
	"go-api/database"
	"go-api/settings"
)
import "go.uber.org/fx"

func main() {

	app := fx.New(
		fx.Provide(
			context.Background,
			settings.New,
			database.New,
		),
		fx.Invoke(
			func(db *sqlx.DB) {
				_, err := db.Query("select * from USERS")
				if err != nil {
					panic(err)
				}
			}),
	)

	app.Run()

}
