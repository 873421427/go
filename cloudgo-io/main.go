package main

import(
	"github.com/873421427/hello/cloudgo-io/service"	
)

func main(){
	server :=service.NewServer()
	server.Run(":8080")
}