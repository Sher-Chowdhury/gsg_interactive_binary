package main

// need to first run: go get github.com/fatih/color
import (
	"bufio" //https://golang.org/pkg/bufio/
	"fmt"
	"os"
)

func main() {
	var f *os.File // "f" is a pointer to memory location that will hold a variable of the type os.File.
	// os.File is a path to a file on the fileystem
	f = os.Stdin // here we chose the Stdin as the the filepath to use. In Linux, stdin, stdout, and stderr are all essentially files behind the scenese

	defer f.Close() // https://tour.golang.org/flowcontrol/12
	// defer means that this line is triggered once the main function finishes running.
	// we are applying the Close method for the f object.

	// This is a bit like opening the file to start writing to it
	scanner := bufio.NewScanner(f)

	// https://golang.org/pkg/bufio/#Scanner.Scan
	for scanner.Scan() {
		fmt.Println(">>>", scanner.Text())
  }
  
  // you need to do ctrl+d to escape this for-loop

}
