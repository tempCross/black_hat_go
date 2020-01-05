package main

import (
	"fmt"
	"net/http"
)

func main(){
req, err := http.NewRequest("DELETE", "https://www.google.com/robots.txt", nil)
var client http.Client
resp, err := client.Do(req)
if err != nil {
 		fmt.Println(err)
 	}
respHeader := resp.Header
fmt.Printf("Header: %v\n", respHeader) //<-- here !
}
