package main

import (
	"fmt"
	"encoding/xml"
	"os"
)

type Servers struct {
	XMLName xml.Name `xml:"servers"`
	Version string `xml:"version,attr"`
	Svs []server  `xml:"server"`
}
type server struct {
	ServerName string `xml:"ServerName"`
	ServerIP string `xml:"ServerIP"`
}
func main()  {
	v:=&Servers{Version:"1"}
	v.Svs=append(v.Svs,server{"shanghai_VPN","127.0.0.1"})
	v.Svs=append(v.Svs,server{"beijing_VPN","127.0.0.1"})
	output,err:=xml.MarshalIndent(v," ","   ")
	if err!=nil{
		fmt.Println("error",err)
	}
	os.Stdout.Write([]byte(xml.Header))
	os.Stdout.Write(output)
}
