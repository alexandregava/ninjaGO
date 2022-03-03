package main

import (
	"fmt"
)

type pessoa struct {
	nome  string
	idade int
}

// Sintaxe para criacao de funcao:
// func (receiver) identifier(parameters) (returns) { code }

func (p pessoa) oibomdia() {
	fmt.Println(p.nome, "diz bom dia!")
}

func main() {

	mauricio := pessoa{"Maurício", 30}
	mauricio.oibomdia()

}
