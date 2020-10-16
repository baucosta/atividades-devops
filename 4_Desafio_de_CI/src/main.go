package main

import (
	"fmt"
	"strconv"

	"github.com/logrusorgru/aurora"
)

func soma(val1, val2 int) int {
	return val1 + val2
}

func main() {
	conv := strconv.Itoa(soma(5, 5))

	result := "Resultado de 5 + 5 = " + conv

	fmt.Println(aurora.Magenta(result))
	// fmt.Println(result)
}
