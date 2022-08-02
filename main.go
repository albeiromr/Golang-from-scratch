package main

import "fmt"

//variable global, se declara así por que
//que sea una variable global o por qué queremos
//definir claramente de que tipo debe ser,
//es decir que golang no infiera el tipado
var age int8 = 30
//las variables con var no necesariamente deben ser
//inicializadas
var money int16
//las variables con var podrían no llevar el tipo de dato
var children = 0

func main() {

	//las variables con var también pueden declararse dentro 
	//de funciones, esto por ejemplo para evitar inferir el 
	//tipado
	var name string = "Albeiro"

	//tambien se pueden declarar variables sin la palabra var
	//estas deben declararse con la notación :=, la cual es 
	//llamada oficialmente short assignment statement, estas
	//estas funciones no pueden ser declaradas de forma global
	//como se hace con las de var
	id := 236563


	fmt.Println("Hello", name, "your age is", age, "and your id is", id)
}
