package main

import (
	"fmt"
	"os"
	"github.com/hprose/hprose-go/hprose"
)

type clientStub struct{
	Update func ( string, int) bool 
	Remove func ( string ) bool
	Get func ( string ) int
}

func main() {
	client := hprose.NewClient("http://127.0.0.1:8080/")
	var ro *clientStub
	var opcao int
	var chave string
	var n int
	client.UseService(&ro)

	for {
		fmt.Printf("\nEscolha uma opção :\n1 - Get\n2 - Atualizar\n3 - Remover\n4 - Sair\n")
		fmt.Scan(&opcao)

		switch opcao {
			case 1:
				
				fmt.Println("Digite a chave a consultar: ")
				fmt.Scan(&chave)

				val := ro.Get(chave)
				
				if val == -1 {
					fmt.Println("-> Chave não encontrada.")
				} else {
					fmt.Printf("-> Valor armazenado: %d\n", val)
				}

			case 2:
				
				fmt.Println("-> Digite a chave para atualizar: ")
				fmt.Scan(&chave)
				fmt.Println("-> Insira seu novo valor: ")
				fmt.Scan(&n)
				
				ok := ro.Update(chave, n)
				if ok {
					fmt.Println("-> Valor atualizado.")
				} else {
					fmt.Println("-> Falha: chave inexistente.")
				}

			case 3:
				
				fmt.Println("Insira uma chave")
				fmt.Scan(&chave)
				ok := ro.Remove(chave)
				if ok {
					fmt.Println("-> Valor removido.")
				} else {
					fmt.Println("-> Falha: chave inexistente.")
				}
			
			case 4:
				os.Exit(0)

			default:
				fmt.Println("-> Falha: Opção inválida.")
		}
	}
}