package main

import (
	"net/http"
	"fmt"
	"log"
)

type Hello struct {

}

func (h Hello) ServerHTTP(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w,"hello!")
}

func main() {
	var h Hello
	err:=http.ListenAndServe("localhost:4001",h)
	if err!=nil {
		log.Fatal(err)
	}
}

