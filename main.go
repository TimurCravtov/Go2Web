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
		Level:      slog.LevelDebug,
		TimeFormat: time.Kitchen,
	}))

	slog.SetDefault(logger)

	cmd.Execute()
}
