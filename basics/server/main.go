package main

import "fmt"

type usage float64

func (u usage) high() bool     { return u >= 0.95 }
func (u usage) set(to float64) { u = usage(to) }

// correct set should be:
// func (u usage) set(to float64) usage { return usage(to) }

func main() {
	cpu := usage(0.99)
	fmt.Println("high usage?", cpu.high()) // ..true

	cpu.set(0.7) // cpu is still 0.99
	// correct usage:
	// cpu = cpu.set(0.7)
	fmt.Println("high usage?", cpu.high()) // ..true
}
