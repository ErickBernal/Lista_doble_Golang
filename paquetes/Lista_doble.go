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

func (l *Lista_doble) Agregar_ordenado(nuevo_dato int) {
	nuevo_Nodo_doble := &Nodo_doble{Dato: nuevo_dato, Ant: nil, Sig: nil}

	if l.Cola == nil && l.Cabeza == nil { //lista vacia
		l.Cabeza = nuevo_Nodo_doble
		l.Cola = nuevo_Nodo_doble
	} else {
		//situacion 1: solo existe cabeza
		if l.Cabeza.Ant == nil && l.Cabeza.Sig == nil {
			if nuevo_dato < l.Cabeza.Dato { // si es menor, a la izquierda
				nuevo_Nodo_doble.Sig = l.Cabeza
				l.Cabeza.Ant = nuevo_Nodo_doble
				l.Cola = l.Cabeza
				l.Cabeza = nuevo_Nodo_doble
			} else { //si es mayor o igual, a la derecha
				l.Cabeza.Sig = nuevo_Nodo_doble
				nuevo_Nodo_doble.Ant = l.Cabeza
				l.Cola = nuevo_Nodo_doble
			}
		} else if nuevo_dato < l.Cabeza.Dato {
			//situacion 2: si es menor a cabeza
			nuevo_Nodo_doble.Sig = l.Cabeza
			l.Cabeza.Ant = nuevo_Nodo_doble
			l.Cabeza = nuevo_Nodo_doble
		} else if nuevo_dato > l.Cola.Dato {
			//si es mayor a la cola
			l.Cola.Sig = nuevo_Nodo_doble
			nuevo_Nodo_doble.Ant = l.Cola
			l.Cola = nuevo_Nodo_doble
		} else {
			//si es mayor a cabeza, pero menor a cola => recorremos
			aux := l.Cabeza
			for aux != nil {
				if nuevo_dato < aux.Dato {
					aux.Ant.Sig = nuevo_Nodo_doble
					nuevo_Nodo_doble.Ant = aux.Ant

					nuevo_Nodo_doble.Sig = aux
					aux.Ant = nuevo_Nodo_doble
					break
				} else if nuevo_dato >= aux.Dato && nuevo_dato < aux.Sig.Dato { //si es mayor o igual
					nuevo_Nodo_doble.Sig = aux.Sig
					aux.Sig.Ant = nuevo_Nodo_doble

					aux.Sig = nuevo_Nodo_doble
					nuevo_Nodo_doble.Ant = aux
					break
				}
				aux = aux.Sig
			}
		}
	}
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

func (l *Lista_doble) Eliminar(data int) {
	delete := false
	tmp := l.Cola
	for tmp != nil {
		if tmp.Dato == data {
			delete = true
			if tmp.Ant == nil {
				//situacion 1; Eliminar Cabeza    { nil <-nodo-> Sig }
				l.Cabeza = tmp.Sig
				l.Cabeza.Ant = nil
			} else if tmp.Sig == nil {
				// situcacion 2: si se quiereien eliminar Cola   { prev <-nodo-> nil }
				l.Cola = tmp.Ant
				l.Cola.Sig = nil
			} else {
				//situacion 3: eliminar cualquiera del centro   { prev <-nodo-> Sig }
				tmp.Ant.Sig = tmp.Sig
				tmp.Sig.Ant = tmp.Ant
			}
			break
		}
		tmp = tmp.Ant
	}
	if !delete {
		fmt.Println("<No existe el dato: \"", data, "\" a eliminar>")
	}
}
