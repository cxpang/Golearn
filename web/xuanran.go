package main

import (
	"html/template"
	"os"
)
type Person struct {
	UserName string
	Email string
}
func main(){
	t:=template.New("fieldname example")
	t,_=t.Parse("hello,{{.UserName}}!{{.Email}}\n")
	p:=Person{UserName:"旁侧后"}
	t.Execute(os.Stdout,p)
}
