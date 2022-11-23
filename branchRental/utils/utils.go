package utils

import (
	"fmt"
	"time"
)

func ConvertToEpoch(dateString string) (int64, error) {
	epochTime, err := time.Parse(time.RFC3339, dateString)
	if err != nil {
		return -1, fmt.Errorf("Err in datetime string conversion")
	}
	return epochTime.Unix(), nil
}
