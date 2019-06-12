package main

import(
	"fmt"
	"shorterner"
)

 storage := shorterner.New()
 fmt.Println(storage)

// import (
// 	"fmt"
// 	"math/rand"
// 	"time"
// )

// const Alphabet = "bcdfghklmnpqrstvwzBCDFGHKLMNPQRSTVWZ0123456789"
// const LinkLength = 12
// const AlphaLength = len(Alphabet)

// type Url struct {
// 	source    string
// 	shortLink string
// }

// func (u *Url) Shorten(url string) string {
// 	var result []rune
// 	initValue := time.Now().UnixNano()
// 	random := rand.New(rand.NewSource(int64(initValue)))

// 	u.source = url
// 	for counter := 0; counter < LinkLength; counter += 1 {
// 		value := rune(Alphabet[random.Intn(AlphaLength)])
// 		result = append(result, value)
// 	}

// 	u.shortLink = string(result)

// 	return u.shortLink
// }

// func (u *Url) Resolve(url string) string {
// 	if len(u.source) > 0 {
// 		return u.source
// 	}

// 	return ""
// }

// func (u *Url) ShortLink() string {
// 	if len(u.shortLink) > 0 {
// 		return u.shortLink
// 	}

// 	return ""
// }

// func main() {
// 	val1 := &Url{}
// 	val2 := &Url{}
// 	val3 := &Url{}

// 	val1.Shorten("http://wwww.mail.ru/directory/subdirectory/")
// 	val2.Shorten("http://wwww.google.com/directory/subdirectory/")
// 	val3.Shorten("http://wwww.yandex.ru/directory/subdirectory/")

// 	fmt.Println(": ", val1.ShortLink())
// 	fmt.Println(": ", val2.ShortLink())
// 	fmt.Println(": ", val3.ShortLink())
	// initValue := time.Now().Second()
	// r := rand.New(rand.NewSource(int64(initValue)))

	// re, err := regexp.Compile(`.*\.[a-zA-Z]{2,4}`)
	// if err != nil {
	// 	fmt.Printf(fmt.Sprintf("Невозможно создать объект для извлечения домена из исходной строки: %s", err))
	// 	os.Exit(1)
	// }

	// domain := re.FindStringSubmatch("https://mail.com/part1/part2/")
	// fmt.Println(domain)
	// fmt.Println(r.Intn(10))
	// fmt.Println(r.Intn(10))
	// fmt.Println(r.Intn(10))
	// fmt.Println(r.Intn(10))
}
