package main

import (
	"github.com/kmdkuk/my-blog-go/server"
)

func main() {
	r := server.SetUpRouting()
	r.Run()
}
