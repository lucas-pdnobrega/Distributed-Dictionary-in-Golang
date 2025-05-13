package server

import (
	"fmt"
	"sync"
	"net/http"
	"github.com/hprose/hprose-go/hprose"
)

var (
	dict := make(map[string]string)
	mutex   sync.Mutex
)

type dictService struct{}

func (dictService) update ( string chave , int valor): bool {

}

func (dictService) remove ( string chave ): bool {

}

func (dictService) get ( string chave ):int {

}

func main() {
	valor = 0
	service := hprose.NewHttpService()
	service.AddMethods(myService{})
	http.ListenAndServe(":8080", service)
}
