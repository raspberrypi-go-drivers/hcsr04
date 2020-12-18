# HC-SR04

[![PkgGoDev](https://pkg.go.dev/badge/github.com/bbayszczak/raspberrypi-go-drivers/hcsr04)](https://pkg.go.dev/github.com/bbayszczak/raspberrypi-go-drivers/hcsr04)

hcsr04 is a driver allowing to usea HC-SR04 ultrasound distance sensor

## Documentation

For full documentation, please visit [![PkgGoDev](https://pkg.go.dev/badge/github.com/bbayszczak/raspberrypi-go-drivers/hcsr04)](https://pkg.go.dev/github.com/bbayszczak/raspberrypi-go-drivers/hcsr04)

## Quick start

```go
package main

import (
	"fmt"
	"os"
	"time"

	"github.com/raspberrypi-go-drivers/hcsr04"
	"github.com/stianeikeland/go-rpio/v4"
)

func main() {
	err := rpio.Open()
	if err != nil {
		os.Exit(1)
	}
	defer rpio.Close()
	h := hcsr04.NewHCSR04(17, 27)
	h.StartDistanceMonitor()
	for {
		fmt.Println(h.GetDistance())
		time.Sleep(hcsr04.MonitorUpdate)
	}
}
```
