package main

import (
	"encoding/json"
	"fmt"
)

type Pessoa struct {
	Nome      string `json:"Apelido"`
	Sobrenome string
	Idade     int
}

func main() {
	pe1 := Pessoa{"Ale", "Gava", 33}

	p, err := json.Marshal(pe1)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(p))
}
