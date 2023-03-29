package main

import "fmt"

func main() {

	// Exercícios do SLIDE.

	// 1. Crie um Array de inteiros com 5 elementos e imprima cada elemento em uma linha separada.

	n := [5]int{1, 2, 3, 4, 5}
	for i := 0; i < len(n); i++ {
		fmt.Println(n[i])
	}

	// 2. Crie um Slice de strings vazio, adicione os nomes "João", "Maria" e "José" a ele e imprima o Slice.

	l := []string{}
	l = append(l, "João,", "Maria,", "José")
	fmt.Print(l)

	// 3. Crie um Slice de inteiros com os valores 1, 2, 3, 4 e 5. Remova o terceiro elemento e imprima o Slice resultante.

	numeros := []int{1, 2, 3, 4, 5}
	fmt.Println(numeros)

	numeros = append(numeros[:2], numeros[3:]...)
	fmt.Println(numeros)

	// 4. Crie um Array de floats com 6 elementos e calcule a média dos valores armazenados no Array.

	x := [6]float64{1, 2, 3, 4, 5, 6}
	var soma float64

	for _, numero := range x {
		soma += numero
	}

	media := soma / float64(len(x))
	fmt.Println(media)

	// 5. Crie uma matriz bidimensional de inteiros com 3 linhas e 4 colunas. Inicialize cada elemento com o valor da soma do índice da linha e o índice da coluna. Imprima a matriz resultante.

	var matriz [3][4]int

	for linha := range matriz {
		for coluna := range matriz[linha] {
			matriz[linha][coluna] = linha + coluna
		}
		fmt.Println(matriz)
	}
}
