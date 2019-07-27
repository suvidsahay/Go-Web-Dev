package main

import (
	"fmt"
	"html/template"
	"os"
)

var tl1 *template.Template

func init(){
	tl1=template.Must(template.ParseFiles("index1.html"))
}

func main(){
	data:=map[string]string{
		"Suvid":"My name is Suvid",
		"Sahay":"My name is Sahay",
		"Divus":"My name is Divus",
		"Yahas":"My name is Yahas",
	}
	err:=tl1.ExecuteTemplate(os.Stdout,"index1.html",data)
	if err!=nil{
		fmt.Println(err)
	}
}
