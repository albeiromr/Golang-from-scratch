package main

import (
	"log"
)

const counter int8 = 1

func main() {

	switch counter {
	case 1:
		//log.Fatal hace que el programa pare después de imprimir un mensaje que nosotros le demos, no
		//da información del error
		log.Fatal("Fatal: esto es un error log.fatal!")
	case 2:
		//log.Panic() hace que el programa pare después de entregar información muy detallada
		// del error como fecha y hora y otra info importante
		log.Panic("Panic: esto es una error de log.Panic!")
	case 3:
		//panic() también da info del error y su ubicación, pero no entrega tanta info como log.Panic()
		panic("Panic: esto es un error de solo panic!!")
	}

}
