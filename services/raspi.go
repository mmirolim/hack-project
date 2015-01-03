package services

import (
	"github.com/hybridgroup/gobot"
	"github.com/hybridgroup/gobot/platforms/gpio"
	"github.com/hybridgroup/gobot/platforms/raspi"

	"github.com/mmirolim/hack-project/conf"
	ds "github.com/mmirolim/hack-project/datastore"
)

func StartStatusBot(rs conf.Raspi, status <-chan ds.Status) {
	gbot := gobot.NewGobot()

	r := raspi.NewRaspiAdaptor(rs.AdaptorName)
	led := gpio.NewLedDriver(r, rs.Led.Name, rs.Led.Pin)

	work := func() {
		for v := range status {
			// @TODO should toggle different leds red and green
			if v == 1 {
				led.Toggle()
			}
		}
	}

	robot := gobot.NewRobot(rs.BotName,
		[]gobot.Connection{r},
		[]gobot.Device{led},
		work,
	)

	gbot.AddRobot(robot)

	gbot.Start()
}
