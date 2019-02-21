package main

import (
	"flag"
	"log"
	"webchart/route"
)
var addr = flag.String("addr", "localhost:8080", "http service address")
func main() {
	 flag.Parse()
	 log.SetFlags(0)
	 log.Fatal(route.StartHttp(*addr))
}
