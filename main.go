package main

import (
	"fmt"
	"log"
	"time"

	"github.com/go-co-op/gocron"
	"github.com/k23dev/tabarra/src/tabarra"
)

func main() {

	t := tabarra.NewTabarraApp()

	appBanner(t.App.Config.App_name, t.App.Config.App_version)

	if t.IsCronJob {
		msg := fmt.Sprintf("Scrapping as cron job every %d ", t.ExecutionInterval)
		intervalTime := t.ExecutionInterval
		if t.IsIntervalSecods {
			msg = msg + "seconds ..."
		} else {
			intervalTime = intervalTime * 60
			msg = msg + "minutes ..."
		}

		printSeparator(msg)

		cron := gocron.NewScheduler(time.UTC)
		cron.StartAsync()

		_, err := cron.Every(intervalTime).Seconds().Do(func() {
			// HERE COMES THE JOB!
			log.Println("Execute this")
		})

		if err != nil {
			fmt.Println(err)
		}

		cron.StartBlocking()
	} else {
		msg := "Scrapping once"
		printSeparator(msg)
		// HERE COMES THE JOB!
		log.Println("Execute this")

	}
}

func appBanner(appName, version string) {
	fmt.Println("######################################")
	fmt.Printf("    %s v(%s) is running\n", appName, version)
	fmt.Println("######################################")
	fmt.Println("")
}

func printSeparator(text string) {
	fmt.Println()
	fmt.Println(text)
	fmt.Println("--------------------------------")
	fmt.Println()
}
