package main

import "strconv"

type parseFunc func(string) error

func stringVar(p *string) parseFunc {
	return func(s string) error {
		*p = s
		return nil
	}
}

func intVar(p *int) parseFunc {
	return func(s string) error {
		var err error
		*p, err = strconv.Atoi(s)
		return err
	}
}
