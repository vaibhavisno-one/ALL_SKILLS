package main

import "fmt"
import "log"

import "example.com/greetings"

func main(){

	// set the properties of predefined logger
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// message,err:=greetings.Hello("vaibhav")  // greetings.hello means that greeting is package and hello is function
	// if there was error the print log and exit the program



	// A slice of names.
    names := []string{"Gladys", "Samantha", "Darrin"}

    // Request greeting messages for the names.
    messages, err := greetings.Hellos(names)
	if err!=nil{
		log.Fatal(err)
	}
	// If no error was returned, print the returned map of
    // messages to the console.
    fmt.Println(messages)
}

