package main

import (
	"fmt"
	"passgen/internal/gen2"
)

func main() {
	gen := gen2.NewGen()

	fmt.Println("New password:", gen.GetNewPass(30))
}
