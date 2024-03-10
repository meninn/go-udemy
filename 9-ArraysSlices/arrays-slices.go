package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Arrays e Slices")

	// Arrays são estruturas homogêneas (mesmo tipo) e estáticos (fixos)
	var array1 [5]int
	fmt.Println(array1)

	array2 := [5]string{"Posição 1", "Posição 2", "Posição 3", "Posição 4", "Posição 5"}
	fmt.Println(array2)

	array3 := [...]int{1, 2, 3, 4, 5} // Compilador conta
	fmt.Println(array3)

	// Slices são estruturas homogêneas e dinâmicas
	slice := []int{10, 11, 12, 13, 14, 15, 16, 17, 18, 19}
	fmt.Println(slice)

	fmt.Println(reflect.TypeOf(array3))
	fmt.Println(reflect.TypeOf(slice))

	slice = append(slice, 20)
	fmt.Println(slice)

	slice2 := array2[1:3] // Pegando do índice 1 até o 3 (sem incluir o 3)
	fmt.Println(slice2)

	// Arrays Internos
	fmt.Println("-----------------------")
	slice3 := make([]float32, 10, 11)
	fmt.Println(slice3)
	fmt.Println(len(slice3)) // Tamanho do slice
	fmt.Println(cap(slice3)) // Tamanho do array interno

	slice3 = append(slice3, 5)
	slice3 = append(slice3, 6) // Cria um novo array interno
	fmt.Println(slice3)
	fmt.Println(len(slice3)) // Tamanho do slice
	fmt.Println(cap(slice3)) // Tamanho do array interno

	slice4 := make([]float32, 5)
	fmt.Println(slice4)
	fmt.Println(len(slice4)) // Tamanho do slice
	fmt.Println(cap(slice4)) // Tamanho do array interno

	slice4 = append(slice4, 5)
	slice4 = append(slice4, 6) // Cria um novo array interno
	fmt.Println(slice4)
	fmt.Println(len(slice4)) // Tamanho do slice
	fmt.Println(cap(slice4)) // Tamanho do array interno

}
