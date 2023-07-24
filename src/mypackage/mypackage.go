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

func printMesseage1(text string) {
	fmt.Println(text)
}

func PrintMesseage(text string) {
	fmt.Println(text)
}
