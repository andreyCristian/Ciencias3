package main

import "fmt"
//import "strings"

type Nodo struct{
    nombre, apellido, eps string
    cedula int
}

type Fifo struct{
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

func (q *Fifo) insertarNodo(nombre string, apellido string, eps string, num int){
    q.insertar(&Nodo{nombre, apellido, eps, num})
}

func (q *Fifo) buscarEPS(eps string) {

    for i := 0; i < len(q.nodos); i++ {
        if(eps == q.nodos[i].eps){
            fmt.Println(q.nodos[i])
        }
    }
}

func (q *Fifo) eliminar() {
    if q.cont != 0 {
        nodes := make([]*Nodo, len(q.nodos) - 1)
        fmt.Print(q.nodos[0].nombre)
        copy(nodes, q.nodos[1:])
        q.nodos = nodes
        q.cont--
    }
}

func (q *Fifo) imprimir(aux int) {
    if aux < q.cont {
        fmt.Println(q.nodos[aux].nombre, q.nodos[aux].apellido, q.nodos[aux].eps, q.nodos[aux].cedula)
        aux++
        q.imprimir(aux)
    }
}

func main(){
    q := nuevaCola()
    ciclo := 1
    var aux int
    var a string

    for ciclo < 1000 {
        fmt.Println("Ingrese la opcion (1,...) que desea realizar:",
            "\n1. Construir cola.",
            "\n2. Atender usuario de la cola.",
            "\n3. Examinar usuarios que se encuentran en una eps.",
            "\n4. Mostrar todos los usuarios de la cola.",
            "\n5. Salir del programa." )

        fmt.Scanln(&aux)

        switch aux {
            case 1:
                q.insertarNodo("Cristian", "Sossa", "Sanitas", 123456789)
                q.insertarNodo("Santiago", "Ruiz", "Sanitas", 987654321)
                q.insertarNodo("Andres", "Gonzalez", "Cafesalud", 192837465)
                q.insertarNodo("Alejandro", "Murcia", "Compensar", 564738291)
                fmt.Println("Cola construida.........")
                fmt.Println("------------------------------------------------------------------------------")
                break
            case 2:
                if q.cont == 0 {
                    q = nuevaCola()
                        fmt.Println("La cola actualmente se encuentra vacia.")
                        fmt.Println("------------------------------------------------------------------------------")
                } else {
                        fmt.Println("Se ha atendido a ")
                        q.eliminar()
                        fmt.Println("------------------------------------------------------------------------------")
                }
            case 3:
                fmt.Println("Ingrese la eps a buscar: ")
                fmt.Scanln(&a)
                q.buscarEPS(a)
                fmt.Println("------------------------------------------------------------------------------")
            case 4:
                fmt.Println("Usuarios registrados en el sistema:\n")
                q.imprimir(0)
                fmt.Println("------------------------------------------------------------------------------")
            default:
                fmt.Println("Gracias por utilizar el sistema")
                ciclo = 1000
        }
    }
}
