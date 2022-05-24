package main

import (
	"fmt"
	"main/paquetes"
)

func main() {

	fmt.Println("------  lista simple ---------------")
	lista := paquetes.Lista{}
	lista.Agregar(1)
	lista.Agregar(2)
	lista.Agregar(3)
	lista.Mostrar()
	fmt.Println("---------------------")
	lista.Actualizar(2, 99)
	lista.Mostrar()

	fmt.Println("------  lista Doble ---------------")
	lista_d := paquetes.Lista_doble{}
	lista_d.Agregar(11)
	lista_d.Agregar(22)
	lista_d.Agregar(33)
	lista_d.Agregar(44)
	lista_d.Mostrar_Ant()
	fmt.Println("---------------------")
	lista_d.Mostrar_Sig()
	lista_d.Actualizar(33, 789)
	fmt.Println("---------------------")
	lista_d.Mostrar_Sig()
	fmt.Println("---------------------")

	lista_d.Mostrar_Ant()
}
