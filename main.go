package main

import (
	"fmt"
	"log"
	"os"
	"time"
	
	"github.com/beevik/ntp"
)

const TimeSyncServer = "ntp2.stratum2.ru"
const ErrorCode = 1

var MMap = map[time.Month]string{
	time.January:   "Января",
	time.February:  "Февраля",
	time.March:     "Марта",
	time.April:     "Апреля",
	time.May:       "Мая",
	time.June:      "Июня",
	time.July:      "Июля",
	time.August:    "Августа",
	time.September: "Сентября",
	time.October:   "Октября",
	time.November:  "Ноября",
	time.December:  "Декабря",
}

func main() {

	ntpResponse, err := ntp.Query(TimeSyncServer)
	logger := log.New(os.Stderr, "", 0)

	if err != nil {
		logger.Println(err)
		os.Exit(ErrorCode)
	}

	dt := ntpResponse.Time

	fmt.Println("Точное время: ", ntpResponse.Time.Local().String())
	fmt.Printf("Текущее время: %02d %s %02d %02d:%02d\n",
		dt.Day(),
		MMap[dt.Month()],
		dt.Year(),
		dt.Local().Hour(),
		dt.Local().Minute())
}
