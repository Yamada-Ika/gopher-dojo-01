package main

import (
	"fmt"
	"os"

	"example.com/ex01/imgconv"
)

func main() {
	if err := imgconv.ConvertImage(); err != nil {
		fmt.Fprint(os.Stderr, err.Error())
		os.Exit(1)
	}
}
