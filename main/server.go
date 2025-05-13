package main

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/hprose/hprose-go/hprose"
)

var (
	dict map[string]int = make(map[string]int);
	mutex   sync.Mutex
	res int
)

type dictService struct{}

func (dictService) Update ( chave string, valor int) bool {

	mutex.Lock()
	defer mutex.Unlock()

	_, exists := dict[chave]
	dict[chave] = valor

	return exists
}

func (dictService) Remove ( chave string ) bool {

	mutex.Lock()
	defer mutex.Unlock()

	_, exists := dict[chave]
	delete(dict, chave)

	return exists
}

func (dictService) Get ( chave string ) int {

	mutex.Lock()
	defer mutex.Unlock()

	res, exists := dict[chave]

	if (!exists) {
		return -1
	} else {
		return res
	}
}

func main() {
	service := hprose.NewHttpService()
	service.AddMethods(dictService{})
	http.ListenAndServe(":8080", service)
	fmt.Println("Server running on localhost:8080")
}
