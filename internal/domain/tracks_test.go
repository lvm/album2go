package domain

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var (
	artist string = "Cattle Decapitation"
	album  string = "Death Atlas"
	title  string = "Anthropogenic End Transmission"
)

func TestCreateTrack(t *testing.T) {
	track := Track{
		Num:    1,
		Title:  title,
		Artist: artist,
		Album:  album,
	}

	assert.Equal(t, 1, track.Num)
	assert.Equal(t, title, track.Title)
	assert.Equal(t, artist, track.Artist)
	assert.Equal(t, album, track.Album)
}

func TestNewTrack(t *testing.T) {
	num := 1
	duration := time.Duration(1)
	startTime := time.Now()
	endTime := startTime.Add(duration)

	newTrack := NewTrack(num, artist, album, title, duration, startTime, endTime)
	expected := Track{
		Num:       num,
		Artist:    artist,
		Album:     album,
		Title:     title,
		Duration:  duration,
		StartTime: startTime,
		EndTime:   endTime,
	}

	assert.Equal(t, newTrack, expected)
}
