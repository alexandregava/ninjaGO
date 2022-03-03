package main

import "fmt"

//Chamando funcao dentro de funcao
func main() {
	resultado := recebeValores(soma, []int{1, 2, 3, 4, 5, 6, 7, 8})
	fmt.Println(resultado)

}

func soma(slice ...int) int {
	total := 0
	for _, v := range slice {
		total += v
	}
	return total
}

func recebeValores(funcao func(x ...int) int, valores []int) int {
	var numImp []int
	for _, v := range valores {
		if v%2 != 0 {
			numImp = append(numImp, v)
		}
	}
	somaImpares := funcao(numImp...)
	return somaImpares
}
