package greetings

import "fmt"

func Hello(name string) string{

	//return the greeting message

	message:=fmt.Sprintf("Hi, %v. Welcome!", name)
	return message

	
}

