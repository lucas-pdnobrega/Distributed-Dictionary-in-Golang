package main

import (
	"fmt"
	"sync"
	"net/http"
	"github.com/hprose/hprose-go/hprose"
)

var (
	dict map[string]int = make(map[string]int);
	mutex   sync.Mutex
	res int
)

type dictService struct{}

func (dictService) update ( chave string, valor int) bool {

	mutex.Lock()
	res, exists := dict[chave]
	if (!exists) {
		return false
	}

	dict[chave] = valor
	mutex.Unlock()

	return true
}

func (dictService) remove ( chave string ) bool {

	mutex.Lock()
	res, exists := dict[chave]
	if (!exists) {
		return false
	} else {
		delete(dict, chave)
	}
	mutex.Unlock()

	return true
}

func (dictService) get ( chave string ) int {

	mutex.Lock()
	res, exists := dict[chave]
	mutex.Unlock()

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
}
