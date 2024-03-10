package main

import "fmt"

func somar(n1 int8, n2 int8) int8 {
	return n1 + n2
}

func calculosMatematicos(n1, n2 int8) (int8, int8) {
	soma := n1 + n2
	subtracao := n1 - n2
	return soma, subtracao
}

func main() {
	resultado := somar(10, 20)
	println(resultado)

	var f = func(txt string) string {
		fmt.Println(txt)
		return txt
	}

	resultadoF := f("Texto da função 1")
	println(resultadoF)

	resultadoSoma, resultadoSubtracao := calculosMatematicos(10, 15)
	println(resultadoSoma, resultadoSubtracao)
}
