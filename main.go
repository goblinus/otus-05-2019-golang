package main

import(
	"fmt"
	"log"
	"os"

	"github.com/otus-05-2019-golang/shortener"
)

const ErrorStatus = 1

func main() {

	urls := []string{
		"https://wwww.mail.ru/directory/subdirectory/",
		"https://google.com/some/resource/placed/here/",
		"https://disk.yandex.ru/the/valued/info/",

	}

	storage := shortener.CreateStorage()
	logger  := log.New(os.Stderr, "Ошибка: ", log.Ldate|log.Ltime)
	
	if storage == nil {
		logger.Println("Хранилище ссылок не создано")
		os.Exit(ErrorStatus)
	}

	//Создаем короткие ссылки для каждого из адресов, указанных в массиве,
	//и сохраняем в storage
	for _, url := range urls {
		record := shortener.CreateRecord()
		if record == nil {
			continue
		}

		if shortLink := record.Shorten(url); len(record.GetShortLink()) > shortener.EmptyValue {
			key := record.GetShortLink()
			storage[key] = record
			
			fmt.Printf("Исходный URL: %s\n", url)
			fmt.Printf("\tКороткая ссылка: %s\n", shortLink)
			fmt.Printf("\tВосстановленная ссылка: %s\n\n", record.Resolve(key))
		}
	}
}
