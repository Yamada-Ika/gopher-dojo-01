package main

import (
	"fmt"

	"example.com/ex01/imgconv"
)

func main() {
	if err := imgconv.ConvertImage(); err != nil {
		fmt.Println(err.Error())
	}
}
