package main

import (
	"fmt"
	"strconv"
	"strings"
    "regexp"
	"os"
	"bufio"
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

	if q.cont > 0 {
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

func EvaluarLenguaje(cadena string) bool{
    numExp, _ := regexp.Compile(`-?\d$`)
    opExp, _ := regexp.Compile(`[-]|[+]|[*]|[/]|[:=]$`)
    varExp, _ := regexp.Compile(`[A-Z]$`)
	err := false

	for _, cad := range Separar(cadena) {
		if numExp.MatchString(cad){
			fmt.Println("Numero = ", cad)
			err = true
		} else if opExp.MatchString(cad) {
			err = true
			fmt.Println("Operador = ", cad)
		} else if varExp.MatchString(cad) {
			err = true
			fmt.Println("Variable = ", cad)
		} else {
			err = false
			fmt.Println("ERROR NO ES UNA ENTRADA VALIDA: ", cad)
			break
		}
	}
    return err
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

func RecorrerPreorden(t *Arbol) {
	if t == nil {
		return
	}

	fmt.Print(t.Valor)
  	fmt.Print(" ")
	RecorrerPreorden(t.Izquierda)
	RecorrerPreorden(t.Derecha)
}

func ImprimiMenu() {
	fmt.Println("\n---------------------------------------------------------",
		"\nIngrese la opcion (1,...) que desea realizar:",
		"\n1. Ingresar expresion en la cola.",
		"\n2. Construir árbol de la expresion actual en cola.",
		"\n3. Ingresar expresion final del árbol. (ex. + A B := C)",
		"\n4. Mostrar construccion en innorden del arbol actual en cola.",
		"\n5. Mostrar construccion en preorden del arbol actual en cola.",
		"\n6. Imprimir resultado del árbol.",
		"\n7. Salir del programa.")
}

func ImprimirTabla(){
	fmt.Println("\nTABLA DE LENGUAJE",
		"\nNumero \t\t [0-9]",
		"\nVariables \t [A-Z]",
		"\nOperadores \t [+ - * / :=]",
		"\n---------------------------------------------------------")
}

func main() {
	ciclo := 0
	sc:= bufio.NewScanner(os.Stdin)
	var t1 *Arbol
	sentencias := nuevaCola()
	m := make(map[string]int64)
	var cadena []string

	for ciclo < 1000 {
		ImprimiMenu()
		sc.Scan()

		switch sc.Text() {
			case "1":
				ImprimirTabla()
				fmt.Println("\nIngrese la expresion a evaluar en preorden, separado por espacios (ex. + 2 2 := A), ")
				sc.Scan()
				if EvaluarLenguaje(sc.Text()) {
					if sentencias.cont == 0{
						sentencias = nuevaCola()
					}
					sentencias.insertarNodo(sc.Text())
				}
			case "2":
				if sentencias.cont == 0 {
					fmt.Println("No se encuentran sentencias para armar arboles")
					sentencias = nuevaCola()
				} else {
					cadena = Separar(sentencias.eliminar())
					fmt.Println("Se construira el siguiente árbol: ", cadena)
					t1 = ConstruirArbol(cadena)
					fmt.Println("ARBOL CREADO.")
					m[cadena[len(cadena) - 1]] = Resultado(t1)
				}
			case "3":
				ImprimirTabla()
				fmt.Println("\nAsegurese de haber ingresado las variables correspondientes antes (ex. + A B := C)")
				sc.Scan()
				cadena = CadenaFinal(Separar(sc.Text()), m)
				t1 = ConstruirArbol(cadena)
				fmt.Println("ARBOL CREADO.")
			case "4":
				fmt.Println("\nLa siguiente es la operacion que se va a realizar (innorden):")
				RecorrerInorden(t1)
			case "5":
				fmt.Println("\nLa siguiente es la construccion en preorden del árbol:")
				RecorrerPreorden(t1)
			case "6":
				fmt.Print("\nSe ha realizado la siguiente operacion: ")
				RecorrerInorden(t1)
				m[cadena[len(cadena) - 1]] = Resultado(t1)
				fmt.Println("El resultado fue:", m[cadena[len(cadena) - 1]])
			case "7":
				fmt.Println("Gracias. \nEl aplicativo ha terminado.")
				ciclo += 1000
			default:
				fmt.Println("La opcion ingresada no fue correcta.")
		}
	}
}
