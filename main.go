package main

import (
	"flag"
	transfer_controller "practice/transfer_token"
	view_controller "practice/view_function"
)

func main() {
	var v = flag.Int("view", 1, "")
	var t = flag.Int("tx", 1, "")
	flag.Parse()
	for i := 0; i < *v; i++ {
		view_controller.Print()
	}
	for i := 0; i < *t; i++ {
		transfer_controller.Print()
	}
}
