package main

import (
	"errors"
	"fmt"
	"strconv"
)

//devolviendo un error personalizado con  errors.new()
func checkInt(a,b int) error {
	if a == 0 && b == 0 {
		return errors.New("Error personalizado: los valores no pueden estar en cero")
	}
	return nil
}

//devolviendo un error formateado con fmt.Errorf()
func formatedError(a int) error {
	if a == 0 {
		return fmt.Errorf("Error formateado: no puedes inicializar con un paremetro en %v", a)
	}
	return nil
}


func main() {

	personalizedError := checkInt(0,0)
	if personalizedError != nil {
		fmt.Println(personalizedError )
	} else {
		fmt.Println("no existen errores")
	}

	formatedError := formatedError(0)
	if formatedError  != nil {
		fmt.Println(formatedError)
	} else {
		fmt.Println("no existen errores")
	}

	//convirtiendo un string a un número, si no se puese convertir arroja un error
	integer, parseError := strconv.Atoi("1234")
	if parseError != nil {
		fmt.Println(parseError)
	} else {
		fmt.Println( "The integer is ", integer)
	}
}
