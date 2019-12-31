package main

import(
	"fmt"
	"encoding/xml"

)

type Foo struct {
	Bar   string   `xml:"id,attr"`
	Baz   string   `xml:"parent>child"`
}

func main(){
	f := Foo{"Joe Junior", "Hello Shabado"}
    b, _ := xml.Marshal(f)
    fmt.Println(string(b))
    xml.Unmarshal(b, &f)	
}