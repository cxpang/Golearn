package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
)
func sayhelloName(w http.ResponseWriter,r *http.Request){
	fmt.Println(r.Form)
	fmt.Println("path",r.URL.Path)
	fmt.Println("scheme",r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k,v:=range r.Form{
		fmt.Println("key:",k)
		fmt.Println("val:",strings.Join(v,""))
	}
	t,_:=template.ParseFiles("index.html")
	t.Execute(w,nil)
	
}
func login(w http.ResponseWriter,r *http.Request)  {
	fmt.Println("Method:",r.Method)
	if r.Method=="GET" {
		t,_:=template.ParseFiles("login.html")
		t.Execute(w,nil)
	}else {
		r.ParseForm()
		fmt.Println("username:",r.Form["username"])
		fmt.Println("password:",r.Form["userpwd"])
	}
	
}
func main(){
	http.HandleFunc("/",sayhelloName)
	http.HandleFunc("/login",login)
	err:=http.ListenAndServe(":9090",nil)
	if err!=nil{
		log.Fatal("ListenAndServer:",err)
	}
}
