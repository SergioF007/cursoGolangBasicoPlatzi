package main

import (
	// pk "curso_golang_platzi/src/mypackage"
	"fmt"
	// "log"
	// "strconv"
	"strings"
)

type figuras2D interface {
	area() float64
}

type cuadrado struct {
	base float64
}

type rectangulo struct {
	base   float64
	altura float64
}

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

	/*
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

	/*
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

		// switch

		switch modulo := 4 % 2; modulo {
		case 0:
			fmt.Println("Es par")
		default:
			fmt.Println("Es impar")
		}

		// switch sin condicion
		value6 := 200
		switch {
		case value6 > 100:
			fmt.Println("Es mayar a 100")
		case value6 < 0:
			fmt.Println("Es menor a 0")
		default:
			fmt.Println("No condicion")
		}

		// uso de los Keywords: defer, break y continue

		// defer
		defer fmt.Println("Mundo")
		fmt.Println("Hola")

		// Continue y break
		for i := 0; i < 10; i++ {
			fmt.Println(i)

			//Continue
			if i == 2 {
				fmt.Println("es 2")
				continue
			}

			// break
			if i == 8 {
				fmt.Println("Break")
				break
			}
		}

		// Array
		// len para mostrar la cantidad de datos en el array
		// cap para mostrar la capacidad que tiene el array
		var array [4]int
		array[0] = 1
		array[1] = 2
		fmt.Println(array, len(array), cap(array))

		// Slice
		slice := []int{0, 1, 2, 3, 4, 5, 6}
		fmt.Println(slice, len(slice), cap(slice))

		// Métodos en el slice
		fmt.Println(slice[0])
		fmt.Println(slice[:3])
		fmt.Println(slice[2:4])
		fmt.Println(slice[4:])

		// Append en los slice
		slice = append(slice, 7)
		fmt.Println(slice)

		// Append nueva lista
		newSlice := []int{8, 9, 10}
		slice = append(slice, newSlice...)
		fmt.Println(slice)

		// Recorrido de Slices con Range.
		slice2 := []string{"Hola", "Que", "Hace"}
		for i := range slice2 {

			fmt.Println(slice2[i])

		}

		// Probando la funcion del palindromo
		isPalindromo("oso")

		m := make(map[string]int)

		m["Jose"] = 14
		m["Sergio"] = 20

		fmt.Println(m)

		// Recorrer el map
		for i, v := range m {
			fmt.Println(i, v)
		}

		// Encontrar un valor
		value8, ok := m["Jose"]
		fmt.Println(value8, ok)

		// instanciar el struc
		myCar := car{brand: "Ford", year: 2020}
		fmt.Println(myCar)

		// otra manera
		var otherCar car
		otherCar.brand = "Ferrari"
		fmt.Println(otherCar)

		var myCar2 pk.CarPublic
		myCar2.Brand = "Ferrari"
		myCar2.Year = 2020
		fmt.Println(myCar2)

		pk.PrintMesseage("Prueba")

		myPC := pc{ram: 16, brand: "msi", disk: 100}
		fmt.Println(myPC)

	*/

	myCudrado := cuadrado{base: 4}
	myRectangulo := rectangulo{base: 4, altura: 2}

	calcular(myCudrado)
	calcular(myRectangulo)

	// Lista de interfaces.
	myInterface := []interface{}{"hola", 12, 4, 90}
	fmt.Println(myInterface...)

	// channels
	// le fefinimos 1 para indicarle que solo tiene un canal
	// por lo que solo va ejecuatar una goroutines a la vez, por lo que la buena practica es idicarle la cantidad limite
	c := make(chan string, 1)

	fmt.Println("Hello")

	go say("Bye", c)
	// le indicamos la salida del canal meidando <-c
	fmt.Println(<-c)

}

type car struct {
	brand string
	year  int
}

func validacionUser(name string, password int) bool {

	if name == "Sergio" && password == 123 {
		return true
	} else {
		return false
	}
}

// Palidromo
func isPalindromo(text string) {

	text = strings.ToLower(text)
	var textReverse string
	// voy a escribir la palabra al reves
	for i := len(text) - 1; i >= 0; i-- {
		textReverse += string(text[i])
	}

	if text == textReverse {
		fmt.Println("Es Palindromo")
	} else {
		fmt.Println("No es Palindromo")
	}

}

// Stringers: Personalizar el output de Strucs
type pc struct {
	ram   int
	brand string
	disk  int
}

// estamos personalizando la funcion de imprimir del struc pc
func (myPC pc) String() string {
	return fmt.Sprintf("Tengo %d GB RAM, %d GB de Disco y es una pc %s", myPC.ram, myPC.disk, myPC.brand)
}

func (c cuadrado) area() float64 {
	return c.base * c.base
}

func (r rectangulo) area() float64 {
	return r.base * r.altura
}

func calcular(f figuras2D) {
	fmt.Println("Area: ", f.area())
}

// channels
// tambien puedo especificar que el parametro del canar va hacer
// o de entrado o de salida, en este caso especificamos
// que va hacer de entrado con chan<- a su lado derecho si fuera de salida seria <-chan
func say(text string, c chan<- string) {
	c <- text
	// si tambine tiene parametros de salid seria
	// text = <- c
}
