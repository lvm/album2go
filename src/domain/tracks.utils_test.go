package domain

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func parseDur(dur string) time.Duration {
	res, _ := time.ParseDuration(dur)
	return res
}

func calcEndTime(startTime time.Time, dur string) time.Time {
	return startTime.Add(parseDur(dur))
}

func TestParseTrack(t *testing.T) {
	artist := "Cattle Decapitation"
	album := "Death Atlas"
	startTime := time.Date(0, 1, 1, 0, 0, 0, 0, time.UTC)

	testData := []struct {
		line     string
		n        int
		title    string
		duration string
	}{
		{
			line:     "1. Test Track 1 - 1:23",
			n:        1,
			title:    "Test Track 1",
			duration: "1m23s",
		},
		{
			line:     "2. Test Track 2 2:34",
			n:        2,
			title:    "Test Track 2",
			duration: "2m34s",
		},
		{
			line:     "3) Test Track 3 3:45",
			n:        3,
			title:    "Test Track 3",
			duration: "3m45s",
		},
		{
			line:     "4) Test Track 4 - 4:56",
			n:        4,
			title:    "Test Track 4",
			duration: "4m56s",
		},
	}

	var expected Track
	for i, td := range testData {
		t.Run(fmt.Sprintf("Running test %d", i), func(t *testing.T) {
			expected = Track{
				Num:       td.n,
				Artist:    artist,
				Album:     album,
				Title:     td.title,
				Duration:  parseDur(td.duration),
				StartTime: startTime,
				EndTime:   calcEndTime(startTime, td.duration),
			}

			track, _ := ParseTrack(artist, album, td.line, startTime)
			assert.Equal(t, track, expected)
		})
	}
}

func testParseDuration(t *testing.T) {

	testData := []struct {
		duration string
		expected time.Duration
	}{
		{
			duration: "1:23",
			expected: parseDur("1m23s"),
		},
		{
			duration: "2:34",
			expected: parseDur("2m34s"),
		},
		{
			duration: "3:45",
			expected: parseDur("3m45s"),
		},
		{
			duration: "4:56",
			expected: parseDur("4m56s"),
		},
	}

	for i, td := range testData {
		t.Run(fmt.Sprintf("Running test %d", i), func(t *testing.T) {
			dur, _ := parseDuration(td.duration)
			assert.Equal(t, dur, td.expected)
		})
	}
}
