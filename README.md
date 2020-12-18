# HC-SR04

[![Go Reference](https://pkg.go.dev/badge/github.com/raspberrypi-go-drivers/hcsr04.svg)](https://pkg.go.dev/github.com/raspberrypi-go-drivers/hcsr04)
![golangci-lint](https://github.com/raspberrypi-go-drivers/hcsr04/workflows/golangci-lint/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/raspberrypi-go-drivers/hcsr04)](https://goreportcard.com/report/github.com/raspberrypi-go-drivers/hcsr04)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

hcsr04 is a driver allowing to usea HC-SR04 ultrasound distance sensor

## Documentation

For full documentation, please visit [![Go Reference](https://pkg.go.dev/badge/github.com/raspberrypi-go-drivers/hcsr04.svg)](https://pkg.go.dev/github.com/raspberrypi-go-drivers/hcsr04)

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

## Raspberry Pi compatibility

This driver has has only been tested on an Raspberry Pi Zero WH using integrated bluetooth but may work well on other Raspberry Pi having integrated Bluetooth

## License

MIT License

---

Special thanks to @stianeikeland

This driver is based on his work in [stianeikeland/go-rpio](https://github.com/stianeikeland/go-rpio/)
