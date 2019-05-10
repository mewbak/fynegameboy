package main

import (
	"log"
	"os"

	"github.com/andydotxyz/gameboy.live/gb"
	"github.com/andydotxyz/fynegameboy/driver"
)

func romPath() string {
	args := os.Args
	if len(args) < 2 {
		return ""
	}

	return args[1]
}

func main() {
	rom := romPath()
	if rom == "" {
		log.Println("ROM parameter missing")
		return
	}

	Driver := new(driver.LCD)
	core := new(gb.Core)
	core.FPS = 60
	core.Clock = 4194304
	core.DisplayDriver = Driver
	core.Controller = Driver
	core.DrawSignal = make(chan bool)
	core.ToggleSound = true
	core.SpeedMultiple = 0
	core.Init(rom)
	go core.Run()
	core.DisplayDriver.Run(core.DrawSignal)
}
