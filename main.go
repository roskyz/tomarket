package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"time"
)

func init() {
	if BASE_URL.Getenv("BASE_URL"); BASE_URL == "" {
		panic("BASE_URL must not be empty")
	}
	if _, err := url.Parse(BASE_URL.String()); err != nil {
		panic("BASE_URL must be valid")
	}

	if DEVICE_KEY.Getenv("DEVICE_KEY"); DEVICE_KEY == "" {
		panic("DEVICE_KEY must not be empty")
	}

	ICON_URL.Getenv("ICON_URL")

	SHORT_BREAK_TITLE.Getenv("SHORT_BREAK_TITLE")

	SHORT_BREAK_REMINDER.Getenv("SHORT_BREAK_REMINDER")

	LONG_BREAK_TITLE.Getenv("LONG_BREAK_TITLE")

	LONG_BREAK_REMINDER.Getenv("LONG_BREAK_REMINDER")

	CYCLE_TITLE.Getenv("CYCLE_TITLE")

	CYCLE_REMINDER.Getenv("CYCLE_REMINDER")

	if cycle, _ := strconv.Atoi(os.Getenv("CYCLE")); cycle > 0 {
		CYCLE = cycle
	}
}

func main() {
	isCyclePhase := true
	cycles := 0

	for {
		cycles++

		var err error
		if isCyclePhase {
			err = performActivity(CYCLE_TITLE, CYCLE_REMINDER, CYCLE)
		} else if cycles%LONG_BREAK_INTERVAL == 0 {
			err = performActivity(LONG_BREAK_TITLE, LONG_BREAK_REMINDER, LONG_BREAK_DURATION)
		} else {
			err = performActivity(SHORT_BREAK_TITLE, SHORT_BREAK_REMINDER, SHORT_BREAK_DURATION)
		}

		if err != nil {
			fmt.Printf("barking error: %v\n", err)
		}

		isCyclePhase = !isCyclePhase
	}
}

func performActivity(title, reminder trimmedStr, durationMinutes int) error {
	err := sendBarking(title, reminder)
	if err != nil {
		return err
	}
	sleepDuration := time.Duration(durationMinutes) * time.Minute
	time.Sleep(sleepDuration)
	return nil
}

func sendBarking(title, body trimmedStr) error {
	r, err := http.Get(fmt.Sprintf("%s/%s/%s/%s?icon=%s", BASE_URL, DEVICE_KEY, title, body, ICON_URL))
	if err != nil {
		return err
	}
	defer r.Body.Close()
	if r.StatusCode != http.StatusOK {
		resp, err := io.ReadAll(r.Body)
		if err != nil {
			return err
		}
		return fmt.Errorf("http status code not 200, response: %s", string(resp))
	}
	return nil
}
