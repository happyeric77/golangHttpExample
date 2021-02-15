package main

import (
	"fmt"
	"net/http"	
)


// Use http.Handler to check URL path to implement multiplexer (mux)
type foo int 
func(f foo) ServeHTTP (res http.ResponseWriter, req *http.Request){
	path := req.URL.Path
	switch path {
	case "/case1":
		fmt.Fprint(res, "This is case1")
	case "/case2":
		fmt.Fprint(res, "This is case2")
	default:
		fmt.Fprintf(res, "Server does not support this endpoint(%s).\nServer only supports '/case1' & '/case2'", path)
	}
}

func main() {
	http.ServeMux
	var f foo
	http.ListenAndServe(":1234", f)
}
