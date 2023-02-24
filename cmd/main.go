package main

import (
	"github.com/xiaofeng-tech/prometheus-demo/src/context"
)

func main() {
	println("Start...")
	context.NewAppContexxt().Server()
	println("End ....")
}