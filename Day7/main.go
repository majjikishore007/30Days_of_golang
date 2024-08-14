package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func Greet(writter io.Writer,name string) {
	// fmt.Printf("Hello, %s", name)
	fmt.Fprintf(writter,"Hello %s",name)
}

func MyGreetHanlder(w http.ResponseWriter, r *http.Request)  {
	Greet(w,"World")
}

func main()  {
	Greet(os.Stdout,"kishore")
	http.ListenAndServe(":3000",http.HandlerFunc(MyGreetHanlder))
}