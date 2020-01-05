package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"os"
	"net/url"
)	
	func main(){
	
    form := url.Values{} 
    form.Add("foo", "bar") 
    r3, err := http.PostForm("https://www.google.com/robots.txt", form) 

 	r3html, err := ioutil.ReadAll(r3.Body)	
	if err != nil {
 		fmt.Println(err)
 	}
 	fmt.Println(os.Stdout, string(r3html)) //<-- here !
} 	