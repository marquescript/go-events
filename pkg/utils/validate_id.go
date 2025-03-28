package utils

import (
	"strconv"
)

func ParseID(id string) (int64, error) {
	eventID, err := strconv.ParseInt(id, 10, 64)
	return eventID, err
}
