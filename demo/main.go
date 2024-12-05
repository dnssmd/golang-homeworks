package main

import (
	"fmt"
	"strconv"
)

func main() {
	var a = 3
	var b = 5
	var summ = a + b
	var result = "Сумма: " + strconv.Itoa(summ)
	fmt.Print(result)
}
