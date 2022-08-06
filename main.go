package main

import (
	"fmt"
	"time"
)

//paralelismo es hacer varias cosas al mismo tiempo
//concurrencia es ocuparse de varias cosas al mismo tiempo

func hello(number int) {
	fmt.Println("Hello number", number)
	//time.sleep es como un setTimeout en javascript
	time.Sleep(5000 * time.Millisecond)
}

func main() {

	//para crear un goroutine usamos la palabra reservada go
	for i := 0; i < 100; i++ {
		go hello(i)
	}

	//las siguientes dos líneas son de rrelleno
	//para que el proceso de la función main no termine antes
	//que la ejecución del goroutine
	var s string
	fmt.Scanln(&s)
}
