package main

import (
	"fmt"
	"time"

	"github.com/beevik/ntp"
)

const TimeSyncServer = "ntp2.stratum2.ru"

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

	if err != nil {
		fmt.Println(err)
		return
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
