package main

import (
	"fmt"
	"net/http"	
)


// Use http.HandleFunc to implement multiplexer (mux) routing.

func case1(res http.ResponseWriter, req *http.Request) {
	fmt.Fprint(res, "This is case1")
}

func case2(res http.ResponseWriter, req *http.Request) {
	fmt.Fprint(res, "This is case2")
}

func main() {

	http.HandleFunc("/case1", case1)
	http.HandleFunc("/case2", case2)
	
	http.ListenAndServe(":1234", nil)
}
