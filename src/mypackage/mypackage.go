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

// como este es un metodo privado no lo puedo implementar en el main del paquete main.go como el metodo de abajo que es publico
func printMesseage1(text string) {
	fmt.Println(text)
}

func PrintMesseage(text string) {
	fmt.Println(text)
}
