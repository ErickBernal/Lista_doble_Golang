package paquetes

import "fmt"

type Nodo_doble struct {
	Dato int
	Ant  *Nodo_doble
	Sig  *Nodo_doble
}

type Lista_doble struct {
	//nil <-> 1 <-> 2 <- n -> nil				nil <-> 1 <->  2 <-> nil
	Cola   *Nodo_doble // nodo final
	Cabeza *Nodo_doble // nodo inicial
}

func (l *Lista_doble) Agregar(nuevo_dato int) {
	nuevo_Nodo_doble := &Nodo_doble{}
	if l.Cola == nil {
		nuevo_Nodo_doble = &Nodo_doble{Dato: nuevo_dato, Ant: nil, Sig: nil}
		l.Cabeza = nuevo_Nodo_doble
	} else {
		nuevo_Nodo_doble = &Nodo_doble{Dato: nuevo_dato, Ant: l.Cola, Sig: nil}
		l.Cola.Sig = nuevo_Nodo_doble
	}
	l.Cola = nuevo_Nodo_doble
}

func (l *Lista_doble) Mostrar_Ant() {
	aux := l.Cola
	for aux != nil {
		fmt.Println("Doble_Cola_ant <", aux.Dato, ">")
		aux = aux.Ant
	}
}

func (l *Lista_doble) Mostrar_Sig() {
	aux := l.Cabeza
	for aux != nil {
		fmt.Println("Doble_Cola_sig <", aux.Dato, ">")
		aux = aux.Sig
	}
}

func (l *Lista_doble) Actualizar(data int, nueva_data int) {
	aux := l.Cola
	for aux != nil {
		if data == aux.Dato {
			aux.Dato = nueva_data
			break
		}
		aux = aux.Ant
	}
}
