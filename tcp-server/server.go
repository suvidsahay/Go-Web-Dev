package main

import (
	"bufio"
	"fmt"
	"net"
)

func main(){
	li,err:=net.Listen("tcp",":8080")
	if err!=nil{
		panic(err)
	}
	for {
		conn,err:=li.Accept()
		if err!=nil{
			fmt.Println(err)
		}
		go handle(conn)
	}
}

func handle(conn net.Conn){
	request(conn)
	//respond(req)
}
 func request(conn net.Conn){
 	scanner:=bufio.NewScanner(conn)
 	defer conn.Close()
 	for scanner.Scan(){
		ln:= scanner.Text()
		fmt.Println(ln)
	}
 }
