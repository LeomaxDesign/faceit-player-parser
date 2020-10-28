package main

import (
	"faceit-parser/web"
	"fmt"
)

func main() {

	fmt.Println("LOL")
	server := web.New(":8080")
	server.Start()

}
