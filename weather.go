package main

import (
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

var echoHandler = hal.Hear(`echo (.+)`, func(res *hal.Response) error {
	return res.Reply(res.Match[1])
})

var quitHandler = &hal.Handler{
	Method:  hal.HEAR,
	Pattern: `quit`,
	Run: func(res *hal.Response) error {
		return res.Send("Told to quit")
	},
}

func run() int {
	robot, err := hal.NewRobot()
	if err != nil {
		hal.Logger.Error(err)
		return 1
	}

	robot.Handle(
		quitHandler,
		echoHandler,
		handler.TableFlip,
		weatherHandler,
		underground.Underground,
	)

	if err := robot.Run(); err != nil {
		hal.Logger.Error(err)
		return 1
	}
	return 0
}

func main() {
	os.Exit(run())
}
