package main

import (
	"fmt"
	"os"
	"ripplerss/config"

	"github.com/mmcdole/gofeed"
	"go.uber.org/zap"
)

func main() {
	logger, err := zap.NewDevelopment()
	if err != nil {
		fmt.Println("failed to initialize logger")
		os.Exit(1)
	}
	config, err := config.LoadFromEnvVariables()
	if err != nil {
		logger.Sugar().Errorf("failed to load config: %s", err.Error())
		os.Exit(1)
	}
	fmt.Println(config)

	fp := gofeed.NewParser()
	feed, _ := fp.ParseURL("https://lukesmith.xyz/index.xml")
	fmt.Printf("%v", feed)
}
