package main

// need to first run: go get github.com/fatih/color 
import (
  "io"
  "os"
)

func main() {
  myMessage := ""
  arguments := os.Args
  // https://golang.org/pkg/builtin/#len
  if len(arguments) == 1 {
    myMessage = "No arguements provided"
  } else {
    myMessage = arguments[1]
  }

  // https://golang.org/pkg/io
  // io.WriteString() works in a similar way to fmt.Print(). the main difference is that fmt.Print() can take multiple parameters, e.g. fmt.Print(var1, "some stuff", var2, "even more stuff", "etc") 
  // wheerase io.WriteString() can only take exactly 2 params, where to write to, and what to write.  
  io.WriteString(os.Stdout, myMessage)
  io.WriteString(os.Stdout, "\n")

}
