package main

import (
	"github.com/danryan/hal"
	_ "github.com/danryan/hal/adapter/shell"
	"github.com/danryan/hal/handler"
	_ "github.com/danryan/hal/store/memory"
	"os"
	"time"
	"fmt"
)

type MyError struct {
	When time.Time
	What string
}

func (e MyError) Error() string {
	return fmt.Sprintf("%v: %v", e.When, e.What)
}

// HAL is just another Go package, which means you are free to organize things
// however you deem best.

// You can define your handlers in the same file...
var weatherHandler = hal.Hear(`weather`, func(res *hal.Response) error {
	return res.Send("Its cold ooutside")
})
var quitFlipHandler = &hal.Handler{
		Method:  hal.HEAR,
		Pattern: `quit`,
		Run: func(res *hal.Response) error {
			return MyError{
				time.Now(),
				"The user quit",
			}
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