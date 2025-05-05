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

func TestQuery(t *testing.T) {
	t.Parallel()
	t.Run("byName", func(t *testing.T) {
		t.Parallel()
		time.Sleep(5 * time.Second)
	})
	t.Run("byInventory", func(t *testing.T) {
		t.Parallel()
		time.Sleep(5 * time.Second)
	})
}
