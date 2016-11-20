package main

import (
	"github.com/danryan/hal"
	_ "github.com/danryan/hal/adapter/shell"
	"github.com/danryan/hal/handler"
	_ "github.com/danryan/hal/store/memory"
	"os"
)

// Key
const kWeatherAPIKey = "ac29d0ab5aa732b4"

// Handlers
var weatherHandler = hal.Hear(`weather`, func(res *hal.Response) error {
	return res.Send("Its cold outside")
})
var quitFlipHandler = &hal.Handler{
		Method:  hal.HEAR,
		Pattern: `quit`,
		Run: func(res *hal.Response) error {
			return res.Robot.Stop()
		},
	}

func run() int {
	robot, err := hal.NewRobot()
	if err != nil {
		hal.Logger.Error(err)
		return 1
	}

	robot.Handle(
		weatherHandler,
		handler.TableFlip,
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