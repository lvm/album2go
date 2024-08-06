package domain

import "time"

type Track struct {
	Num       int
	Artist    string
	Album     string
	Title     string
	Duration  time.Duration
	StartTime time.Time
	EndTime   time.Time
}

func NewTrack(num int, artist, album, title string, duration time.Duration, startTime, endTime time.Time) Track {
	return Track{
		Num:       num,
		Artist:    artist,
		Album:     album,
		Title:     title,
		Duration:  duration,
		StartTime: startTime,
		EndTime:   endTime,
	}
}
