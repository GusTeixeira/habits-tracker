package database

import (
	"context"
	"database/sql"
	"log"
	"os"
	"time"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"github.com/uptrace/bun/extra/bundebug"
)

type QueryHook struct{}

func (h *QueryHook) BeforeQuery(ctx context.Context, event *bun.QueryEvent) context.Context {
	return ctx
}

func (h *QueryHook) AfterQuery(ctx context.Context, event *bun.QueryEvent) {
	log.Println("\n", string(event.Query))
}

func ConnDb() *bun.DB {
	dsn := os.Getenv("DATABASE_URL")
	connector := pgdriver.NewConnector(pgdriver.WithDSN(dsn))
	sqldb := sql.OpenDB(connector)

	sqldb.SetConnMaxLifetime(60 * time.Second)
	sqldb.SetMaxIdleConns(5)
	sqldb.SetMaxOpenConns(10)

	db := bun.NewDB(sqldb, pgdialect.New())
	if _, isDevMode := os.LookupEnv("DEV"); isDevMode {
		db.AddQueryHook(bundebug.NewQueryHook(bundebug.WithVerbose(true)))
	}
	return db
}
