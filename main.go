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
	client *APIClient
	conf   config
)

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

func main() {
	conf = parseConfig()

	client = &APIClient{Url: url, Uid: conf.AppUid, Secret: conf.AppSecret}
	if err := client.Auth(); err != nil {
		log.Fatal(err.Error())
	}

	time.Sleep(1 * time.Second)

	if len(os.Args) == 2 && os.Args[1] == "--print" {
		loc, level, bh, emote := getDisplayString(true)
		fmt.Printf("%s - %s - %s %s", loc, level, bh, emote)
		return
	}

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT)
	err := rpcClient.Login("837269730787852299")
	if err != nil {
		panic(err)
	}
	<-sigs

	setActivity(getDisplayString(false))

	fmt.Println("exiting")
	rpcClient.Logout()
}
