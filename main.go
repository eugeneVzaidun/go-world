// init main package
package main

import (
	"strings"
)

// import fmt package

// main function
func main() {
	// create new helper 
	helper := NewHelper("Eugene" ,37)
	helper.PrintInfo()
	helper.PrintAge()
	helper.PrintName()
	// send request to server by using sendRequest function
	// init request body with body as io.Reader
	body:= strings.NewReader(`{"name":"Eugene","age":37}`)
	
	res, error := sendRequest("http://localhost:1880/go-world", "POST", body)
	if error != nil {
		panic(error)
	}
	
	// declare response  as array of bytes
	var response []byte
	response, _ = parseResponse(res)

	// print response
	println(string(response))
	
	// use ReadHelpers function and print result
	// loop over result of ReadHelpers function
	for _, helper := range ReadHelpers() {
		helper.PrintInfo()
	}
}
