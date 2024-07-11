package main

import (
	"fmt"

	"github.com/ysnbhb/ascii-art-web/serve"
	// package creat by team
)

func main() {
	serve := serve.NewPort(":8081")
	fmt.Println(serve.Start())
}
