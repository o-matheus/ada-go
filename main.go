package main

import "fmt"

func main() {
	// Implemente o exercício descrito no README.md
	var opcao int
	var nota float32
	var maiorNota float32
	var menorNota float32
	contador := 0
	notas := make([]float32, 0, 6)
	var totalNotas float32

	for {
		fmt.Println("Digite a opção desejada: ")
		fmt.Println("1 - Adicionar nota")
		fmt.Println("2 - Ver relatório")
		fmt.Println("0 - Sair")

		// Quantidade de itens lidos | Código do erro ou nil
		_, _ = fmt.Scan(&opcao)

		switch opcao {
		case 1:
			fmt.Print("Digite a nota: ")
			_, _ = fmt.Scan(&nota)
			if nota < 0 || nota > 10 {
				fmt.Println("Insira uma nota entre o intervalo de 0 a 10.")
			} else {
				totalNotas += nota
				notas = append(notas, nota)

				if contador == 0 {
					maiorNota = nota
					menorNota = nota
				} else if nota > maiorNota {
					maiorNota = nota
				} else if nota < menorNota {
					menorNota = nota
				}
				contador += 1
				fmt.Println("Nota adicionada com sucesso!")
				fmt.Printf("Nota adicionada: %.2f\n", nota)
			}
		case 2:
			if len(notas) == 0 {
				fmt.Println("Nenhuma nota cadastrada.")

			} else {
				var media = totalNotas / float32(len(notas))
				var resultado string
				if media < 7 {
					resultado = "Reprovado"
				} else {
					resultado = "Aprovado"
				}
				fmt.Println("-----RELATORIO-----")
				for indice, nota := range notas {
					fmt.Printf("Nota %d: %.2f\n", indice+1, nota)
				}
				fmt.Printf("Maior nota: %.2f\n", maiorNota)
				fmt.Printf("Menor nota: %.2f\n", menorNota)
				fmt.Printf("Média:  %.2f\n", media)
				fmt.Printf("Resultado: %s\n", resultado)
			}
		case 0:
			fmt.Println("Até mais.")
			return
		default:
			fmt.Println("Escolha uma opção valida.")
		}
	}
}
