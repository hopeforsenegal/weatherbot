package main

import (
	"fmt"
	"github.com/danryan/hal"
	_ "github.com/danryan/hal/adapter/shell"
	"github.com/danryan/hal/handler"
	_ "github.com/danryan/hal/store/memory"
	"github.com/hopeforsenegal/weatherbot/underground"
	"os"
)

// Handlers
var weatherHandler = hal.Hear(`weather`, func(res *hal.Response) error {
	return res.Send("Its cold outside")
})

var gloatHandler = hal.Hear(`gloat`, func(res *hal.Response) error {
	return res.Send("Kamau is a boss. Or at least he tries to be")
})

var echoHandler = hal.Hear(`echo (.+)`, func(res *hal.Response) error {
	return res.Reply(res.Match[1])
})

func run() int {
	robot, err := hal.NewRobot()
	if err != nil {
		hal.Logger.Error(err)
		return 1
	}

	robot.Handle(
		echoHandler,
		gloatHandler,
		handler.TableFlip,
		weatherHandler,
		underground.Underground,
	)

	defer fmt.Println("Terminating")

	if err := robot.Run(); err != nil {
		hal.Logger.Error(err)
		return 1
	}
	return 0
}

func main() {
	os.Exit(run())
}
