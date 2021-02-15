package main

import (
	"fmt"
	"net/http"	
)


// Use http.ServeMux to implement multiplexer (mux) routing.

type case1 int

type case2 int

func (c case1) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	fmt.Fprint(res, "This is case1")
}

func (c case2) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	fmt.Fprint(res, "This is case2")
}


func main() {
	mux := http.NewServeMux()
	var c1 case1
	var c2 case2
	mux.Handle("/case1", c1)
	mux.Handle("/case2", c2)
	
	http.ListenAndServe(":1234", mux)
}
