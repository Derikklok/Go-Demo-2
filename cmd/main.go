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


func GetFruitDate(){
	fruits := []string{"apple","banana","mango"}
	fmt.Println("Printing available fruits...\t:\t",fruits)
	fmt.Println("Array length\t:\t",len(fruits))
	// add item to the array
	fruits = append(fruits, "orange")
	fmt.Println("Fruits array after append\t:\t",fruits)
	fmt.Println("New array size\t:\t",len(fruits))
}

// check the functionality of the * marks
 var memAddress int = 457854
 func PrintMemoryAddress(){
	fmt.Println(memAddress)
 }

 
 func PrintNewMemAddress(newMemAdd int){
	fmt.Println(newMemAdd)
	GetCheckRoute()
 }