package Variables

import (
	"fmt"
	"time"
)

//NOTA: Si las varibles o el nombre de las funciones se encuentran en mayusculas, estas podrán ser vistas por los otros archivos .GO ya que son publicas, si de lo contrario, estas son minusculas, el lenguaje .GO las tomará como privadas y solo podran ser vistas dentro del propio documento el cual se creó dicha variable

var Nombre string
var Estado bool
var Sueldo float64
var Fecha time.Time

func RestoVariables() {

	Nombre = "Diego Fernando"
	Estado = false
	Sueldo = 536576.897987
	Fecha = time.Now()

	fmt.Println("Nombre: ", Nombre)
	fmt.Println("Su estado: ", Estado)
	fmt.Println("Su sueldo: ", Sueldo)
	fmt.Println("La fecha de hoy es: ", Fecha)

}
