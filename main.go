package main

//IMPORTACIONES
//para importar un paquete externo usamos el prefijo con el cual fue creado el proyecto
//el cual se encuentra en el go.mod file y luego agregamos la ruta del paquete
import (
	"fmt"
	"primeros-pasos/hello"
)

func main() {
	fmt.Println("Hello world desde el main")
	hello.HelloWorld()
}
