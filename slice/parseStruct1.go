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
	data1:=[]string{
		"Suvid",
		"Sahay",
		"Divus",
		"Yahas",
	}
	err:=tl1.ExecuteTemplate(os.Stdout,"index1.html",data1)
	if err!=nil{
		fmt.Println(err)
	}
}
