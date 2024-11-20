package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	dontKillMeOrIKillYou()

	var notifyTime time.Duration = 1
	fmt.Printf("Notifying every %d minutes.\n", notifyTime)
	fmt.Println("Taking care of your kidneys...")

	ticker := time.NewTicker(notifyTime * time.Minute)
	drinkWaterMyFriend()
	for range ticker.C {
		if !inTimeRange() {
			continue
		}

		drinkWaterMyFriend()
	}
}

func dontKillMeOrIKillYou() {
	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-signalChannel
		sendNotification("Do you want to open a mine?")
		fmt.Println("Take care of your kidney stone.")
		os.Exit(0)
	}()
}

func inTimeRange() bool {
	newLayout := "15:04"
	start, _ := time.Parse(newLayout, "08:00")
	end, _ := time.Parse(newLayout, "19:00")

	actualHour := time.Now().Format(newLayout)
	now, _ := time.Parse(newLayout, actualHour)
	if now.Minute() == 0 || now.Minute() == 30 {
		return start.Before(now) && end.After(now)
	}

	return false
}

func drinkWaterMyFriend() {
	fmt.Printf("(%s) Taking care of your kidneys!!!\n", time.Now().Format(time.DateTime))
	sendNotification("Drink water!!!")
	openYoutubeVideo()
}

func sendNotification(args ...string) {
	runCommand("notify-send", args...)
}

func openYoutubeVideo() {
	youtubeVideoIds := []string{
		"S2lal37EiqU",
		"YrwBDSUhV-c",
		"4xgN8mTByZ8",
	}

	index := rand.Intn(len(youtubeVideoIds))
	runCommand("/usr/bin/firefox", fmt.Sprintf("https://www.youtube.com/embed/%s?autoplay=1", youtubeVideoIds[index]))
}

func runCommand(executable string, arg ...string) {
	exec.
		Command(executable, arg...).
		Run()
}
