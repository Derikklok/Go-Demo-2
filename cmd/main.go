package cmd

import (
	"fmt"

)

func Connect() {
	fmt.Println("Read from cmd")
}

func GetUserData() {
	fmt.Println("User data loading...")
}

func GetLoopData() {
	fmt.Println("Loop started....")

	// Classical Loop
	for i := 0; i < 5; i++ {
		fmt.Println("i\t:=\t", i)
	}

	// While loop
	counter := 0
	for counter < 3{
		fmt.Println("Conyer is:\t:\t",counter)
		counter++;
	}
}

func SetName(name string) {
	fmt.Println("First Name\t:\t", name)
}
