package main

import (
	"fmt"
)


func main() {

	//forma cencilla de obtener el input del usuario desde la misma terminal
	var name string
	var age int8
	fmt.Printf("introduce your name:  ")
	fmt.Scanln(&name)
	fmt.Printf("introduce your age:  ")
	fmt.Scanln(&age)

	switch {
	case age < 18:
		fmt.Printf("%v no eres mayor de 18 años, lo sentimos \n", name)
	case age > 18:
		fmt.Printf("%v eres mayor de 18 años, puedes acceder al servicio \n", name)
	default:
		fmt.Println("no hemos podido ver tu edad")
	}
}
