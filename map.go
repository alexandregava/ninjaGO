package main

import "fmt"

func main() {

	mapa := map[string][]string{
		"nome":  {"ale", "mari"},
		"sobre": {"gava", "franco"},
	}

	mapa["hobby"] = []string{"teste1", "teste2"} //adiciona item

	delete(mapa, "sobre") //deleta item

	for k, v := range mapa {
		fmt.Println(k)
		for _, v2 := range v {
			fmt.Println("\t", v2)
		}
	}

}
