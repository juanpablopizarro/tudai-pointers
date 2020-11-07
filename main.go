package main

import "fmt"

type Person struct {
	Name string
}

// & devuelve la direccion de memoria de una variable
// * a nivel definicion de funcion indica que se requiere una posicion de memoria
// * a nivel operador hace referencia al valor que contiene el puntero

func (x *Person) changeName(n string) {
	x.Name = n
	fmt.Printf("%p %v\n", &x, x)
}

func main() {
	p := Person{"Juan"}
	fmt.Printf("%p %v\n", &p, p)
	p.changeName("Pedro")
	fmt.Printf("%p %v\n", &p, p)
}
