package main

import (
	"fmt"
	"reflect"
)

func main() {

	// Print hello world with GO
	fmt.Println("Hola, Go!")

	/*
		comentario
			multilinea
	*/

	// Variales
	// String
	var varString string = "Contenido de la variable"
	fmt.Printf(varString)

	var varString2 = "Esta es una variable sin indicar el tipo"
	fmt.Println(varString2)

	// Int

	var varInt int = 7
	varInt = varInt + 4
	fmt.Println(varInt)
	fmt.Println(varInt - 1)
	fmt.Println(varInt)

	// Concatenar texto y numero
	fmt.Println(varString, varInt)

	// Ver tipo de datos
	fmt.Println(reflect.TypeOf(varInt))

	// Floar
	var varFloat float64 = 6.5
	fmt.Println(varFloat)
	fmt.Println(reflect.TypeOf(varFloat))
	//suma de float y enteros
	fmt.Println(varInt + int(varFloat))
	fmt.Println(float64(varInt) + varFloat)

	// Boolean
	var varBool bool = false
	varBool = true
	fmt.Println(varBool)

	// Crear variable asignadas automaticamente
	myVar3 := "Esto es una variable autodeclarada"
	fmt.Println(myVar3)

	// Constantes
	const varConstante = "Esto es una constante" //Aúnque una constante no se use, el error permite declaralas y no usarlas
	fmt.Println(varConstante)

	// Control de flujo (if, else, else if)
	if varInt == 10 && !varBool {
		fmt.Println("El valor si es 10")
	} else if varInt == 12 || varBool {
		fmt.Println("El valor es 11")
	} else {
		fmt.Println("el valor no es 10")
	}

	//array
	var myArray [3]int //el numero indica el numero de elemntos
	myArray[0] = 1
	myArray[1] = 2
	myArray[2] = 3
	fmt.Println(myArray)

	var myArray2 [3]string
	myArray2[0] = "hola"
	myArray2[1] = "Mundo"
	myArray2[2] = "desde Go"
	fmt.Println(myArray2[2])

	//map
	var myMap map[string]int
	myMap = make(map[string]int)
	myMap["Miguel"] = 24
	myMap["Daniel"] = 25
	myMap["Javier"] = 43
	fmt.Println("mapa 1", myMap)
	fmt.Println("mapa 1 key Miguel, valor:", myMap["Miguel"])

	myMap2 := make(map[string]int)
	myMap2["Miguel"] = 23
	myMap2["Gonzalo"] = 24
	myMap2["Andres"] = 28
	fmt.Println("mapa 2", myMap2)

	myMap3 := map[string]int{"Andrea": 23, "Maria": 12, "Fonseca": 5}
	fmt.Println("mapa 3", myMap3)


	

}
