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
	Izquierda *Arbol
	Valor Expresion
	Derecha *Arbol
}

type Nodo struct {
	expresion string
}

type Expresion struct {
	nombreArbol string
	tipoDato string
	Dato string
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
	pop := ""

	if q.cont > 0 {
		nodes := make([]*Nodo, len(q.nodos) - 1)
		pop = q.nodos[0].expresion
		copy(nodes, q.nodos[1:])
		q.nodos = nodes
		q.cont--
	}

	return pop
}

func ConstruirArbol(sentencia []Expresion) *Arbol{
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

func Insertar(t *Arbol, a Expresion) *Arbol{
	if t == nil {
		return &Arbol{nil, a, nil}
	}

	if !EvaluarOp(a.Dato) {
		InsertarNum(t, a)
	}

	InsertarOp(t, a)
	return t
}

func InsertarOp(t *Arbol, a Expresion) *Arbol{
	if a.Dato == "+" || a.Dato == "-" {
		if EvaluarProfundidad(t.Izquierda){
			t.Derecha = Insertar(t.Derecha, a)
			return t
		}else {
			t.Izquierda = Insertar(t.Izquierda, a)
			return t
		}
	} else if a.Dato == "*" || a.Dato == "/"{
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

func InsertarNum(t *Arbol, a Expresion) *Arbol{
	if EvaluarProfundidad(t.Izquierda){
		t.Derecha = Insertar(t.Derecha, a)
	} else if EvaluarOp(t.Valor.Dato){
		if t.Izquierda == nil || EvaluarOp(t.Izquierda.Valor.Dato){
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

func EvaluarLenguaje(cadena string) (bool, []Expresion){
	varExp, _ := regexp.Compile(`[A-Z]$`)
	asigExp, _ := regexp.Compile(`[:=]$`)
	numExp, _ := regexp.Compile(`-?\d$`)
	opExp, _ := regexp.Compile(`[-]|[+]|[*]|[/]$`)

	err := false
	expresiones := make([]Expresion, 0)

	for _, cad := range Separar(cadena) {
		if numExp.MatchString(cad) {
			expresiones = append(expresiones, Expresion{tipoDato: "Numero\t->\t", Dato: cad})
			err = true
		} else if opExp.MatchString(cad) {
			expresiones = append(expresiones, Expresion{tipoDato: "Operador\t->\t", Dato: cad})
			err = true
		} else if varExp.MatchString(cad) {
			expresiones = append(expresiones, Expresion{tipoDato: "Variable\t->\t", Dato: cad})
			err = true
		} else if asigExp.MatchString(cad) {
			expresiones = append(expresiones, Expresion{tipoDato: "Asignacion\t->\t", Dato: cad})
			err = true
		} else {
			err = false
			fmt.Println("ERROR NO ES UNA ENTRADA VALIDA: ", cad)
			break
		}
	}
	return err, expresiones
}

func SintaxisArbol(expresiones []Expresion, mapa map[string][]Expresion) []Expresion{
	aux := true
	for aux {
		for a, b := range expresiones {
			for c, d := range mapa {
				if b.Dato == c {
					partA := make([]Expresion, 0)
					partB := make([]Expresion, 0)
					partA = append(partA, expresiones[0:a]...)
					partB = append(partB, expresiones[a + 1:len(expresiones)]...)
					partA = append(partA, d...)
					fmt.Println(partB)
					if len(partB) > 0 {
						partB[0].nombreArbol = "\n"
					}
					partA = append(partA, partB...)
					partA[a].nombreArbol = "\n" + c + " :="
					expresiones = make([]Expresion, len(partA))
					copy(expresiones, partA)
					return expresiones
				}
			}
		}
		return expresiones
	}

	return expresiones
}

func buscarVariable(expresiones []Expresion) bool{
	varExp, _ := regexp.Compile(`[A-Z]$`)

	for _, b := range expresiones {
		if varExp.MatchString(b.Dato) {
			return true
		}
	}

	return false
}

func EvaluarProfundidad(t *Arbol) bool{
	if t == nil{
		return false
	}

	if EvaluarOp(t.Valor.Dato){
		if t.Izquierda == nil{
			return false
		} else if t.Derecha == nil {
			return false
		}
	}
	return true;
}

func Resultado(t *Arbol) int64{
	if t == nil {
		return 0
	}

	Resultado(t.Izquierda)
	if EvaluarOp(t.Valor.Dato){
		switch t.Valor.Dato {
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
	a, _ := strconv.ParseInt(t.Valor.Dato, 10, 64)
	return a
}

func RecorrerInorden(t *Arbol) {
	if t == nil {
		return
	}

	RecorrerInorden(t.Izquierda)
	if(t.Valor.nombreArbol != ""){
		fmt.Print(t.Valor.tipoDato, t.Valor.Dato)
	} else {
		fmt.Print(t.Valor)
	}
	RecorrerInorden(t.Derecha)
}

func RecorrerPreorden(t *Arbol) {
	if t == nil {
		return
	}

	fmt.Print(t.Valor, "\n")
	RecorrerPreorden(t.Izquierda)
	RecorrerPreorden(t.Derecha)
}

func ImprimiMenu() {
	fmt.Println("\n---------------------------------------------------------",
		"\nIngrese la opcion (1,...) que desea realizar:",
		"\n1. Ingresar expresion en la cola.",
		"\n2. Construir árbol de la expresion actual en cola.",
		"\n3. Mostrar construccion en innorden del arbol actual en cola.",
		"\n4. Mostrar construccion en preorden del arbol actual en cola.",
		"\n5. Imprimir resultado del árbol.",
		"\n6. Salir del programa.")
}

func ImprimirTabla(){
	fmt.Println("\n---------------------------------------------------------",
		"\n\t\tTABLA DE LENGUAJE",
		"\n\t\tNumero \t\t [0-9]",
		"\n\t\tVariables \t [A-Z]",
		"\n\t\tOperadores \t [+ - * /]",
		"\n\t\tAsignacion \t [:=]",
		"\n---------------------------------------------------------")
}

func main() {
	ciclo := true
	sc:= bufio.NewScanner(os.Stdin)
	var t1 *Arbol
	sentencias := nuevaCola()
	m := make(map[string][]Expresion)

	for ciclo {
		ImprimiMenu()
		sc.Scan()

		switch sc.Text() {
			case "1":
				ImprimirTabla()
				fmt.Println("\nIngrese la expresion a evaluar en preorden, separado por espacios (ex. := A + 2 2), ")
				sc.Scan()
				if err, _ := EvaluarLenguaje(sc.Text()); err {
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
					err, expresiones := EvaluarLenguaje(sentencias.eliminar())

					if err{
						aux := true
						m[expresiones[1].Dato] = SintaxisArbol(expresiones[2:], m)
						for aux {
							if buscarVariable(m[expresiones[1].Dato]){
								m[expresiones[1].Dato] = SintaxisArbol(m[expresiones[1].Dato], m)
							} else {
								aux = false
							}
						}
						m[expresiones[1].Dato][0].nombreArbol = expresiones[1].Dato + ":="

						fmt.Print("\n---------------------------------------------------------",
							"\n\nSe construira el siguiente árbol:", "\n", m[expresiones[1].Dato], "\n")
						t1 = ConstruirArbol(m[expresiones[1].Dato])
						fmt.Println("\nARBOL CREADO.")
					}
				}
			case "3":
				fmt.Print("\n---------------------------------------------------------",
					"\n\nLa siguiente es la construccion en INNORDEN del árbol:\n")
				RecorrerInorden(t1)
			case "4":
				fmt.Print("\n---------------------------------------------------------",
					"\n\nLa siguiente es la construccion en PREORDEN del árbol:\n")
				RecorrerPreorden(t1)
			case "5":
				fmt.Print("\n---------------------------------------------------------",
					"\n\nSe ha realizado la siguiente operacion:\n")
				RecorrerInorden(t1)
				fmt.Println("\nEl resultado fue:", Resultado(t1))
			case "6":
				fmt.Print("\n---------------------------------------------------------",
					"\n\nGracias. \nEl aplicativo ha terminado.\n",
					"\n---------------------------------------------------------\n",)
				ciclo = false
			default:
				fmt.Print("\n---------------------------------------------------------",
					"\n\nERROR: OPCION INVALIDA - \nPor favor ingrese una opcion correcta\n")
		}
	}
}
