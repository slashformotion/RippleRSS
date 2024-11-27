package main

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"ripplerss/config"
	"ripplerss/db/dbconn"
	"ripplerss/db/gooselogger"
	"ripplerss/db/migrations"
	"ripplerss/db/query"

	_ "github.com/mattn/go-sqlite3"
	"github.com/pressly/goose/v3"
	"go.uber.org/zap"
)

func main() {
	logger, err := zap.NewProduction()
	if err != nil {
		fmt.Println("failed to initialize logger")
		os.Exit(1)
	}
	config, err := config.LoadFromEnvVariables()
	if err != nil {
		logger.Sugar().Errorf("failed to load config: %s", err.Error())
		os.Exit(1)
	}
	logger.Debug("config loaded", zap.Any("config", config))

	db, err := sql.Open("sqlite3", config.DbPath)
	if err != nil {
		logger.Sugar().Fatalf("failed to open database: %w ", err)
	}

	err = dbconn.InitDB(db)
	if err != nil {
		logger.Sugar().Fatalf("failed to apply pragma to db: %w", err)
	}

	// migrations
	// TODO: put this behind a flag
	goose.SetBaseFS(migrations.Embed)
	goose.WithLogger(gooselogger.NewGooseLogger(logger.Sugar()))

	if err := goose.SetDialect("sqlite3"); err != nil {
		panic(err)
	}

	if err := goose.Up(db, "."); err != nil {
		panic(err)
	}

	// fp := gofeed.NewParser()
	// feed, _ := fp.ParseURL("https://lukesmith.xyz/index.xml")
	// fmt.Printf("%v", feed)
	err = dbconn.DeactivateWAL(db)
	if err != nil {
		logger.Fatal("failed to run db exit hooks", zap.Error(err))
	}

}

func insertDummyData(db *sql.DB) {
	ctx := context.Background()
	q := query.New(db)
	q.InsertFeed(ctx, query.InsertFeedParams{
		Description:   nil,
		Link:          nil,
		FeedLink:      nil,
		Updated:       nil,
		UpdatedParsed: sql.NullString{},
		Guid:          "",
	})
}
