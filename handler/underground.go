package handler

import (
	"github.com/danryan/hal"
)

// TableFlip is an example of a Handler
var Underground = &hal.Handler{
	Method:  hal.HEAR,
	Pattern: `underground`,
	Run: func(res *hal.Response) error {

		client := wunderground_api.JsonClient{*api_key}
		request := wunderground_api.Request{Features: []string{"conditions"}, Location: *location}

		resp, err := client.Execute(&request)
		if err != nil {
			fmt.Println(err)
			return
		}

		if resp.CurrentConditions == nil {
			fmt.Println("The current conditions were not returned. Is the API key correct?")
			return
		}
	
		return res.Send(`(╯°□°）╯︵ ┻━┻`)
	},
}