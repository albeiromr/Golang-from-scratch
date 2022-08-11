package main

import (
	"fmt"
)

//arrays y slices

func main() {

	//ARRAYS

	//limitaciones de los arrays

	//You cannot change the size of an array after you have created it

	//When you pass an array to a function, what is happening is that
	//Go creates a copy of that array and passes that copy to that
	//function—therefore any changes you make to an array inside a function
	//are lost when the function returns

	//para definir un arrglos hay que poner su tamaño desde un inicio o usar los tres puntos ...
	testArray1 := [4]int{5, 6, 4, 6}      //se debe poner el tamaño exacto, en este caso 4
	testArray2 := [...]int{5, 6, 4, 6, 9} // go infiere el tamaño del array

	fmt.Printf("la variable testArray1 es de tipo %T \n", testArray1)
	fmt.Printf("la variable testArray2 es de tipo %T \n", testArray2)

	//-----------------------------------------------------------------------------------------

	//SLICES
	//si no se declara el tamaño del array o se usan los (...) automáticamente se crea un slice

	//creando un slice vacio
	emptySlice := []float64{}
	//imprimiendo el array, su tamaño y su capacidad
	fmt.Println(emptySlice, len(emptySlice), cap(emptySlice))

	//agregando elementos a un slice, se debe guardar el valor retornado por el append() en una
	//variable existente o una nueva
	emptySlice = append(emptySlice, 45.23, 26.58)
	emptySlice = append(emptySlice, -34.0)
	fmt.Println(emptySlice, "with length", len(emptySlice))

	//creando un slice con la función make, se debe asignar el valor retornado por make() a una variable
	makeEmptySlice := make([]int, 4)
	//agregando valores
	makeEmptySlice[0] = 21
	makeEmptySlice[1] = 15
	makeEmptySlice[2] = 56
	makeEmptySlice[3] = 589
	makeEmptySlice = append(makeEmptySlice, 111)

	fmt.Println(makeEmptySlice)

	//slices de dos dimenciones, son como dos arreglos hijos anidados en un padre
	twoDimensionsSlide := [][]int{{1, 2, 3}, {4, 5, 6}}
	makeTwoDimensionsSlide := make([][]int, 2)

	//accediendo y asignando valores a un arreglo de dos dimenciones
	fmt.Println(makeTwoDimensionsSlide)
	makeTwoDimensionsSlide[0] = []int{1, 2, 3, 4}
	makeTwoDimensionsSlide[1] = []int{-1, -2, -3, -4}
	fmt.Println(makeTwoDimensionsSlide)

	//accediendo a todos los valores de un slice doble
	for _, i := range twoDimensionsSlide {
		for _, k := range i {
			fmt.Print(k, " ")
		}
		fmt.Println()
	}

	//-------------------------------------------------------------------------

	//SELECCIONANDO PARTES DE UN SLICE
	aSlice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	// seleccionango los primeros 5 elementos
	fmt.Println(aSlice[0:5])
	fmt.Println(aSlice[:5])

	//seleccionango los dos ultimos elementos del slice usando el length del mismo
	aSliceLength := len(aSlice)
	fmt.Println(aSlice[aSliceLength-2 : aSliceLength])

	//-------------------------------------------------------------------------

	//BORRANDO ELEMENTOS DE UN SLICE

	//primer metodo, dividir el slice en 2 partes y luego rearmarlo excluyendo el
	//elemento que queremos borrar
	originalSliceOne := []int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println("Original slice:", originalSliceOne)
	index := 3
	originalSliceOne = append(originalSliceOne[:index], originalSliceOne[index+1:]...)
	fmt.Println("Original slice despues de primer metodo:", originalSliceOne)

	//segundo metodo, reemplazando el elemento que queremos borrar con el ultimo
	originalSliceTwo := []int{1, 2, 3, 4, 5, 6, 7, 8}
	originalSliceTwo[index] = originalSliceTwo[len(originalSliceTwo)-1]
	originalSliceTwo = append(originalSliceTwo[:len(originalSliceTwo)-1])
	fmt.Println("Original slice despues de segundo metodo:", originalSliceTwo)

}
