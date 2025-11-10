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

	for i := 0; i < 5; i++ {
		fmt.Println("i\t:=\t", i)
	}
}

func SetName(name string) {
	fmt.Println("First Name\t:\t", name)
}
