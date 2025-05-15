package main

import "fmt"

type usage int

func (u usage) high() bool { return u >= 95 }
func (u usage) set(to int) { u = usage(to) }

func main() {
	var cpu usage = 99
	cpu.set(70)
	fmt.Printf("cpu: %d%% high:%t\n", cpu, cpu.high())
}
