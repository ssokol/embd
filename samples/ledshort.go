// +build ignore

// Short LED example, works OOTB on a BBB.

package main

import (
	"flag"
	"time"

	"github.com/ssokol/embd"

	_ "github.com/ssokol/embd/host/bbb"
)

func main() {
	flag.Parse()

	embd.InitLED()
	defer embd.CloseLED()

	embd.LEDOn(3)
	time.Sleep(1 * time.Second)
	embd.LEDOff(3)
}
