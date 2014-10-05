package main

import (
	"os"
	"fmt"
)

type Reader interface {
	Read(b []byte)(n int,err error)
}

type Writer interface {
	Write(b []byte)(n int,err error)
}

type ReadWiter interface {
	Reader
	Writer
}

func main(){
	var w Writer

	w = os.Stdout

	fmt.Fprintf(w,"hello,world!")
}
