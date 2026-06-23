package main

import "fmt"
import "golang.org/x/example/hello/reverse"

func main(){
	fmt.Println(reverse.String("Hello, world!"), reverse.Int(12345))
}