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
  slogopts := slog.HandlerOptions{Level: lvl}
  handler := slog.NewJSONHandler(os.Stdout, &slogopts)
  logger := slog.New(handler)
  return logger
}

func GetLogger() *slog.Logger {
  if logger == nil {
    logger = NewLogger("")
  }
  return logger
}

func LogFatal(msg string) {
  GetLogger().Error(msg)
  os.Exit(1)
}

