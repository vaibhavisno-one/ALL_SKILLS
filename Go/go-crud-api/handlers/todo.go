package handlers

import "fmt"

func Profile(w http.ResponseWriter, r *http.request){
	fmt.Println("Profile handler executed")
}