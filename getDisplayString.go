package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

var (
	drool = []string{"drool", "🤤"}
	smile = []string{"smile", "🙂"}
	think = []string{"think", "🤔"}
	frown = []string{"frown", "🙁"}
	clown = []string{"clown", "🤡"}
	poop  = []string{"poop", "💩"}
)

func emoteSelector(bh time.Time, emoji bool) string {
	selector := 0
	if emoji {
		selector = 1
	}
	emote := ""
	switch tt := time.Until(bh).Hours() / 24; {
	case tt >= 100: // x > 100
		emote = drool[selector]
		break
	case tt >= 50: // x < 100
		emote = smile[selector]
		break
	case tt >= 30: // x < 50
		emote = think[selector]
		break
	case tt >= 10: // x < 30
		emote = frown[selector]
		break
	case tt >= 3: // x < 10
		emote = clown[selector]
		break
	case tt >= 1: // x < 3
		emote = poop[selector]
		break
	}
	return emote
}

func getDisplayString(emoji bool) (loc, level, bh, emote string) {
	time.Sleep(1 * time.Second)
	user, err := client.GetUser(os.Getenv("LOGNAME"))
	if err != nil {
		log.Fatal(err)
	}

	if user.Location == nil {
		user.Location = "N/A"
	}

	loc = fmt.Sprintf("%s@%s", user.Login, user.Location)
	if user.CursusUsers != nil && len(user.CursusUsers) > 0 && len(os.Args) >= 2 {
		for _, cursus := range user.CursusUsers {
			if cursus.Cursus.Slug == conf.Cursus {
				level = fmt.Sprintf("Level: %.2f", cursus.Level)
				if time.Until(cursus.BlackholedAt) > 0 {
					bh = fmt.Sprintf("Blackhole: %.0f days",
						time.Until(cursus.BlackholedAt).Hours()/24)
					emote = emoteSelector(cursus.BlackholedAt, emoji)
				}
			}
		}
	}
	return
}
