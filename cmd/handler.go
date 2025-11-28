package cmd

import(
	"fmt"
	"example.com/hello/internal/pkg"
)

func GetCheckRoute(){
	fmt.Println("Hi from Handler")
}

func LoadEnvs(){
	fmt.Println("---------------------")
	pkg.UseRef()
}