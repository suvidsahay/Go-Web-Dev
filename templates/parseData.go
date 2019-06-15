package main

import (
	"log"
	"os"
	"text/template"
)

var tl *template.Template

func init(){
	tl =template.Must(template.ParseFiles("index.html"))
}

func main(){
	err:=tl.ExecuteTemplate(os.Stdout,"index.html",`Suvid`)
	if(err!=nil){
		log.Fatalln(err)
	}
}
