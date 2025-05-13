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
	_, exists := dict[chave]
	if (!exists) {
		return false
	}

	dict[chave] = valor
	defer mutex.Unlock()

	return true
}

func (dictService) Remove ( chave string ) bool {

	mutex.Lock()
	_, exists := dict[chave]
	if (!exists) {
		return false
	} else {
		delete(dict, chave)
	}
	defer mutex.Unlock()

	return true
}

func (dictService) Get ( chave string ) int {

	mutex.Lock()
	res, exists := dict[chave]
	defer mutex.Unlock()

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
