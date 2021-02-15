package main

import (
	"fmt"
	"net/http"	
)

func case1(res http.ResponseWriter, req *http.Request) {
	fmt.Fprint(res, "This is case1")
}

func case2(res http.ResponseWriter, req *http.Request) {
	fmt.Fprint(res, "This is case2")
}

func main() {

	http.HandleFunc("/case1", AddSomethingInMiddle(case1))
	http.HandleFunc("/case2", case2)
	
	http.ListenAndServe(":1234", nil)
}

// AddSomethingInMiddle Defines the logic inside the middleware
func AddSomethingInMiddle(hf http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request){
		// Do some additional logic in middle
		fmt.Println("Do some addition in middle")
		hf.ServeHTTP(res, req)
	})
}
