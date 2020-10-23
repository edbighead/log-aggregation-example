package main

import (
	"math/rand"
	"os"
	"time"

	log "github.com/sirupsen/logrus"
)

func random(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max-min) + min
}

func InsertLog() {
	pages := []string{"/login", "/account", "/view/friends", "/logout", "/feedback"}
	users := []string{"Ben", "John", "Bill", "Jane", "Stacy"}
	ips := []string{"2.19.84.119", "2.22.126.12", "13.107.60.25", "91.228.151.55", "178.17.172.140"}

	rand.Seed(time.Now().UnixNano())

	np := rand.Int() % len(pages)
	page := pages[np]

	nu := rand.Int() % len(users)
	user := users[nu]

	iu := rand.Int() % len(ips)
	ip := ips[iu]

	myrand := random(1, 7)

	switch myrand {
	case 1:
		log.WithFields(log.Fields{
			"ip":   ip,
			"user": user,
			"page": page,
		}).Debug()
	case 2:
		log.WithFields(log.Fields{
			"ip":   ip,
			"user": user,
			"page": page,
		}).Warn()
	case 6:
		log.WithFields(log.Fields{
			"ip":   ip,
			"user": user,
			"page": page,
		}).Error()
	default:
		log.WithFields(log.Fields{
			"ip":   ip,
			"user": user,
			"page": page,
		}).Info()

	}

}

func waitFor() {
	myrand := random(1, 10)
	time.Sleep(time.Duration(myrand) * time.Second)
	InsertLog()
}

func main() {
	file, err := os.OpenFile("/var/log/app/info.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	customFormatter := new(log.TextFormatter)
	customFormatter.TimestampFormat = "2006-01-02 15:04:05"
	customFormatter.FullTimestamp = true

	log.SetOutput(file)
	log.SetFormatter(customFormatter)
	log.SetLevel(log.DebugLevel)

	for {
		waitFor()
	}

}
