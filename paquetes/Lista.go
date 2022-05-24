package paquetes

import "fmt"

type Nodo struct {
	Dato int
	Ant  *Nodo
}

type Lista struct {
	//nil<- 1 <- 2 <- n
	Cola *Nodo
}

func (l *Lista) Agregar(nuevo_dato int) {
	nuevo_nodo := &Nodo{}
	if l.Cola == nil {
		nuevo_nodo = &Nodo{Dato: nuevo_dato, Ant: nil}
	} else {
		nuevo_nodo = &Nodo{Dato: nuevo_dato, Ant: l.Cola}
	}
	l.Cola = nuevo_nodo
}

func (l *Lista) Mostrar() {
	aux := l.Cola
	for aux != nil {
		fmt.Println("Cola<", aux.Dato, ">")
		aux = aux.Ant
	}
}

func (l *Lista) Actualizar(data int, nueva_data int) {
	aux := l.Cola
	for aux != nil {
		if data == aux.Dato {
			aux.Dato = nueva_data
			break
		}
		aux = aux.Ant
	}
}
