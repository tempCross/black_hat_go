package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"log"
	"os"
)

func main(){
	r1, err := http.Get("http://www.google.com/robots.txt")
	if err != nil {
		log.Fatal(err)
	}
	// Read response body
	defer r1.Body.Close()

	r2, err := ioutil.ReadAll(r1.Body)	
	if err != nil {
 		fmt.Println(err)
 	}

 	// print out
 	fmt.Println(os.Stdout, string(r2)) //<-- here !
}