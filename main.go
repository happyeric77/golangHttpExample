package main

import (
	"fmt"
	"net/http"	
)

func main() {
	// Pass the end point and handler function
	http.HandleFunc("/", func(http.ResponseWriter, *http.Request){
		fmt.Println("Default http routing function")
	})
	http.ListenAndServe(":1234", nil)
}
