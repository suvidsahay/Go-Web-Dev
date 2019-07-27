package main

import (
	"fmt"
	"os"
	"text/template"
)
var tl *template.Template

func init(){
	tl=template.Must(template.ParseFiles("index.html"))
}

func main(){
	data:=[]string {
		"Suvid",
		"Sahay",
		"Divus",
		"Yahas",
	}
	err:=tl.ExecuteTemplate(os.Stdout,"index.html",data)
	if err!=nil{
		fmt.Println(err)
	}
}