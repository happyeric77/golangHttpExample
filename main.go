package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Make a foo int type which immplements ServeHTTP method to be a http.Handler interface.
type foo int

func (f foo) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	fmt.Println("url: ", req.URL)
	fmt.Println("reqeust method", req.Method)

	// Header type is map[string]string
	for key := range req.Header {
		fmt.Println(key, ":", req.Header[key])
	}

	// Set the header
	res.Header().Set("Content-Type", "application/json")
	fmt.Println(res.Header())

	// Write the status code to header
	res.WriteHeader(200)

	// Write something into Body
	data := struct {
		Foo string `json:"foo"`
	}{Foo: "A foo response"}

	rawJSON, _ := json.Marshal(data)

	_, err := res.Write(rawJSON)
	if err != nil {
		fmt.Println(err)
	}
	
}

// Pass the foo as http.Handler into ListenAndServe()
func main() {
	var f foo
	http.ListenAndServe(":1234", f)
}
