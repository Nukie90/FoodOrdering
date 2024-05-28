package main

import (
	_ "fmt"
	"foodOrder/cmd"
)

func main() {
	cmd.Start("env", "lite", "a string")
}