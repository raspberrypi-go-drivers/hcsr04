package hcsr04_test

import (
	"fmt"
	"os"
	"time"

	"github.com/raspberrypi-go-drivers/hcsr04"
	log "github.com/sirupsen/logrus"
	"github.com/stianeikeland/go-rpio/v4"
)

func Example() {
	err := rpio.Open()
	if err != nil {
		os.Exit(1)
	}
	defer rpio.Close()
	h := hcsr04.NewHCSR04(17, 27)
	fmt.Println(h.MeasureDistance())
}

func Example_monitor() {
	err := rpio.Open()
	if err != nil {
		os.Exit(1)
	}
	defer rpio.Close()
	h := hcsr04.NewHCSR04(17, 27)
	if err := h.StartDistanceMonitor(); err != nil {
		log.WithField("error", err).Error("impossible to start distance monitor")
		os.Exit(1)
	} else {
		defer h.StopDistanceMonitor()
	}
	for {
		fmt.Println(h.GetDistance())
		time.Sleep(hcsr04.MonitorUpdate)
	}
}

func ExampleNewHCSR04() {
	err := rpio.Open()
	if err != nil {
		os.Exit(1)
	}
	defer rpio.Close()
	h := hcsr04.NewHCSR04(17, 27)
	fmt.Println(h.MeasureDistance())
}

func ExampleHCSR04_MeasureDistance() {
	err := rpio.Open()
	if err != nil {
		os.Exit(1)
	}
	defer rpio.Close()
	h := hcsr04.NewHCSR04(17, 27)
	fmt.Println(h.MeasureDistance())
}

func ExampleHCSR04_GetDistance() {
	err := rpio.Open()
	if err != nil {
		os.Exit(1)
	}
	defer rpio.Close()
	h := hcsr04.NewHCSR04(17, 27)
	if err := h.StartDistanceMonitor(); err != nil {
		log.WithField("error", err).Error("impossible to start distance monitor")
		os.Exit(1)
	} else {
		defer h.StopDistanceMonitor()
	}
	for {
		fmt.Println(h.GetDistance())
		time.Sleep(hcsr04.MonitorUpdate)
	}
}

func ExampleHCSR04_StopDistanceMonitor() {
	err := rpio.Open()
	if err != nil {
		os.Exit(1)
	}
	defer rpio.Close()
	h := hcsr04.NewHCSR04(17, 27)
	if err := h.StartDistanceMonitor(); err != nil {
		log.WithField("error", err).Error("impossible to start distance monitor")
		os.Exit(1)
	} else {
		defer h.StopDistanceMonitor()
	}
	for {
		fmt.Println(h.GetDistance())
		time.Sleep(hcsr04.MonitorUpdate)
	}
}
