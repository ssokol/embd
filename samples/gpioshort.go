// +build ignore

package main

import (
	"flag"

	"github.com/ssokol/embd"

	_ "github.com/ssokol/embd/host/all"
)

func main() {
	flag.Parse()

	embd.InitGPIO()
	defer embd.CloseGPIO()

	embd.SetDirection(10, embd.Out)
	embd.DigitalWrite(10, embd.High)
}
