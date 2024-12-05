package main

import "fmt"

func main() {
	const usdEuro = 0.95
	const usdRuble = 101.81
	var euroRuble = usdRuble / usdEuro
	fmt.Print(euroRuble)
}
