package main

import (
	"fmt"

	"example.com/hello/cmd"
)

// build command - go build -o ./build/main.exe
func main() {
	fmt.Println("Hello From Go...!")
	var name string = "Sachin"
	fmt.Println(name)

	pageCount := 499
	fmt.Println(pageCount)

	pageCount = 500
	fmt.Println(pageCount)

	age := 18

	if age >= 18 {
		fmt.Println("You are old enough to vote", age)
	} else {
		fmt.Println("You need to be at least 18 or above to vote.")
	}

	// use other go files
	cmd.Connect()
	cmd.GetUserData()
	cmd.GetLoopData()
	cmd.SetName("Sachin")

	// Get Data from Slices - basic collections in Go
	cmd.GetFruitDate()
	cmd.PrintMemoryAddress()
	cmd.PrintNewMemAddress(12345)
}

// Double shift to switch between multiple files
