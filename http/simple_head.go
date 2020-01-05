package main

import (
	"fmt"
	"net/http"
	"log"
	"io/ioutil"
	"os"
	"strings"
	"net/url"
)	

func main(){
	r2, err := http.Head("http://www.google.com/robots.txt")
	if err != nil {
		log.Fatal(err)
	}
	// Read response body
	defer r2.Body.Close()

	r2HeadStatus := r2.Status
	r2HeadProto  := r2.Proto
	r2HeadHeader := r2.Header
 	// print out
 	fmt.Printf("Status: %v\nProto: %v\nHeader: %v", r2HeadStatus, r2HeadProto, r2HeadHeader) //<-- here !
 	
 	

}