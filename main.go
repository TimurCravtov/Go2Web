package main

import (
	"go2web/cmd"
	"log/slog"
	"os"
	"time"
	"github.com/lmittmann/tint"
)

func main() {

	logger := slog.New(tint.NewHandler(os.Stdout, &tint.Options{
		Level:      slog.LevelInfo,
		TimeFormat: time.Kitchen,
	}))

	slog.SetDefault(logger)

	cmd.Execute()

	// cache := connect.NewFileCache("cache")

	// cachedGet := cache.WithCache(connect.Get)
	// redirectGet := connect.WithRedirects(cachedGet)

	// fmt.Println(html.ParsePage("https://point.md", redirectGet))

}
