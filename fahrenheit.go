package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if celcius, err := strconv.ParseFloat(os.Args[1], 64); err == nil {
		fmt.Println((celcius * 9 / 5) + 32)
	} else {
		fmt.Println(err)
	}
}
