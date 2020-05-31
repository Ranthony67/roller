package main

import (
	"time"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/drivers/gpio"
	"gobot.io/x/gobot/platforms/raspi"
)

func main() {
	adaptor := raspi.NewAdaptor()
	led7 := gpio.NewLedDriver(adaptor, "7")

	blinkLed := func() {
		gobot.Every(1*time.Second, func() {
			led7.Toggle()
		})
	}

	robot := gobot.NewRobot("Blink Test", []gobot.Connection{adaptor}, []gobot.Device{led7}, blinkLed)

	robot.Start()
}
