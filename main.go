package main

import "fmt"


//POINTERS: a continuación un ejemplo

type aStructure struct {
	field1 complex128
	field2 int
}

//recibe un pointer y cambia el valor en memoria del mismo
func processPointer(number *float64) {
	*number = *number * *number
}

//recibe un float64 como parámetro y devuelve un pointer de tipo float64
func returnPointer(number float64) *float64 {
	temp := 2 * number
	return &temp
}

//recibe un pointer de tipo float64 y devuelve un pointer de tipo float64
func bothPointers(number *float64) *float64 {
	temp := 2 * *number
	return &temp
}

func main() {

	var f float64 = 12.123

	//creando un pointer apuntando a f
	fPointer := &f
	fmt.Println("dirección en momoria de f es: ", fPointer)
	fmt.Println("valor en momoria de f es: ", *fPointer)

	//cambiando el valor de f
	processPointer(fPointer)
	fmt.Printf("el nuevo valor de f es: %.2f\n", f)

	//el valor de f no se cambia solo se utiliza para retornar un nuevo pointer
	x := returnPointer(f)
	fmt.Printf("el pointer retornado es : %v\n", x)

	//el valor de f no cambia por que la función solo usa su valor para operaciones internas
	xx := bothPointers(fPointer)
	fmt.Printf("el valor de xx: %.2f\n", *xx)

	//--------------------------------------------------------------------------------------

	//validando si el struct stá vacío
	var k *aStructure
	//el valor de k es nil por que no está apuntando a ningún lado
	fmt.Println(k)
	// por lo tanto se puede hacer esto:
	if k == nil {
		k = new(aStructure)// inicializando valores con zero value
	}
	fmt.Printf("%+v\n", k)
}
