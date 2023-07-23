package main

import (
	"fmt"
	"log"
	"strconv"
)

func normalFunction(menssage string) {
	fmt.Println(menssage)
}

func tripleArgument(a, b int, c string) {
	fmt.Println(a, b, c)

}

func returnValue(a int) int {
	return a * 2
}

func doubleReturn(a int) (c, d int) {
	return a, a * 2
}

func main() {

	normalFunction("Hola Mundo")

	tripleArgument(4, 6, "hola")

	value := returnValue(2)
	fmt.Println("Value: ", value)

	helloMessage := "Hello"
	//worldMessage := "World"

	// Declaracion de constantes
	//Forma - 1
	const pi float64 = 3.14
	//Forma - 2
	const pi2 = 3.1415

	// Declaracion de variables enteras
	base := 12
	var altura int = 14
	var area int

	fmt.Println(base, altura, area)

	// Zero values
	var a int
	var b float64
	var c string
	var d bool

	fmt.Println(a, b, c, d)

	// Area Cuadrado

	const baseCuadrado = 10
	areaCuadrado := baseCuadrado * baseCuadrado
	fmt.Println(areaCuadrado)

	fmt.Println("pi: ", pi)
	fmt.Println("pi2: ", pi2)

	x := 10
	y := 50

	// Suma
	result := x + y
	fmt.Println("Suma: ", result)

	// Resta
	result = y - x
	fmt.Println("Resta: ", result)

	// Multiplicacion
	result = x * y
	fmt.Println("Multipliación: ", result)

	// División
	result = y / x
	fmt.Println("División: ", result)

	// Modulo
	result = y % x
	fmt.Println("Modulo: ", result)

	// Incremental
	x++
	fmt.Println("Incremental: ", x)

	// Decremental
	x--
	fmt.Println("Decremental: ", x)

	// Printf
	nombre := "Platzi"
	cursos := 500
	fmt.Printf("%s tiene más de %d cursos\n", nombre, cursos)
	fmt.Printf("%v tiene más de %v cursos\n", nombre, cursos)

	// Sprintf
	message := fmt.Sprintf("%s tiene mas de %d cursos\n", nombre, cursos)
	fmt.Println(message)

	// tipo de datos

	fmt.Printf("helloMessage: %T \n", helloMessage)
	fmt.Printf("cursos: %T \n", cursos)

	// For condicional
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	fmt.Printf("\n")

	// for decremencal
	for i := 10; i > 0; i-- {
		fmt.Println(i)
	}

	fmt.Printf("\n")
	// For while
	counter := 0
	for counter < 10 {
		fmt.Println(counter)
		counter++
	}

	// for range
	listaNumerosPares := []int{2, 4, 6, 8, 10, 12, 14, 16, 18, 20}
	for i, par := range listaNumerosPares {
		fmt.Printf("posicion %d número par: %d \n", i, par)
	}

	/*
		conterForever := 0
		for {
			fmt.Println(conterForever)
			conterForever++
		}
	*/

	valor1 := 1

	if valor1 == 1 {

		fmt.Println("Es 1")
	} else {
		fmt.Println("No es 1")
	}

	// convertir texto a numero  con strconv.Atoi
	value, err := strconv.Atoi("53")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Vlaue: ", value)

	value4 := 4

	if value4%2 == 0 {
		fmt.Println("Es par")
	} else {
		fmt.Println("Es impar")
	}
	fmt.Printf("\n")

	fmt.Println(validacionUser("Sergio", 123))

}

func validacionUser(name string, password int) bool {

	if name == "Sergio" && password == 123 {
		return true
	} else {
		return false
	}
}
