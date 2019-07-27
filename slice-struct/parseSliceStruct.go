package main

import (
	"fmt"
	"html/template"
	"os"
)

var tl *template.Template
type person struct{
	Name string
	Age int
}
func init(){
	tl=template.Must(template.ParseFiles("index.html"))
}

func main(){
	suvid:=person{
			"Suvid",
			20,
		}
	Divus:=person{
			"Divus",
			20,
		}
	Rahul:=person{
			"Rahul",
			19,
		}
	data:=[]person{suvid,Divus,Rahul}
	err:=tl.ExecuteTemplate(os.Stdout,"index.html",data)
	if err!=nil{
		fmt.Println(err)
	}
}
