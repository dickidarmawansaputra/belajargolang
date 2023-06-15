package test

import (
	"fmt"
	"testing"
	"time"
)

func TestTicker(t *testing.T) {
	ticker := time.NewTicker(1 * time.Second)

	go func() {
		time.Sleep(5 * time.Second)
		ticker.Stop()
	}()

	// kalo kayak gini jangan pake for range, mungkin bisa pake select & jangan lupa close channelnya
	// error deadlock
	for time := range ticker.C {
		fmt.Println(time)
	}

}

// hanya mengembalikan channelnya saja tanpa ticker
func TestTick(t *testing.T) {
	channel := time.Tick(1 * time.Second)

	for time := range channel {
		fmt.Println(time)
	}

}
