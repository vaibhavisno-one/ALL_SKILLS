package main

import "fmt"
import "net/http"
import "io/ioutil"

const url="https://vbhv-six.vercel.app"


func main(){
	fmt.Println("HTTP web request")

	response, err:= http.Get(url)

	if err != nil{
		panic(err)
	}

	fmt.Printf("Response is of type :  %T\n", response)

	defer response.Body.Close()  // callers responsibility to close the response

	databytes,err:=ioutil.ReadAll(response.Body)

	if err!=nil{
		panic(err)
	}

	content:= string(databytes)

	fmt.Println(content)
}