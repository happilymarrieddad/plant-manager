package utils

import "time"

func TimeNow() *time.Time {
	t := time.Now()
	return &t
}
