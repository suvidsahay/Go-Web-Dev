package main

import (
	"html/template"
	"log"
	"os"
	"strings"
)

var tl *template.Template
var fm= template.FuncMap{
	"uc":strings.ToUpper,
	"f3":firstThree,
}
func firstThree(s string) string{
	s=s[:3]
	return s
}

func init(){
	tl=template.Must(template.New("").Funcs(fm).ParseFiles("index.html"))
}

func main(){
	data:=[]string{
		"Suvid",
		"Sahay",
		"Divus",
		"Yahas",
	}
	err:=tl.ExecuteTemplate(os.Stdout,"index.html",data)
	if err!=nil{
		log.Fatalln(err)
	}
}
