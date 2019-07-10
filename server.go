package main

import (
	"github.com/TarsCloud/TarsGo/tars"
	"impl/demo"
)

func main() {
	demo.AddServant()
	tars.Run()
}
