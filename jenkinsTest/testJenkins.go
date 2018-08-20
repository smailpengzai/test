package main

import "fmt"

func main(){
	// Initialize Jenkins API
	//
	// For example:
	  jenkinsApi := Init(&Connection {
	    Username: "luopeng",
	    AccessToken: "11b1e3c296e0b59e5d2bcfa37709469658",
	    BaseUrl: "http://10.94.10.18:8080",
	  })
	//http://10.94.10.18:8080/job/test/api/json
	fmt.Println(jenkinsApi.GetJob("test"))

}
