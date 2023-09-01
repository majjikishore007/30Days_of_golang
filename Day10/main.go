package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
)

func makeRGB(i int)(int , int, int)  {
	var f= 0.1

	return int(math.Sin(f*float64(i)+0)*127+128),
	int(math.Sin(f*float64(i)+2*math.Pi/3)*127+128),
	int(math.Sin(f*float64(i)+4*math.Pi/3)*127+128)
}
func Printout(output[] rune)  {
	for j:=0 ; j<len(output);j++{
		r,g,b := makeRGB(j)
		fmt.Printf("\033[38;2;%d;%d;%dm%c\033[0m", r, g, b, output[j])
	}
}

func main()  {
	info,_:=os.Stdin.Stat()
	var output []rune

	if info.Mode() & os.ModeCharDevice !=0 {
		fmt.Println("The command is intended to work with pipes.")
        fmt.Println("Usage: fortune | gorainbow")
	}
	reader:= bufio.NewReader(os.Stdin)

	for {
        input, _, err := reader.ReadRune()
        if err != nil && err == io.EOF {
            break
        }
        output = append(output, input)
    }
	Printout(output)
}