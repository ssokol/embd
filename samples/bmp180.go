// +build ignore

package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/ssokol/embd"
	"github.com/ssokol/embd/sensor/bmp180"

	_ "github.com/ssokol/embd/host/all"
)

func main() {
	flag.Parse()

	if err := embd.InitI2C(); err != nil {
		panic(err)
	}
	defer embd.CloseI2C()

	bus := embd.NewI2CBus(1)

	baro := bmp180.New(bus)
	defer baro.Close()

	for {
		temp, err := baro.Temperature()
		if err != nil {
			panic(err)
		}
		fmt.Printf("Temp is %v\n", temp)
		pressure, err := baro.Pressure()
		if err != nil {
			panic(err)
		}
		fmt.Printf("Pressure is %v\n", pressure)
		altitude, err := baro.Altitude()
		if err != nil {
			panic(err)
		}
		fmt.Printf("Altitude is %v\n", altitude)

		time.Sleep(500 * time.Millisecond)
	}
}
