package tripit

// This file was generated by a tool. Do not edit.

import (
	"encoding/json"
)

// A specialization of Vector for WeatherObject objects
type WeatherObjectVector []WeatherObject

func (p *WeatherObjectVector) UnmarshalJSON(b []byte) error {
	var arr *[]WeatherObject
	arr = (*[]WeatherObject)(p)
	*arr = nil
	err := json.Unmarshal(b, arr)
	if err != nil {
		*arr = make([]WeatherObject, 1)
		err := json.Unmarshal(b, &(*arr)[0])
		if err != nil {
			if err2, ok := err.(*json.UnmarshalTypeError); ok && err2.Value == "null" {
				*arr = (*arr)[0:0]
			} else {
				return err
			}
		}
		
	}
	return nil
}
