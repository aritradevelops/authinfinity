package logx

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"os"
	"strings"
)

const (
	colorReset   = "\x1b[0m"
	colorRed     = "\x1b[31m"
	colorGreen   = "\x1b[32m"
	colorYellow  = "\x1b[33m"
	colorBlue    = "\x1b[34m"
	colorMagenta = "\x1b[35m"
	colorCyan    = "\x1b[36m"
	colorWhite   = "\x1b[97m"
)

func colorLevel(l slog.Level) string {
	switch l {
	case slog.LevelDebug:
		return colorMagenta + "DEBUG" + colorReset
	case slog.LevelInfo:
		return colorGreen + "INFO" + colorReset
	case slog.LevelWarn:
		return colorYellow + "WARN" + colorReset
	case slog.LevelError:
		return colorRed + "ERROR" + colorReset
	default:
		return colorBlue + l.String() + colorReset
	}
}

type colorHandler struct {
	w        io.Writer
	minLevel slog.Level
	attrs    []slog.Attr
	group    string
}

func (h *colorHandler) Enabled(_ context.Context, level slog.Level) bool {
	return level >= h.minLevel
}

func (h *colorHandler) Handle(_ context.Context, r slog.Record) error {
	var b strings.Builder
	// level and message
	b.WriteString(colorLevel(r.Level))
	b.WriteString(" ")
	b.WriteString(colorWhite)
	b.WriteString(r.Message)
	b.WriteString(colorReset)

	writeAttrs := func(a slog.Attr) {
		if a.Equal(slog.Attr{}) {
			return
		}
		key := a.Key
		val := a.Value
		// flatten groups
		if h.group != "" {
			key = h.group + "." + key
		}
		b.WriteString(" ")
		b.WriteString(colorCyan)
		b.WriteString(key)
		b.WriteString(colorReset)
		b.WriteString("=")
		// Avoid quoted strings
		if val.Kind() == slog.KindString {
			if s, ok := val.Any().(string); ok {
				b.WriteString(s)
				return
			}
		}
		b.WriteString(fmt.Sprint(val.Any()))
	}
	for _, a := range h.attrs {
		writeAttrs(a)
	}
	r.Attrs(func(a slog.Attr) bool { writeAttrs(a); return true })
	b.WriteByte('\n')
	_, err := io.WriteString(h.w, b.String())
	return err
}

func (h *colorHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	n := *h
	n.attrs = append(append([]slog.Attr{}, h.attrs...), attrs...)
	return &n
}

func (h *colorHandler) WithGroup(name string) slog.Handler {
	n := *h
	if n.group == "" {
		n.group = name
	} else {
		n.group = n.group + "." + name
	}
	return &n
}

// New returns a slog.Logger that prints colored logs without timestamps.
func New() *slog.Logger {
	return slog.New(&colorHandler{w: os.Stderr, minLevel: slog.LevelInfo})
}
