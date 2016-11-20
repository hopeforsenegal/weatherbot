package wunderground_api

import (
	"encoding/json"
	"fmt"
	"io"
	"strconv"
	"time"
)

// The Json attribute containing the current conditions
const CURRENT_CONDITIONS_ATTR = "current_observation"

type parseError struct {
	field string
	err   error
}

func (e *parseError) String() string {
	if e.err != nil {
		return fmt.Sprintf("error parsing field %s: %v", e.field, e.err)
	}
	return fmt.Sprintf("error parsing field %s", e.field)
}

// Parses the JSON response from the body parameter and returns a model struct.
// The type information of the data returned by the Wunderground API is not
// very reliable. For some locaitons, some attributes are returned as string,
// while for different locaitons they are returned as float.
//
// The library attemps to normalize the types in order to produce consistent
// results, but this comes at the cost of not being able to use the field
// mapping feautres of the encoding JSON library.
func parseWeatherResponse(body io.ReadCloser) (weatherResponse *Response, err error) {
	decoder := json.NewDecoder(body)
	var responseMap interface{}
	localErr := decoder.Decode(&responseMap)
	if localErr != nil {
		return nil, localErr
	}

	weatherResponse = &Response{}
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("%v", r)
		}
	}()

	weatherResponse.CurrentConditions = parseConditions(responseMap.(map[string]interface{}))

	return
}

func parseConditions(m map[string]interface{}) *Conditions {
	c := &Conditions{}

	if m[CURRENT_CONDITIONS_ATTR] == nil {
		return nil
	}

	conditionsMap := m[CURRENT_CONDITIONS_ATTR].(map[string]interface{})

	c.TempC = decodeToFloat(conditionsMap["temp_c"], "temp_c")
	c.TempF = decodeToFloat(conditionsMap["temp_f"], "temp_f")
	c.FeelsLikeC = decodeToFloat(conditionsMap["feelslike_c"], "feelslike_c")
	c.FeelsLikeF = decodeToFloat(conditionsMap["feelslike_f"], "feelslike_f")
	c.PressureMb = decodeToFloat(conditionsMap["pressure_mb"], "pressure_mb")
	c.RelativeHumidity = decodeToString(conditionsMap["relative_humidity"], "relative_humidity")
	c.VisibilityKm = decodeToFloat(conditionsMap["visibility_km"], "visibility_km")
	c.Weather = decodeToString(conditionsMap["weather"], "weather")
	c.Wind = decodeToString(conditionsMap["wind_string"], "wind_string")
	c.WindKph = decodeToFloat(conditionsMap["wind_kph"], "wind_kph")
	c.WindGustKph = decodeToFloat(conditionsMap["wind_gust_kph"], "wind_gust_kph")
	c.WindDegrees = decodeToFloat(conditionsMap["wind_degrees"], "wind_degrees")
	c.WindDir = decodeToString(conditionsMap["wind_dir"], "wind_dir")
	c.ObservationTime = time.Unix(decodeToInt(conditionsMap["observation_epoch"], "observation_epoch"), 0)

	return c
}

func decodeToFloat(v interface{}, field string) float32 {
	switch vv := v.(type) {
	case string:
		if vv == "N/A" {
			// Workaround for the case when the value is not provided by the API
			return -1
		}

		f64, e := strconv.ParseFloat(vv, 32)
		if e != nil {
			panic(&parseError{field, e})
		}

		return float32(f64)
	case float32:
		return vv
	case float64:
		return float32(vv)
	default:
		panic(&parseError{field, nil})
	}
}

func decodeToString(v interface{}, field string) string {
	switch vv := v.(type) {
	case string:
		return vv
	case float32:
		return strconv.FormatFloat(float64(vv), 'f', 2, 32)
	case float64:
		return strconv.FormatFloat(vv, 'f', 2, 64)
	case int:
		return strconv.Itoa(vv)
	default:
		panic(&parseError{field, nil})
	}
}

func decodeToInt(v interface{}, field string) int64 {
	switch vv := v.(type) {
	case string:
		t, e := strconv.ParseInt(vv, 10, 32)
		if e != nil {
			panic(&parseError{field, e})
		}

		return int64(t)
	case int64:
		return vv
	case int32:
		return int64(vv)
	default:
		panic(&parseError{field, nil})
	}
}
