package main

import "fmt"

type usuario struct {
	nome     string
	idade    uint8
	endereco endereco
}

type endereco struct {
	logradouro string
	numero     uint8
}

func main() {
	var u usuario

	u.nome = "Gabriel"
	u.idade = 24

	fmt.Println(u)

	u2 := usuario{"Gabriel", 24, endereco{"Ingleses", 123}}
	fmt.Println(u2)

	u3 := usuario{nome: "Gabriel", idade: 24, endereco: endereco{logradouro: "Ingleses", numero: 123}}
	fmt.Println(u3)
}
