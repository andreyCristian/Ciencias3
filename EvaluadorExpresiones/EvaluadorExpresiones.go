package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Arbol struct {
	Izquierda  *Arbol
	Valor string
	Derecha *Arbol
}

type Nodo struct {
	expresion string
}

type Fifo struct {
	nodos []*Nodo
	cont int
}

func nuevaCola() *Fifo{
	return &Fifo{
		nodos: make([]*Nodo, 1),
	}
}

func (q *Fifo) insertar(n *Nodo){
    if  q.cont > 0 {
		nodes := make([]*Nodo, len(q.nodos) + 1)
		copy(nodes, q.nodos[0:])
		q.nodos = nodes
	}
    q.nodos[len(q.nodos) - 1] = n
	q.cont++
}

func (q *Fifo) insertarNodo(expresion string){
    q.insertar(&Nodo{expresion})
}

func (q *Fifo) eliminar() string{
	pop := "";

	if q.cont != 0 {
        nodes := make([]*Nodo, len(q.nodos) - 1)
		pop = q.nodos[0].expresion
		copy(nodes, q.nodos[1:])
        q.nodos = nodes
        q.cont--
    }

	return pop
}

func ConstruirArbol(sentencia []string) *Arbol{
	var t *Arbol

	for i := 0; i < len(sentencia); i++ {
		t = Insertar(t, sentencia[i])
	}

	return t
}

func Separar(aux string) []string{
	cadena := strings.Split(aux, " ")
	return cadena
}

func Insertar(t *Arbol, a string) *Arbol{
	if t == nil {
		return &Arbol{nil, a, nil}
	}

	if !EvaluarOp(a) {
		InsertarNum(t, a)
	}

	InsertarOp(t, a)
	return t
}

func InsertarOp(t *Arbol, a string) *Arbol{
	if a == "+" || a == "-" {
		if EvaluarProfundidad(t.Izquierda){
			t.Derecha = Insertar(t.Derecha, a)
			return t
		}else {
			t.Izquierda = Insertar(t.Izquierda, a)
			return t
		}
	} else if a == "x" || a == "*" || a == "/"{
		if EvaluarProfundidad(t.Izquierda){
			t.Derecha = Insertar(t.Derecha, a)
			return t
		}else {
			t.Izquierda = Insertar(t.Izquierda, a)
			return t
		}
	}
	return t
}

func InsertarNum(t *Arbol, a string) *Arbol{
	if EvaluarProfundidad(t.Izquierda){
		t.Derecha = Insertar(t.Derecha, a)
	} else if EvaluarOp(t.Valor){
		if t.Izquierda == nil || EvaluarOp(t.Izquierda.Valor){
			t.Izquierda = Insertar(t.Izquierda, a)
			return t
		} else {
			t.Derecha = Insertar(t.Derecha, a)
			return t
		}
	}
	return t
}

func EvaluarOp(aux string) bool{
	_, b := strconv.ParseInt(aux, 10, 64)
	if b == nil{
		return false
	} else{
		return true
	}
}

func EvaluarProfundidad(t *Arbol) bool{
	if t == nil{
		return false
	}

	if EvaluarOp(t.Valor){
		if t.Izquierda == nil{
			return false
		} else if t.Derecha == nil {
			return false
		}
	}
	return true;
}

func CadenaFinal(cadena []string, mapa map[string]int64) []string {
	for a, b := range cadena {
		for c, d := range mapa {
			if b == c {
				cadena[a] = strconv.FormatInt(d, 10)
			}
		}
	}

	return cadena
}

func Resultado(t *Arbol) int64{
	if t == nil {
		return 0
	}

	Resultado(t.Izquierda)
	if EvaluarOp(t.Valor){
		switch t.Valor {
		case "+":
			return Resultado(t.Izquierda) + Resultado(t.Derecha)
		case "-":
			return Resultado(t.Izquierda) - Resultado(t.Derecha)
		case "*":
			return Resultado(t.Izquierda) * Resultado(t.Derecha)
		case "/":
			return Resultado(t.Izquierda) / Resultado(t.Derecha)
		case "x":
			return Resultado(t.Izquierda) + Resultado(t.Derecha)
		}
	}
	a, _ := strconv.ParseInt(t.Valor, 10, 64)
	return a
}

func RecorrerInorden(t *Arbol) {
	if t == nil {
		return
	}

	RecorrerInorden(t.Izquierda)
	fmt.Print(t.Valor)
  	fmt.Print(" ")
	RecorrerInorden(t.Derecha)
}

func main() {
	sentencias := nuevaCola()
	m := make(map[string]int64)
	var t1 *Arbol
	var cadena []string

	sentencias.insertarNodo("/ 8 - 5 1 := X")
	sentencias.insertarNodo("+ 2 2 := Y")
	sentencias.insertarNodo("* X Y := Z")

	cadena = Separar(sentencias.eliminar())
	fmt.Println("La entrada fue la siguiente: " , cadena)
	t1 = ConstruirArbol(cadena)
	fmt.Print("Se ha realizado la siguiente operacion: ")
	RecorrerInorden(t1)
	m[cadena[len(cadena) - 1]] = Resultado(t1)
	fmt.Println("\nEl resultado fue:", m["X"])

	cadena = Separar(sentencias.eliminar())
	fmt.Println("La entrada fue la siguiente: " , cadena)
	t1 = ConstruirArbol(cadena)
	fmt.Print("Se ha realizado la siguiente operacion: ")
	RecorrerInorden(t1)
	m[cadena[len(cadena) - 1]] = Resultado(t1)
	fmt.Println("\nEl resultado fue:", m["X"])

	cadena = CadenaFinal(Separar(sentencias.eliminar()), m)
	fmt.Println("La entrada fue la siguiente: " , cadena)
	t1 = ConstruirArbol(cadena)
	fmt.Print("Se ha realizado la siguiente operacion: ")
	RecorrerInorden(t1)
	m[cadena[len(cadena) - 1]] = Resultado(t1)
	fmt.Println("\nEl resultado fue:", m["Z"])
}
