package main

import (
	"strings"
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
