package Variables

import (
	"fmt"
	"time"
)

func MostrarInformacion() {

	Nombre := "Diego Fernando"
	Apellido := "Palmieri Barreto"
	Edad := 19
	Peso := 50.09
	Fecha := time.Now()

	fmt.Println("\nNombre:", Nombre, "\nApellido:", Apellido, "\nEdad:", Edad, "\nPeso:", Peso, "\nFecha:", Fecha)

}
