package main

import (
	"time"

	"github.com/pflog/go/log"

	"github.com/pflog/go/container"
)

func main() {
	for {
		log.With(container.NewCtx("4b560c51-358b-4908-b87c-79c0613c457a", "request")).Info("Yeah")

		log.Info(`This is
a multiline message.

Yeah.`)
		log.Info("Hallo")
		log.Infof("Hallo %s! Es ist %d Stunden nach 12", "Heinz", 12)

		time.Sleep(time.Second)
	}
}
