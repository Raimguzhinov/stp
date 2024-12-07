package main

import (
	"fmt"
	"stp/lab5/upnumber"
)

func main() {
	num, _ := upnumber.NewTPNumberFromNumber(10.5, 10, 2)
	fmt.Println("Value:", num.GetValue())
	fmt.Println("String Representation:", num.ToString())

	squared, _ := num.Square()
	fmt.Println("Square:", squared.GetValue())
}
