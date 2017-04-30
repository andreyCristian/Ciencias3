package main

import (
	"fmt"
	"strconv"
	"strings"
	"bufio"
	"os"
)

type Arbol struct {
	Izquierda  *Arbol
	Valor string
	Derecha *Arbol
}

func ConstruirArbol(sentencia string) *Arbol{
	cadena := Separar(sentencia)
	var t *Arbol

	for i := 0; i < len(cadena); i++ {
		t = Insertar(t, cadena[i])
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

func RecorrerPreorden(t *Arbol) {
	if t == nil {
		return
	}

	fmt.Print(t.Valor)
  	fmt.Print(" ")
	RecorrerPreorden(t.Izquierda)
	RecorrerPreorden(t.Derecha)
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

func ImprimiMenu() {
	fmt.Println("\n---------------------------------------------------------",
		"\nIngrese la opcion (1,...) que desea realizar:",
		"\n1. Insertar un Arbol.",
		"\n2. Mostrar construccion en innorden.",
		"\n3. Mostrar construccion en preorden.",
		"\n4. Imprimir resultado del árbol.",
		"\n5. Salir del programa.")
}

func main() {
	ciclo := 0
	sc:= bufio.NewScanner(os.Stdin)
	var t1 *Arbol

	for ciclo < 1000 {
		ImprimiMenu()
		sc.Scan()

		switch sc.Text() {
			case "1":
				fmt.Println("\nIngrese el árbol en preorden, separado por espacios (ex. + 2 2): ")
				sc.Scan()
				t1 = ConstruirArbol(sc.Text())
				fmt.Println("\nARBOL CREADO.")
			case "2":
				fmt.Println("\nLa siguiente es la operacion que se va a realizar (innorden): \n")
				RecorrerInorden(t1)
			case "3":
				fmt.Println("\nLa siguiente es la construccion en preorden del árbol: \n")
				RecorrerPreorden(t1)
			case "4":
				fmt.Print("\nSe ha realizado la siguiente operacion: ")
				RecorrerInorden(t1)
				fmt.Println("\nEl resultado fue:", Resultado(t1))
			case "5":
				fmt.Println("\nGracias. \nEl aplicativo ha terminado.")
				ciclo += 1000
			default:
				fmt.Println("\nLa opcion ingresada no fue correcta.")
		}
	}
}
