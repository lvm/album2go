package domain

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func parseDuration(durationStr string) (time.Duration, error) {
	parts := strings.Split(durationStr, ":")
	mult := []time.Duration{time.Second, time.Minute, time.Hour}
	if len(parts) > 3 {
		return 0, fmt.Errorf("invalid duration format")
	}
	var duration time.Duration
	for i, part := range parts {
		val, err := strconv.Atoi(part)
		if err != nil {
			return 0, err
		}
		duration += time.Duration(val) * mult[len(parts)-i-1]
	}
	return duration, nil
}

func ParseTrack(artist, album, trackInfo string, startTime time.Time) (Track, error) {
	re := regexp.MustCompile(`(?P<n>\d+[\.\)])\s+(?P<track>.+?)\s*(-\s+)?(?P<time>(\d{1,2}:)?\d{1,2}:\d{2})`)

	match := re.FindStringSubmatch(trackInfo)
	if match != nil {
		params := make(map[string]string)
		for i, name := range re.SubexpNames() {
			if i != 0 && name != "" {
				params[name] = match[i]
			}
		}

		duration, err := parseDuration(params["time"])
		if err != nil {
			return Track{}, err
		}

		num, _ := strconv.Atoi(strings.Trim(params["n"], ".)"))
		endTime := startTime.Add(duration)

		return NewTrack(num, artist, album, strings.TrimSpace(params["track"]), duration, startTime, endTime), nil
	}

	return Track{}, fmt.Errorf("no matching pattern found")
}
