package main

import "fmt"

import "example.com/greetings"

func main(){
	message:=greetings.Hello("Gladys")  // greetings.hello means that greeting is package and hello is function
	fmt.Println(message)
}

