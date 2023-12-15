package main

import "fmt"

// usage is a type that represents the usage of a server,
// like CPU or memory. It's a float64 between 0 and 1.
type usage float64

// high returns true if the usage is higher than 95%.
func (u usage) high() bool {
	// u acts like a float64 because its underlying type
	// is float64.
	return u >= 0.95
}

// set sets the usage to the given value.
func (u usage) set(to float64) usage {
	// conversion is required because u is usage, not float64.
	return usage(to)
}

// incorrect:
// func (u usage) set(to float64) { u = usage(to) }

// consider this func as the main function for the value receivers
// section of the book. this is called from main() in main.go.
func mainForValueReceivers() {
	cpu := usage(0.99)
	fmt.Println("CPU usage:", cpu)             // prints 0.99
	fmt.Println("high CPU usage?", cpu.high()) // prints ..true
	
	cpu = cpu.set(0.7)
	fmt.Println("CPU usage:", cpu)         // prints 0.7
	fmt.Println("high usage?", cpu.high()) // prints ..false

	// incorrect:
	// cpu.set(0.5)                               // cpu is still 0.95 instead of 0.5
	// fmt.Println("high CPU usage?", cpu.high()) // still prints true
}
