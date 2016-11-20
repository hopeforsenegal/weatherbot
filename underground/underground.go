package underground

import (
	"github.com/greyfocus/go-wunderground-api"
	"github.com/danryan/hal"
	"fmt"
)

// Key
const WeatherAPIKey = "ac29d0ab5aa732b4"
const CurrentLocation = "France/Paris"

// TableFlip is an example of a Handler
var Underground = &hal.Handler{
	Method:  hal.HEAR,
	Pattern: `underground`,
	Run: func(res *hal.Response) error {

		client := wunderground_api.JsonClient{WeatherAPIKey}
		request := wunderground_api.Request{Features: []string{"conditions"}, Location: CurrentLocation}

		resp, err := client.Execute(&request)
		if err != nil {
			fmt.Println(err)
			return
		}

		if resp.CurrentConditions == nil {
			fmt.Println("The current conditions were not returned. Is the API key correct?")
			return
		}

		return res.Send("       Temperature: %6.1f C (feels like %2.1f C)\n", resp.CurrentConditions.TempC, resp.CurrentConditions.FeelsLikeC)
	},
}