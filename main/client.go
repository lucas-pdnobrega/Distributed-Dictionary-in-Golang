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
		fmt.Printf(`
			Escolha uma opção :

			1 - Get
			2 - Atualizar
			3 - Remover
			4 - Sair
		`)
		fmt.Scan(&opcao)

		switch opcao {
			case 1:
				
				fmt.Println("Insira uma chave")
				fmt.Scan(&chave)
				fmt.Println(ro.Get(chave))

			case 2:

				
				fmt.Println("Insira uma chave")
				fmt.Scan(&chave)
				
				fmt.Println("Insira um número")
				fmt.Scan(&n)
				
				fmt.Println(ro.Update(chave, n))

			case 3:
				
				fmt.Println("Insira uma chave")
				fmt.Scan(&chave)
				
				fmt.Println("Insira um número")
				fmt.Println(ro.Remove(chave))
			
			case 4:
				os.Exit(0)
		}
	}
}