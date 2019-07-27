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

func main(){
	data:=make(map[string]string)
	data["Suvid"]="My name is Suvid"
	data["Sahay"]="My name is Sahay"
	data["Divus"]="My name is Divus"
	data["Yahas"]="My name is Yahas"
	err:=tl.ExecuteTemplate(os.Stdout,"index.html",data)
	if err!=nil{
		fmt.Println(err)
	}
}
