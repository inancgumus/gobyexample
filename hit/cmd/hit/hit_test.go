package main

import "strings"

type testEnv struct {
	stdout strings.Builder
	stderr strings.Builder
}
