package main

import (
	"flag"
	"fmt"
	"github.com/gregdel/pushover"
	"os"
	"strconv"
	"strings"
	"time"
)

func mustAtoi(s string) (i int) {
	i, _ = strconv.Atoi(s)
	return
}

var (
	argToken           = flag.String("token", os.Getenv("PUSHOVER_TOKEN"), "Pushover application token")
	argUser            = flag.String("user", os.Getenv("PUSHOVER_USER"), "Pushover user token")
	argTitle           = flag.String("title", os.Getenv("PUSHOVER_TITLE"), "Notification title")
	argURL             = flag.String("url", os.Getenv("PUSHOVER_URL"), "Notification url")
	argURLTitle        = flag.String("url_title", os.Getenv("PUSHOVER_URL_TITLE"), "Notification url title")
	argDevices         = flag.String("devices", os.Getenv("PUSHOVER_DEVICES"), "Notification devices comma separated")
	argPriority        = flag.Int("priority", mustAtoi(os.Getenv("PUSHOVER_PRIORITY")), "Notification priority")
	argSound           = flag.String("sound", os.Getenv("PUSHOVER_SOUND"), "Notification sound")
	argTimestamp       = flag.String("timestamp", os.Getenv("PUSHOVER_TIMESTAMP"), "Notification timestamp")
	argTimestampFormat = flag.String("timestamp_format", os.Getenv("PUSHOVER_TIMESTAMP"), "Notification timestamp format, e.g. Mon Jan 2 15:04:05 -0700 MST 2006")
)

func main() {
	flag.Parse()

	app := pushover.New(*argToken)
	recipient := pushover.NewRecipient(*argUser)
	msg := pushover.NewMessage(strings.Join(flag.Args(), " "))
	if *argTitle != "" {
		msg.Title = *argTitle
	}
	if *argURL != "" {
		msg.URL = *argURL
	}
	if *argURLTitle != "" {
		msg.URLTitle = *argURLTitle
	}
	if *argDevices != "" {
		msg.DeviceName = *argDevices
	}
	msg.Priority = *argPriority
	if *argSound != "" {
		msg.Sound = *argSound
	}
	if *argTimestamp != "" {
		ts, err := time.Parse(*argTimestampFormat, *argTimestamp)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error parsing timestamp: %s\n", err)
			os.Exit(1)
		}
		msg.Timestamp = ts.Unix()
	}
	resp, err := app.SendMessage(msg, recipient)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s", err)
		os.Exit(1)
	}
	if len(resp.Errors) > 0 {
		fmt.Fprintf(os.Stderr, "Error: %s", err)
		os.Exit(1)
	}
}
