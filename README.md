# cursoGolangBasicoPlatzi
Variables, Constantes y Zero values -  Tipos de datos primitivos- Operadores aritméticos - cilo for - switch - keywords defer, break y continue-    Manejao de errores - Arrays y Slice - Slice de interfaces - map - punteros -  structs 


# Curso Basico de Go

En esta lectura te llevarás algunos tópicos que siempre vale la pena tener a la mano al momento de programar en Go, además te dejo algunos tips extras.

- Hola mundo

```
package main

import "fmt"

func main(){
    fmt.Println("Hola mundo")
}

```

- ¿Hacer una impresión en consola rápida?

```
package main

func main(){
    print("Hola")
}

```

- Importar una librería sin usarla

```
package main

/*
    Hazlo solo y únicamente cuando la librería externa
    que estés usando lo pida explícitamente
*/
import (
    "fmt"
    _ "math"
)

func main(){
    fmt.Println("Hola mundo")
}

```

- Agregar un alias a un import (no suele usarse, pero es bueno saberlo)

```
package main

import (
	"fmt"
	mth "math"
)

func main() {
	fmt.Println(mth.Pi)
}

```

- Diferentes formas de declarar variables

```
v := 12
var vint = 12
var vint
```

- Zero values de primitivos

```
var aint // 0
var bfloat64 // 0
var cstring // ""
var dbool // false

```

- Incremental y decremental

```
x++ // Suma 1 a x
x-- // Resta 1 a x

```

- Imprimir tipo de variables (hay otras formas, pero esta es la más fácil)

```
a := 2
fmt.Printf("%T", a)

```

- Función para tomar los errores (ahorra mucho código)

```
func isError(e error) {
if e !=nil {
        log.Fatal(e)
    }
}

// Ejemplo de uso
func main() {
	_, err := strconv.Atoi("53a")
	isError(err)
}

```

- Arrays vs Slices

```
// Array
var myList [2]int// Slice
var myList2 []int
```

- Slice de interfaces (Úsalo con sabiduría)

```
// Permite guardar diferentes tipos de datos en un mismo slice
myList := []interface{}{"Hola", 12, 4.90}

// Iterar sobre los distintos tipos de datos de ese slice
for _, v :=range myList {
switch v.(type) {
caseint:
        fmt.Println("Es int")
casestring:
        fmt.Println("Es string")
casefloat64:
        fmt.Println("Es float64")
    }
}

```

- Asegurarnos si un key existe en el map

```
m := make(map[string]int)

m["hola"] = 1

// Nota, usalmente se usa "ok" para recibir la segunda variable
value, ok := m["hello"]

/*
Si existe, ok será "true"
Si no existe, ok será "false"

En este caso, ok es "false" porque no existe.
*/

```

- Punteros

```
a := 10 // Variable int
b := &a // "b" es el puntero de "a"
c := *b // "c" adquiere el valor del puntero de "b", es decir toma el mismo valor de "a"

```

- Comandos de Go modules

```
// Inicializar un proyecto
go mod init path_del_proyecto

// Verificar que el código externo no esté corrupto
go mod verify

// Reemplazar fuente del código
go mod edit -replace path_del_repo_online=path_del_repo_en_local

// Quitar el replace
go mod edit -dropreplace path_del_repo_online

// Empaquetar todo el código de terceros que usa nuestro código
go mod vendor

// Eliminar todos los paquetes externos que no estemos usando
go mod tidy

// Aprender más de go modules
go help mod
```

## Inicio

### Nota:

Go no usa los objetos como la suele manejar Java, Aquí no hay clases ni la herencia tradicional. En su lugar utilizan **********************composición y delegación********************** para lograr la reutilización de código.

El manejo de errores no es con excepciones como en Java. Go usa los ******************retornos****************** donde es posible devolver el valor y un error en las funciones para manejar las condiciones excepcionales. 

**************El rendimiento en Go:**************  

- Go es conocido por su rendimiento eficiente y su bajo consumo de recursos.
- La combinación  de su recolector de basura eficiente.
- Su modelo de **************************concurrencia************************** y ****************su compilación estática**************** producen programas que suelen ser mas rápidos que disminuye el uso de la memoria ram en comparación con java.

### Instalación de Go

>Home

```go
sudo mkdir temp
wget http://go.dev/dl/go1.19.2.linux-amd64.tar.gz
sudo tar -xvf go.1.19.2.linux.amd64.tar.gz
sudo mv go /usr/local
export GOPATH=$home/ingsergio/go
export GOBIN=$GOPATH/bin
export GOROOT=/usr/local/go

export PATH=$PATH:$GOBIN:$GOROOT/bin

```

Dentro de la carpeta go, creo: bin, pkg, src

************bin :************ para guardar todos los ejecutables que vallamos utilizando o creando, ya se de librerías de terceros o que creemos nosotros. 

************************************pkg :************************************ guardar cierto código que va a depender tambien de go modules y cada una de las dependencias y las librerías que valla a utilizar. 

**********************Hola mundo:**********************

```go
pakage main

import "fmt"

func main() {
	fmt.Println("Hola Mundo")
}
```

generar un archivó compilador:

>>go buil src/main.go

ó si solo quiero probar 

>>go run src/main.go

### Variables, Constantes y Zero values

********************Constantes********************

```go
const pi float64 = 3.14
const pi2 = 3.1416

```

******************Variables******************

```go
// Declaracion de variables enteras
		base := 12
		var altura int = 14
		var area int
```

**********************Zero values**********************

```go
// Zero values
		var a int // output: 0
		var b float64 // output: 0
		var c string // output: 
		var d bool // output: false
```

### ****Operadores aritméticos****

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/12e7fb7b-4e5b-4c76-934b-f38507c81d36/Untitled.png)

```go
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
```

Reto Área del cuadrado 

```go
const baseCuadrado = 10
areaCuadrado := baseCuadrado * baseCuadrado
fmt.Println(areaCuadrado)
```

### ****Tipos de datos primitivos****

Al codificar en `Go` podemos especificar el tipo de dato, permitiéndonos ganar gran preformas en nuestro código

`int` Cuando no se declara el tamaño tomara la referencia del OS (Sistema operativo) (32 o 64 bits)

********************************Números Enteros********************************

`int` Cuando no se declara el tamaño tomara la referencia del OS (Sistema operativo) (32 o 64 bits)

```go
int8 8bits ⇒ -128 a 127
int16 16bits ⇒ -2^15 a 2^15-1
int32 32bits ⇒ -2^31 a 2^31-1
int64 64bits ⇒ -2^63 a 2^63-1
```

**********Números Enteros Positivos********** 

Optimizar memoria cuando sabemos que el dato siempre va ser positivo

```go
uint ⇒ Depende del OS (32 o 64 bits)
uint8 ⇒ 8bits = 0 a 127
uint16 ⇒ 16bits = 0 a 2^15-1
uint32 ⇒ 32bits = 0 a 2^31-1
uint64 ⇒ 64bits = 0 a 2^63-1
```

****Números Decimales****

```go
float32 ⇒ 32 bits = +/- 1.18e^-38 +/- -3.4e^38
float64 ⇒ 64 bits = +/- 2.23e^-308 +/- -1.8e^308
```

**********************Textos y Booleanos********************** 

• A diferencia de otros lenguajes de programación donde para definir una variable de tipo `string` es permitido utilizar `“”`, `‘’` o ```` en Go solo podemos utilizar las comillas dobles `""`

```go
boolean ⇒ Trueo False
```

************************************Números complejos************************************ 

```go
Complex64 ⇒ Real e Imaginario float32
Complex128 ⇒ Real e Imaginario float64
// EJ:
c := 10 + 8i
c2 := complex(10, 8) // complex 128 por defecto 
c3 := complex64(10 + 8i)

```

### ****Paquete fmt: algo más que imprimir en consola****

```go
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
```

### ****Uso de funciones****

```go
package main

import "fmt"

//Creando funciones
func calculateRectangleArea(b, h int) {
	rectangleArea := b * h
	fmt.Println("El área del rectángulo es:", rectangleArea)
}

func calculateTrapezeArea(b1, b2, h1 float64) {
	trapezeArea := (b1 + b2) * h1 / 2
	fmt.Println("El área del trapecio es:", trapezeArea)
}

func calculateCircleArea(r float64) {
	const pi float64 = 3.14
	circleArea := pi * (r * r)
	fmt.Println("El área del circulo es:", circleArea)
}

//En main epezará a correr en cónsola
func main() {
	calculateRectangleArea(5, 10)
	calculateTrapezeArea(15, 12, 6)
	calculateCircleArea(10)
}
```

### ****El poder de los ciclos en Golang: for, for while y for forever****

En Go solo tengo un caso de ciclo iterativo el cual es el For

```go
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

	// for infinito 
	conterForever := 0
	for {
		fmt.Println(conterForever)
		conterForever++
	}
```

### ****Operadores lógicos y de comparación****

Son operadores que nos permiten hacer una comparación de condiciones y en caso de cumplirse como sino ejecutarán un código determinado. Si se cumple es VERDADERO/TRUE y si no se cumple son FALSO/FALSE.

Empecemos con los operadores de comparación:

## Operadores de comparación

Son aquellos que retornan TRUE o FALSE en caso de cumplirse o no una expresión. Son los siguientes:

- *valor1 == valor2*: Retorna TRUE si valor1 y valor2 son exactamente iguales.
- *valor1 != valor2*: Retorna TRUE si valor1 es diferente de valor2.
- *valor1 < valor2*: Retorna TRUE si valor1 es menor que valor2
- *valor1 > valor2*: Retorna TRUE si valor1 es mayor que valor2
- *valor1 >= valor2*: Retorna TRUE si valor1 es igual o mayor que valor2
- *valor1 <= valor2*: Retorna TRUE si valor1 es menor o igual que valor2.

## Operadores lógicos

Son aquellos que retorna TRUE o FALSE si cumplen o no una condición utilizando [puertas lógicas](https://platzi.com/clases/1050-programacion-basica/15968-que-son-tablas-de-verdad-y-compuertas-logicas/).

### Operador AND:

Este operador indica que todas las condiciones declaradas deben cumplirse para poderse marcar como TRUE. En Go, se utiliza este símbolo `&&`.

Ejemplo1: `1>0 && 2 >0` Esto retornará TRUE porque tanto la primera como la segunda condición son verdaderas.

Ejemplo2: `2<0 && 1 > 0` Esto retornará FALSE porque una de las condiciones no es verdadera.

### Operador OR:

Este operador indica que al menos una de las condiciones debe cumplirse para marcarse como TRUE. En Go, se representa con el símbolo `||`.

Ejemplo: `2<0 || 1 > 0` Esto retornará TRUE porque la segunda condición se cumple, a pesar que la primera no.

### Operador NOT:

Este operador retornará el opuesto al boleano que está dentro de la variable. Ejemplo:

```go
myBool :=true
fmt.Println(!myBool) // Esto retornará false

```

---

Una vez ya estudiado la teoría, en la siguiente clase vamos a ver cómo utilizarlo con más detalles en Go.

### ****El condicional if****

```go
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
```

```go
package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {

	if true {
		fmt.Println("Hello World")
	} else {
		fmt.Println("Bye World")
	}

	//Converting strings to integers
	value, err := strconv.Atoi("23") // Atoi is equivalent to ParseInt(s, 10, 0), converted to type int.

	if err != nil {
		log.Fatal(err) //Fatal is equivalent to Print() followed by a call to os.Exit(1).
	}
	fmt.Println(value)

	//Even or odd
	numb := 10
	if numb%2 != 0 {
		fmt.Println("Number is odd")
	} else {
		fmt.Println("Number is even")
	}

	//login
	user := "Pedro"
	password := "12345"

	if user == "Pedro" && password == "12345" {
		fmt.Println("Logged in")
	} else if user == "Pedro" {
		fmt.Println("Password incorrect")
	} else if password == "12345" {
		fmt.Println("User incorrect")
	} else {
		fmt.Println("Credencials incorrect")
	}

}
```

### ****Múltiple condiciones anidadas con Switch****

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/db05fde4-5864-4165-882f-64c0b1daffe1/Untitled.png)

```go
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
```

### ****El uso de los keywords defer, break y continue****

```go
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
```

**Nota importante para la instrucción `defer`**

Algo importante que se debe tomar en cuenta con el uso del defer, es que si dentro de una misma función usas más de un defer, estos son tratados como una pila o en inglés “stack” uno encima del otro lo que genera que los defer se ejecuten con el principio LIFO (Last In First Out / Último en entrar, primero en salir) y hay que ser cuidadosos en ese aspecto.

Aquí dejo un ejemplo donde se puede apreciar mejor el asunto:

```
package main

import "fmt"

funcmain() {
   fmt.Println("Inicio")
for i := 0; i < 5; i++ {
defer fmt.Println(i)
   }
   fmt.Println("Fin")
}

```

Si revisas la salida verás lo siguiente:

```
Inicio
Fin
4
3
2
1
0

```

y no

```
Inicio
Fin
0
1
2
3
4

```

Como tal vez se podría pensar.

## ****Arrays y Slices****

```go
// Array
// len para mostrar la cantidad de datos en el array
// cap para mostrar la capacidad que tiene el array
var array [4]int
array[0] = 1
array[1] = 2
fmt.Println(array, len(array), cap(array))
```

Se usa **len** para saber la cantidad de objetos en el array y ********cap******** para saber la cantidad de objetos que puede almacenar el array. 

```go
// Slice
slice := []int{0, 1, 2, 3, 4, 5, 6}
fmt.Println(slice, len(slice), cap(slice))

// Métodos en el slice
fmt.Println(slice[0])
fmt.Println(slice[:3]) // el numero despuesd los : es exclusivo, ose que no se muestra
fmt.Println(slice[2:4])// la primera posicion es inclusiva y la ultima es exclusiva 
fmt.Println(slice[4:])// el primer numero es inclusivo

// Append en los slice
slice = append(slice, 7)
fmt.Println(slice) // output: [0, 1, 2, 3, 4, 5, 6, 7]

// Append nueva lista
newSlice := []int{8, 9, 10}
slice = append(slice, newSlice...)
fmt.Println(slice) // output: [0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10]
```

## ****Recorrido de Slices con Range****

```go
// Recorrido de Slices con Range.
slice2 := []string{"Hola", "Que", "Hace"}
for i := range slice2 { 
	// aqui solo pretendo tabajar con el indice, si tambien quisiera trabajar con el valor seria asi: 
	// for i, valor := range slice2
	fmt.Println(slice2[i]) 

}
```

************************************************Ejercicio de Palíndromo************************************************ 

```go
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

//En el main
// Probando la funcion del palindromo
isPalindromo("oso")

```

## Estructura de ****Llave valor con Maps(en go en otros lenguajes se les conoce como diccionarios)****

 ********

```go
m := make(map[string]int)

m["Jose"] = 14
m["Sergio"] = 20

fmt.Println(m)

// Recorrer el map
for i, v := range m {
	fmt.Println(i, v)
}
// usualmente no vamos a obtener el llave, valor en el mismo orden que fue ingresado ya que este proceso ocurre de forma concurrente.
// si me interes es que imprima en el mismo orden, se recomienda usar los Slice

// Encontrar un valor
	value8, ok := m["Jose"]
	fmt.Println(value8, ok)
// si Jose no existe, Output: 0 false
// si Jose existe, nos imprmira el valor de Jose y true: Output: 14 true

```

## ****Structs: La forma de hacer clases en Go****

```go
type car struct {
	brand string
	year  int
}

func main() {
	// instanciar el struc
	myCar := car{brand: "Ford", year: 2020}
	fmt.Println(myCar) // output: {Ford 2020}

	// otra manera de intanciar un struct
	var otherCar car
	otherCar.brand = "Ferrari"
	fmt.Println(otherCar)

	var myCar2 pk.CarPublic
	myCar2.Brand = "Ferrari"
	myCar2.Year = 2020
	fmt.Println(myCar2)
	
}
```

## Resolver problemas al tratar de manejar el modulo.

En mi caso tuve varios problemas al tratar de manejar con módulos porque los archivos los creo en una carpeta diferente. Para esto uso el siguiente comando en la carpeta que quiero que sea la raíz de mi proyecto:

```
go mod init proyecto

```

Al ejecutar este comando se crea un archivo go.mod. Luego en esta misma carpeta creo el archivo main.go y creo el paquete. De forma que la estructura de carpetas sería la siguiente:

```
-proyecto
	|- go.mod
	|- main.go
	|- my_package/
		my_package.go

```

Para poder importar el paquete my_package, no solo especifico la ruta en el import (que seria proyecto/my_package), sino que también ejecuto el siguiente comando:

```
go mod tidy

```

Con este comando hace, según ChatGPT, lo siguiente:

1. Analiza el código fuente de tu proyecto para detectar qué módulos son requeridos por tu proyecto y actualiza el archivo “go.mod” con las dependencias necesarias.
2. Elimina las dependencias innecesarias y obsoletas del archivo “go.mod”.
3. Actualiza el archivo “go.sum” para incluir los hashes de las versiones de las dependencias utilizadas por tu proyecto.
4. Descarga automáticamente las dependencias faltantes o actualiza las dependencias existentes a las últimas versiones disponibles, si es necesario.

Espero les sirva

## ****Modificadores de acceso en funciones y Structs****

```go
package mypackage

import "fmt"

type CarPublic struct {
	Brand string
	Year  int
}

type CarPrivate struct {
	brand string
	year  string
}

// funciones de acceso para los atributos de tipo privado

func (c *CarPrivate) brandSet(brand string) {
	c.brand = brand
}

func PrintMesseage(text string) {
	fmt.Println(text)
}
```

```go

import (
	pk "curso_golang_platzi/src/mypackage"
	"fmt"
}
	
func main() {

	var myCar2 pk.CarPublic
	myCar2.Brand = "Ferrari"
	myCar2.Year = 2020
	fmt.Println(myCar2)
	
	// provando la funcion publica
	pk.PrintMesseage("Prueba")

}

```

## ****Structs y Punteros****

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/d847c8f1-2765-400b-bb46-73385eb48cb9/Untitled.png)

```go

package main

import (
	"fmt"
}

type pc struct {
	ram   int
	brand string
	disk  int
}

func (myPc pc) ping() {
	fmt.Println(myPc.brand, "Pong")
}

//funcion para duplicar la ram 
// y como vamos a acceder a los datos de Pc para hacer uso de sus datos usamos el *Pc
func (myPc *pc) duplicateRAM{
	myPc.ram = myPc.ram * 2
}

func main() {
	// difino los endpoint
	myPc := pc{ram: 16, disk: 200, brand: "msi" }
	fmt.Println(myPc) //output: {16, 200, msi}

	myPc.ping() //output: msi Pong
	
	// para cambiar el dato de la ram podemos hacer esto
	//myPc.ram = 20 //pero esto no suele ser lo adecuado 
	// asi que solemos implementar funciones para realizar estos cambios 
	// cremos un funcion para duplicar la ram 
	
	fmt.Println(myPC)
	myPc.duplicateRAM()
	
	fmt.Println(myPC)
	myPc.duplicateRAM()

	fmt.Println(myPC)
	myPc.duplicateRAM()

	fmt.Println(myPC)

}

//OutPut:
//{16 200 msi}
//msi Pong
//{16 200 msi}
//{32 200 msi}
//{32 200 msi}

```

## ****Stringers: personalizar el output de Structs****

```go
package main

import (
	"fmt"
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

func main() {
	// difino los endpoint
	myPc := pc{ram: 16, disk: 200, brand: "msi" }
	fmt.Println(myPc) //output: {16, 200, msi}

}

//OutPut:
// Tengo 16 GB RAM, 100GB Disco y es una msi
```

## ****Interfaces y listas de interfaces****

```go
package main

import (
	"fmt"
}

// creo la interfaz y el metodo que va a compartir los dos Structs
type figuras2D interface {
	area() float64
}

// Creamos los Structs
type cuadrado struct {
	base float64
}

type rectangulo struct {
	base   float64
	altura float64
}

// Creo la funciones del area respecto al tipo del Struct
// esta funciones no le creo lo que va a resivir si no el tipo de objeto que va a implementar 
// por que la funcion que me va a implemetar esta funciones es la del calcular mediante el uso de la interfaz que implementa el metodo area()
func (c cuadrado) area() float64 {
	return c.base * c.base
}

func (r rectangulo) area() float64 {
	return r.base * r.altura
}

// esta funcion me toma como entrada la interfaz y hacemos que esta funcion ejecute la funcion de la interfaz del are (f.area())
func calcular(f figuras2D) {
	fmt.Println("Area: ", f.area())
}

func main() {

	myCudrado := cuadrado{base: 4}
	myRectangulo := rectangulo{base: 4, altura: 2}
	
	calcular(myCudrado)
	calcular(myRectangulo)

	// Lista de interfaces. aqui puedo manejar distintos tipos de datos
	myInterface := []interface{}{"hola", 12, 4.90}
	fmt.Println(myInterface...)
}

//Output: 
//Area: 16
//Area: 8
//hola 12 4.9

```

## ****¿Qué es la concurrencia?****

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/29c216c1-61ce-4558-b2c2-6cb6934610e6/Untitled.png)

Concurrencia te permite estar pendiente de varios procesos, comienzas uno, empiezas otro, ves si el anterior ya terminó, luego crear otro así

El paralelismo, es usar varios hilos del procesador para hacer varios procesos a la vez, pero siempre estas esperando que la tarea termine.

## ****Primer contacto con las Goroutines****

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/ec93bf6a-6bb6-4fba-aced-583425d2d9b3/Untitled.png)

Al agregar un tiempo de espera para que para que responda la ****Goroutines, esto de igual forma es contraproducente al momento de hacer eficiente nuestro código.**** 

### Para solucionar esto hay dos formas:

sync.WaitGroup  

Por lo al agregar una ****Goroutines al**** WaitGroup para que espere su ejecución antes de que la ****Goroutine base(main) muera !****

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/42773ec6-5fd4-4d27-bfe3-c546865a1c7a/Untitled.png)

Para poder mantener esta función tenemos que agregarle una modificación  

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/376c5f02-b76d-4fe8-b0bd-d9da03f1cfb4/Untitled.png)

wg.Wait   esto le indica que tiene que esperar ha que las dos **Goroutines terminen.** 

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/0ae9d148-e958-4b50-8ca3-304d6742a17b/Untitled.png)

en este otro caso implementamos una función anónima dentro del main para que tambien imprima un text.

Pero usualmente de esta forma no se suele utilizar, por lo que son mas complejas de mantener. Así que para esto utilizamos los ****Channels.****

## ****Channels: La forma de organizar las goroutines****

Permite compartir datos entre los ****goroutines****

Los channels con un conducto en el cual solo puedo manejar un tipo de dato. 

Primero debemos definir el channel: 

**c := make(chan string, 1)**

le indicamos cantidad de groutines que va a ejecutar a la vez - **esto es opcional pero es una buena practica.** 

luego en la función que pretendemos que se ejecute de manera concurrente vamos a agregarle el canal. 

```go
func say(text string, c chan string) { 
 // ahora vamos a tomar el texto y se lo vamos a pasar al canal 
	c <- text
}

func main() {
	
	c := make(chan string, 1)
	
	fmt.Println("Hello")
	
	go say("Bye", c)
	
	fmt.Println(<-c)
```

tambien en el chan podremos especificar si se va a comporta de forma de entrada o salida o de ambas, asi que se lo podemos indicar de las siguiente formas: 

Si solo es de entrada de datos, como en nuestro caso: 

```go
func say(text string, c chan<- string) {
	c <- text
}
```

Si solo es de salida: 

```go
func say(text string, c <-chan<- string) {
 text = <- c
}
```

Resumen: 

```go
// channels
// tambien puedo especificar que el parametro del canal va hacer
// o de entrado o de salida, en este caso especificamos
// que va hacer de entrado con chan<- a su lado derecho si fuera de salida seria <-chan
func say(text string, c chan<- string) {
	c <- text
	// si tambine tiene parametros de salid seria
	// text = <- c
}
```

## ****Range, Close y Select en channels****

para capturar información de la estructura del cahnnels: 

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/5aeb6500-d1c7-4357-8f5e-d857ecb72cd5/Untitled.png)

### Es recomendable siempre cerrar el Channels.

por eso: 

### Close:

```go
close(c)
```

y una ves que esta cerrado no puedo agregarle datos al channels: 

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/05dd04bf-bf4f-4f90-a846-affac7358b2d/Untitled.png)

### Range

Es ideal cuando quiero iterar en iterar entre cada uno de los elementos de un canal. 

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/7e107064-0d5f-4ee6-8640-f3b51ebaea36/Untitled.png)

### Select

Cuando estamos ejecutando múltiples canales y no sabemos cual de todos va a responder primero, por eso utlizamos **Select**

Vamos hacer el ejercicio de agregar un texto a un canal especifico. Por lo que creamos la función message que va a resibir el texto   y un dato de tipo canal. 

por lo que el canal va a recibir un text:

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/1e911556-2c20-4d7b-a5b0-bea7a1f98797/Untitled.png)

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/e4d01789-6c1e-4607-9604-a910e3c36fd2/Untitled.png)

en el main creamos los dos canales email1 y email2  pero no le definimos su tamaño por que va hacer dinámico. 

Pero al momento no sabríamos cual de los dos canales van a responder a tiempo y para eso **tenemos que tener presente la cantidad de channels** que vamos a manejar. 

Este dato lo necesitamos para poder recorrer los chan. como en este caso se estan manejando 2 creamos un for con 2 iteraciones  y con el  **************select************** vamos a utlizarlo para esperar cual de los canales responde primero, lo hacemos de la siguiente manera: 

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/969490c6-9351-4b6a-8c3d-e20ad94460f3/Untitled.png)

 

## ****Go get: El manejador de paquetes****

**En la terminal:** 

```go
go install -v golang.org/x/website/tour@latest
```

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/ecfcf98d-2861-4efc-8c53-8763fa3a1744/Untitled.png)

## ****Go modules: Ir más allá del GoPath con Gin****

### Vamos a la documentacion de Gin para estudiar sus primeros pasos.

```go

func main() {

r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
		c.String(http.StatusOK, "Hola mundo")
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

}
```

Validar en nuestro go.mod del vedana

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/0dbd3c95-adbf-463e-b6c3-ba72d01f5476/Untitled.png)

En Go puedo aplicar modificación de módulos lo que significa es que si quiero puedo hacer cambios a librerías de terceros. y que se actualicen en nuestros proyectos. 

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/d93db5ae-3559-447d-b6d8-7fa41206e2ea/Untitled.png)

 

https://gobyexample.com/

https://play-with-go.dev/

[**https://gophercises.com/**](https://gophercises.com/)

https://blog.gopheracademy.com/gophers-slack-community/
