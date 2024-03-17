package utils

import (
	"encoding/json"

	"github.com/khelechy/argus/models"
)


func IsJsonString(str string) (bool, models.Event, string) {
	var event models.Event
	if json.Unmarshal([]byte(str), &event) == nil {
		return true, event, str
	}

	return false, event, str
}
