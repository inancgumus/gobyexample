package link

import (
	"log/slog"
	"math"
	"testing"
)

func TestMain(m *testing.M) {
	// silence the logger
	slog.SetLogLoggerLevel(math.MaxInt)
	m.Run()
}
