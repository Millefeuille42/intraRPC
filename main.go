package main

import (
	"fmt"
	rpcClient "github.com/hugolgst/rich-go/client"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var (
	url    = "https://api.intra.42.fr"
	uid    = os.Getenv("APP_UID")
	secret = os.Getenv("APP_SECRET")
	client = &APIClient{Url: url, Uid: uid, Secret: secret}

	statuses = [4]string{"Activity", "À table", "En pause", "Aux toilettes"}
	index    = 0

	done chan bool
)

func checkEnv() bool {
	return os.Getenv("APP_UID") == "" || os.Getenv("APP_SECRET") == "" || os.Getenv("LOGNAME") == ""
}

func setActivity(loc, level, bh, emote string) {
	tt := time.Now()

	err := rpcClient.SetActivity(rpcClient.Activity{
		Details:    loc,
		State:      level + " - " + bh,
		LargeImage: "largeimage",
		SmallImage: emote,
		Timestamps: &rpcClient.Timestamps{
			Start: &tt,
		},
	})

	if err != nil {
		panic(err)
	}
}

func activityPicker() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGUSR1, syscall.SIGUSR2, syscall.SIGINT)

	for {
		sig := <-sigs
		switch sig {
		case syscall.SIGUSR1:
			index = (index + 1) % 4
			break
		case syscall.SIGUSR2:
			loc, level, bh, emote := getDisplayString(false)
			if statuses[index] != "Activity" {
				loc = statuses[index]
			}
			setActivity(loc, level, bh, emote)
			index = 0
			break
		case syscall.SIGINT:
			done <- true
			break
		}
	}
}

func main() {
	if checkEnv() {
		log.Fatal("missing env values")
	}

	if len(os.Args) < 2 {
		log.Fatal("not enough arguments")
	}

	if err := client.Auth(); err != nil {
		log.Fatal(err.Error())
	}

	if len(os.Args) == 3 && os.Args[2] == "--print" {
		loc, level, bh, emote := getDisplayString(true)
		fmt.Printf("%s - %s - %s %s", loc, level, bh, emote)
		return
	}

	err := rpcClient.Login("837269730787852299")
	if err != nil {
		panic(err)
	}

	done = make(chan bool, 1)

	setActivity(getDisplayString(false))
	go activityPicker()

	<-done
	fmt.Println("exiting")
	rpcClient.Logout()
}
