package domain

import "time"

type (
	Track struct {
		Num       int
		Artist    string
		Album     string
		Title     string
		Duration  time.Duration
		StartTime time.Time
		EndTime   time.Time
	}
)
