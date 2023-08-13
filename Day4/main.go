package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

var ErrNoPath = errors.New("path required")
func execInput(input string) error{
	// remove the needw line char 
	input=strings.TrimSuffix(input,"\n")

	// split the input to seperate the args
	args:=strings.Split(input," ")

	switch args[0] {
    case "cd":
        // 'cd' to home with empty path not yet supported.
        if len(args) < 2 {
            return ErrNoPath
        }
        // Change the directory and return the error.
        return os.Chdir(args[1])
    case "exit":
        os.Exit(0)
    }

	// ... spread opreator .... mean any number of trailling aruguments
	cmd := exec.Command(args[0],args[1:]...)

	cmd.Stderr=os.Stderr
	cmd.Stdout=os.Stdout

	return cmd.Run()
}

func main()  {
	reader := bufio.NewReader(os.Stdin)
	
	for {
		input,err := reader.ReadString('\n')
		if err != nil{
			fmt.Fprintln(os.Stderr,err)
		}

		if err=execInput(input); err !=nil{
			fmt.Fprintln(os.Stderr,err)
		}
	}
}

