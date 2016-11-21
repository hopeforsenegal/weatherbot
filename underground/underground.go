package underground

import (
	"fmt"
	"github.com/danryan/hal"
	"github.com/greyfocus/go-wunderground-api"
)

// Key
const WeatherAPIKey = "ac29d0ab5aa732b4"
const CurrentLocation = "TX/Austin"

// Returns the conditions in a certain type of weather
var Underground = &hal.Handler{
	Method:  hal.HEAR,
	Pattern: `underground`,
	Run: func(res *hal.Response) error {

		client := wunderground_api.JsonClient{WeatherAPIKey}
		request := wunderground_api.Request{Features: []string{"conditions"}, Location: CurrentLocation}

		resp, err := client.Execute(&request)
		if err != nil {
			fmt.Println(err)
			return err
		}

		if resp.CurrentConditions == nil {
			fmt.Println("The current conditions were not returned. Is the API key correct?")
			return err
		}

		returnString := fmt.Sprintf("%s       Temperature: %6.1f C (feels like %2.1f C)\n", CurrentLocation, resp.CurrentConditions.TempC, resp.CurrentConditions.FeelsLikeC)
		return res.Send(returnString)
	},
}
