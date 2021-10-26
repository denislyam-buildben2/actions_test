package main

import (
	"fmt"
	"os"
)

func main() {


	//eqwe
	fmt.Println("Hello World!", os.Args)
	de, err := os.ReadDir("./")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	for i := range de {
		fmt.Println(de[i].Name())
	}
}
