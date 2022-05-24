package main

import (
	"fmt"
	"main/paquetes"
)

func main() {

	lista := paquetes.Lista{}
	lista.Agregar(1)
	lista.Agregar(2)
	lista.Agregar(3)
	lista.Mostrar()
	fmt.Println("---------------------")
	lista.Actualizar(2, 99)
	lista.Mostrar()

}
