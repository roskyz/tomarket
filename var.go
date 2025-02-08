package main

import (
	"os"
	"strings"
)

type trimmedStr string

func (s trimmedStr) String() string { return string(s) }

func (s *trimmedStr) Getenv(key string) {
	if ts := strings.TrimSpace(os.Getenv(key)); ts != "" {
		*s = trimmedStr(ts)
	}
}

var (
	CYCLE                     = 35
	CYCLE_TITLE    trimmedStr = "Mindfulness"
	CYCLE_REMINDER trimmedStr = "Get Back to Work!"
)

var (
	BASE_URL   trimmedStr
	DEVICE_KEY trimmedStr
	ICON_URL   trimmedStr
)

var (
	SHORT_BREAK_DURATION            = 5
	SHORT_BREAK_TITLE    trimmedStr = "Relax"
	SHORT_BREAK_REMINDER trimmedStr = "Break Time! Take a Breath!"

	LONG_BREAK_DURATION            = 10
	LONG_BREAK_TITLE    trimmedStr = "Burnout"
	LONG_BREAK_REMINDER trimmedStr = "Session Complete! Reset!"
	LONG_BREAK_INTERVAL            = 4
)
