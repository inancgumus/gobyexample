package url

import (
	"testing"
	"time"
)

func TestParallelOne(t *testing.T) {
	t.Parallel()
	time.Sleep(5 * time.Second)
}

func TestParallelTwo(t *testing.T) {
	t.Parallel()
	time.Sleep(5 * time.Second)
}

func TestSequential(t *testing.T) {
}
