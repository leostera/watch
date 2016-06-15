package main

import (
	"fmt"
	"os"
	"time"
)

func dieIf(err error, status int, message string) {
	if err != nil {
		die(status, message)
	}
}

func die(status int, message string) {
	if len(message) > 0 {
		fmt.Println(message)
	}
	os.Exit(status)
}

func intervalToTime(i float64) time.Duration {
	return time.Duration(i) * time.Millisecond
}

func loop(d time.Duration, fn func()) {
	fn()
	for range time.Tick(d) {
		fn()
	}
}
