package log

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"os/user"
	"path/filepath"
)

var Logger *slog.Logger

func Init(filename string) {
	// Get User Home Dir
	usr, err := user.Current()
	if err != nil {
		slog.Error(fmt.Sprintf("Couldn't Get User: %v", err))
	}
	home := usr.HomeDir

	// Log To Standard Dir
	logDir := filepath.Join(home, "logs")
	logFilename := filename
	logFP := filepath.Join(logDir, logFilename)
	configure(logFP) // or pass this in from elsewhere if needed
}

func configure(logFilePath string) {
	// Open log file
	f, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		panic("failed to open log file: " + err.Error())
	}

	// Text handler for console (INFO+)
	consoleHandler := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	})

	// JSON handler for file (DEBUG+)
	fileHandler := slog.NewJSONHandler(f, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	})

	// Combine handlers with split-level control
	Logger = slog.New(NewSplitLevelHandler(consoleHandler, fileHandler))
}

// --- SplitLevelHandler implementation ---

type SplitLevelHandler struct {
	handlers []slog.Handler
}

func NewSplitLevelHandler(handlers ...slog.Handler) *SplitLevelHandler {
	return &SplitLevelHandler{handlers: handlers}
}

func (s *SplitLevelHandler) Enabled(ctx context.Context, level slog.Level) bool {
	for _, h := range s.handlers {
		if h.Enabled(ctx, level) {
			return true
		}
	}
	return false
}

func (s *SplitLevelHandler) Handle(ctx context.Context, record slog.Record) error {
	for _, h := range s.handlers {
		if h.Enabled(ctx, record.Level) {
			if err := h.Handle(ctx, record); err != nil {
				return err
			}
		}
	}
	return nil
}

func (s *SplitLevelHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	newHandlers := make([]slog.Handler, 0, len(s.handlers))
	for _, h := range s.handlers {
		newHandlers = append(newHandlers, h.WithAttrs(attrs))
	}
	return NewSplitLevelHandler(newHandlers...)
}

func (s *SplitLevelHandler) WithGroup(name string) slog.Handler {
	newHandlers := make([]slog.Handler, 0, len(s.handlers))
	for _, h := range s.handlers {
		newHandlers = append(newHandlers, h.WithGroup(name))
	}
	return NewSplitLevelHandler(newHandlers...)
}
