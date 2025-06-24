package main

import (
	"strings"
	"testing"
)

type testEnv struct {
	stdout strings.Builder
	stderr strings.Builder
}

func testRun(args ...string) (*testEnv, error) {
	var tenv testEnv
	err := run(&env{
		args:   append([]string{"hit"}, args...),
		stdout: &tenv.stdout,
		stderr: &tenv.stderr,
		dryRun: true,
	})
	return &tenv, err
}

func TestRunValidInput(t *testing.T) {
	t.Parallel()

	tenv, err := testRun("https://github.com/inancgumus")
	if err != nil {
		t.Fatalf("got %q;\nwant nil err", err)
	}
	if n := tenv.stdout.Len(); n == 0 {
		t.Errorf("stdout = 0 bytes; want >0")
	}
	if n := tenv.stderr.Len(); n != 0 {
		t.Errorf("stderr = %d bytes; want 0; stderr:\n%s", n, tenv.stderr.String())
	}
}

func TestRunInvalidInput(t *testing.T) {
	t.Parallel()

	tenv, err := testRun("-c=2", "-n=1", "invalid-url")
	if err == nil {
		t.Fatalf("got nil; want err")
	}
	if n := tenv.stderr.Len(); n == 0 {
		t.Error("stderr = 0 bytes; want >0")
	}
}
