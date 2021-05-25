package main

import "fmt"

type formatter string

func (g formatter) Format(in string) {
	fmt.Println(in)
}

var Formatter formatter
