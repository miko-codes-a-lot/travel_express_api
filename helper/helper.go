package helper

import "time"

func Now() string {
	now := time.Now().UTC()
	iso8601NanoString := now.Format(time.RFC3339Nano)
	return iso8601NanoString
}
