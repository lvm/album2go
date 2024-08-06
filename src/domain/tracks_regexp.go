package domain

import "regexp"

var trackPatterns = []*regexp.Regexp{
	regexp.MustCompile(`(?P<time>(\d+:)?\d{2}:\d{2}) - (?P<track>.+)`),
	regexp.MustCompile(`(?P<time>(\d+:)?\d{2}:\d{2}) (?P<n>\d+\.) (?P<track>.+)`),
	regexp.MustCompile(`(?P<time>(\d+:)?\d{2}:\d{2}) (?P<n>\d+\)) (?P<track>.+)`),
	regexp.MustCompile(`(?P<n>\d+\.) (?P<time>(\d{1,2}:)?\d{2}:\d{2}) - (?P<track>.+)`),
	regexp.MustCompile(`(?P<n>\d+\.)?( )?(?P<track>.+) (?P<time>(\d{1,2}:)?\d{1,2}:\d{2})`),
	regexp.MustCompile(`(?P<n>\d+\)) (?P<track>.+) (?P<time>(\d{1,2}:)?\d{1,2}:\d{2})`),
	regexp.MustCompile(`(?P<time>(\d{1,2}:)?\d{1,2}:\d{2}) - (?P<track>.+)`),
	regexp.MustCompile(`(?P<time>(\d{1,2}:)?\d{1,2}:\d{2}) (?P<n>\d+(\.|\))) (?P<track>.+)`),
	regexp.MustCompile(`(?P<time>(\d{1,2}:)?\d{1,2}:\d{2}) (?P<track>.+)`),
	regexp.MustCompile(`(?P<track>.+) (?P<time>(\d{1,2}:)?\d{1,2}:\d{2})`),
	regexp.MustCompile(`(?P<n>\d+\.) (?P<track>.+) - (?P<time>(\d{1,2}:)?\d{1,2}:\d{2})`),
	regexp.MustCompile(`(?P<track>.+) (?P<time>(\d{1,2}:)?\d{1,2}:\d{2})`),
	regexp.MustCompile(`\[(?P<time>(\d{1,2}:)?\d{1,2}:\d{2})\] (?P<track>.+)`),
}
