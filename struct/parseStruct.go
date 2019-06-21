package main

import (
	"fmt"
	"html/template"
	"os"
)

var tl *template.Template

func init(){
	tl=template.Must(template.ParseFiles("index.html"))
}
type person struct{
	Name string
	Title string
}
func main(){
	data:=person{
		"Suvid",
		"Sahay",
	}
	err:=tl.ExecuteTemplate(os.Stdout,"index.html",data)
	if err!=nil{
		fmt.Println(err)
	}
}
