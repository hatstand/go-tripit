package tripit

// This file was generated by a tool. Do not edit.

import (
	"encoding/json"
)

// A specialization of Vector for ProfileEmailAddress objects
type ProfileEmailAddressVector []ProfileEmailAddress

func (p *ProfileEmailAddressVector) UnmarshalJSON(b []byte) error {
	var arr *[]ProfileEmailAddress
	arr = (*[]ProfileEmailAddress)(p)
	*arr = nil
	err := json.Unmarshal(b, arr)
	if err != nil {
		*arr = make([]ProfileEmailAddress, 1)
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
