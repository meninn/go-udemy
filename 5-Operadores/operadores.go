package main

import "fmt"

func main() {
	// Operadores aritméticos
	soma := 1 + 2
	subtracao := 1 - 2
	divisao := 10 / 4
	multiplicacao := 10 * 5
	restoDaDivisao := 10 % 2

	fmt.Println(soma, subtracao, divisao, multiplicacao, restoDaDivisao)

	// Operadores de atribuição
	var variavel1 string = "String"
	variavel2 := "String 2"

	fmt.Println(variavel1, variavel2)

	// Operadores relacionais
	fmt.Println(1 > 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 == 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 < 2)
	fmt.Println(1 != 2)

	// Operadores lógicos
	verdadeiro, falso := true, false
	fmt.Println(verdadeiro && falso)
	fmt.Println(verdadeiro || falso)
	fmt.Println(!verdadeiro)

	// Operadores unários
	numero := 10
	numero++
	numero += 15 // numero = numero + 15
	fmt.Println(numero)

	numero--
	numero -= 20 // numero = numero - 20
	fmt.Println(numero)

	numero *= 3  // numero = numero * 3
	numero /= 10 // numero = numero / 10
	numero %= 3  // numero = numero % 3
	fmt.Println(numero)

	// Operadores ternários
	// Não existe em Go

}
