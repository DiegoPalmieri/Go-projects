package main

import (
	"fmt"

	"github.com/DiegoPalmieri/Go-projects/Variables"
) //De esta forma se importa un paquete en GO

func main() {

	Variables.MostrarEnteros()

	Variables.RestoVariables()

	Estado, texto := Variables.ConvertiraTexto(8786)

	fmt.Println(Estado)
	fmt.Println(texto)

}
