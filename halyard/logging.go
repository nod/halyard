package halyard

import (
	"os"

	"golang.org/x/exp/slog"
)


var logger *slog.Logger

func NewLogger(loglevel string) *slog.Logger {
  var lvl slog.Level
  switch loglevel {
  case "warn": lvl = slog.LevelWarn
  case "debug": lvl = slog.LevelDebug
  default: lvl = slog.LevelInfo
  }
  handler := slog.NewJSONHandler(os.Stdout, nil)
  logger := slog.New(handler)
  return logger
}


func GetLogger() *slog.Logger {
  if logger == (slog.Logger{}) {
    logger = NewLogger("")
  }
  return &logger
}
