package main

import (
	"errors"
	"fmt"
)

func main() {
	var numero int64 = -100000000000000000
	fmt.Println(numero)

	var numero2 uint64 = 100000000000000000
	fmt.Println(numero2)

	// alias
	// INT32 = RUNE
	var numero3 rune = 123456
	fmt.Println(numero3)

	// INT8 = BYTE
	var numero4 byte = 123
	fmt.Println(numero4)

	// float32, float64
	var numeroReal1 float32 = 123.45
	fmt.Println(numeroReal1)

	var numeroReal2 float64 = 123
	fmt.Println(numeroReal2)

	// boolean
	var booleano1 bool = true
	fmt.Println(booleano1)

	var booleano2 bool = false
	fmt.Println(booleano2)

	// string
	var cadeiaDeCaracteres string = "Texto"
	fmt.Println(cadeiaDeCaracteres)

	// char não existe em Go
	caractere := 'B'
	fmt.Println(caractere)

	var textoDefault string
	fmt.Println(textoDefault) // valor zero

	var numeroDefault int16
	fmt.Println(numeroDefault) // valor zero

	var booleanoDefault bool
	fmt.Println(booleanoDefault) // valor zero

	var erro error = errors.New("Erro interno")
	fmt.Println(erro) // valor zero
}
