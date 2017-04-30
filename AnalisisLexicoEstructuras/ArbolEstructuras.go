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
	Valor Expresion
	Derecha *Arbol
}

type Nodo struct {
	expresion string
}

type Expresion struct {
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

func Diccionario(valor string) string{
	numExp, _ := regexp.Compile(`-?\d$`)
    opExp, _ := regexp.Compile(`[-]|[+]|[*]|[/]|[:=]$`)
    varExp, _ := regexp.Compile(`[A-Z]$`)

	if numExp.MatchString(valor){
		return "numero"
	} else if opExp.MatchString(valor) {
		return "operador"
	} else if varExp.MatchString(valor) {
		return "variable"
	}

	return ""
}

func EvaluarLenguaje(cadena string) (bool, []Expresion){
    numExp, _ := regexp.Compile(`-?\d$`)
    opExp, _ := regexp.Compile(`[-]|[+]|[*]|[/]|[:=]$`)
    varExp, _ := regexp.Compile(`[A-Z]$`)
	err := false
    expresiones := make([]Expresion, 0)

	for _, cad := range Separar(cadena) {
		if numExp.MatchString(cad){
            expresiones = append(expresiones, Expresion{tipoDato: "Numero\t->\t", Dato: cad})
			err = true
		} else if opExp.MatchString(cad) {
            expresiones = append(expresiones, Expresion{tipoDato: "Operador\t->\t", Dato: cad})
			err = true
		} else if varExp.MatchString(cad) {
            expresiones = append(expresiones, Expresion{tipoDato: "Variable\t->\t", Dato: cad})
			err = true
		} else {
			err = false
			fmt.Println("ERROR NO ES UNA ENTRADA VALIDA: ", cad)
			break
		}
	}
    return err, expresiones
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

func CadenaFinal(cadena []Expresion, mapa map[string]int64) ([]Expresion, bool){

	for a, b := range cadena {
		for c, d := range mapa {
			if b.Dato == c {
				cadena[a].Dato = strconv.FormatInt(d, 10)
			}
		}
	}

	for _, b := range cadena[:len(cadena) - 2]{
		if aux := Diccionario(b.Dato); aux == "variable" {
			fmt.Println("Por favor inserte los arboles correctos <<", b, ">> \nERROR: DATO NO ENCONTRADO ")
			return cadena, false
			break
		}
	}

	return cadena, true
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
	fmt.Print(t.Valor, "\n")
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
		"\n3. Ingresar expresion final del árbol. (ex. + A B := C)",
		"\n4. Mostrar construccion en innorden del arbol actual en cola.",
		"\n5. Mostrar construccion en preorden del arbol actual en cola.",
		"\n6. Imprimir resultado del árbol.",
		"\n7. Salir del programa.")
}

func ImprimirTabla(){
	fmt.Println("\n---------------------------------------------------------",
		"\n\t\tTABLA DE LENGUAJE",
		"\n\t\tNumero \t\t [0-9]",
		"\n\t\tVariables \t [A-Z]",
		"\n\t\tOperadores \t [+ - * / :=]",
		"\n---------------------------------------------------------")
}

func main() {
    ciclo := 0
	sc:= bufio.NewScanner(os.Stdin)
	var t1 *Arbol
	var tFinal *Arbol
	sentencias := nuevaCola()
	m := make(map[string]int64)

	for ciclo < 1000 {
		ImprimiMenu()
		sc.Scan()

		switch sc.Text() {
			case "1":
				ImprimirTabla()
				fmt.Println("\nIngrese la expresion a evaluar en preorden, separado por espacios (ex. + 2 2 := A), ")
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
                    _, expresiones := EvaluarLenguaje(sentencias.eliminar())
					fmt.Print("\n---------------------------------------------------------",
								"\n\nSe construira el siguiente árbol: \n", expresiones[:len(expresiones) - 2], "\n")
					t1 = ConstruirArbol(expresiones)
					fmt.Println("\nARBOL CREADO.")
					m[expresiones[len(expresiones) - 1].Dato] = Resultado(t1)
				}
			case "3":
				ImprimirTabla()
				fmt.Println("\nAsegurese de haber creado los arboles correspondientes antes (ex. + A B := C)")
				sc.Scan()
				_, expresiones := EvaluarLenguaje(sc.Text())
				expresiones, err := CadenaFinal(expresiones, m)
				if err {
					tFinal = ConstruirArbol(expresiones)
					fmt.Println("\nARBOL CREADO. \nEl resultado del arbol creado es:", Resultado(tFinal))
				}
			case "4":
				fmt.Print("\n---------------------------------------------------------",
						"\n\nLa siguiente es la construccion en INNORDEN del árbol:\n")
				RecorrerInorden(t1)
			case "5":
				fmt.Print("\n---------------------------------------------------------",
						"\n\nLa siguiente es la construccion en PREORDEN del árbol:\n")
				RecorrerPreorden(t1)
			case "6":
				fmt.Print("\n---------------------------------------------------------",
						"\n\nSe ha realizado la siguiente operacion:")
				RecorrerInorden(t1)
				fmt.Println("\nEl resultado fue:", Resultado(t1))
			case "7":
				fmt.Print("\n---------------------------------------------------------",
						"\n\nGracias. \nEl aplicativo ha terminado.\n",
						"\n---------------------------------------------------------\n",)
				ciclo += 1000
			default:
				fmt.Print("\n---------------------------------------------------------",
						"\n\nERROR: OPCION INVALIDA - \nPor favor ingrese una opcion correcta\n")
		}
	}
}
