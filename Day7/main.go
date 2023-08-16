package main

import (
	"fmt"
	"io"
	"os"
)


func Greet(writer io.Writer, name string) {
	// fmt.Printf("Hello, %s", name)
	fmt.Fprintf(writer,"Hello, %s",name)
}

func main()  {
	Greet(os.Stdout,"majj")
}