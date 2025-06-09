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

var counter int

func incr() { counter++ } // data race

func TestIncr(t *testing.T) {
	t.Parallel()
	t.Run("once", func(t *testing.T) {
		t.Parallel()
		incr()
		if counter != 1 {
			t.Errorf("counter = %d, want 1", counter)
		}
	})

	t.Run("twice", func(t *testing.T) {
		t.Parallel()
		incr()
		incr()
		if counter != 3 {
			t.Errorf("counter = %d, want 3", counter)
		}
	})
}
