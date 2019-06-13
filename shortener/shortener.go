package shortener

import(
	"regexp"
	_ "fmt"
	"math/rand"
	"time"
	"strings"
)

const alphabet = "123456789bcdfghjkmnpqrstvwxyzBCDFGHJKLMNPQRSTVWXYZ" 
const LinkLength = 14
const EmptyValue = 0
const FirstElement  = 0
const SecondElement = 1
const EqualValues 	= 0

type Shortener interface {
	Shorten(url string) string
	Resolve(url string) string
}

type record struct {
	base string			// Переменная для  хранение базового адреса ссылки
	rawLink string		// Ссылка (относительная) 
	shortLink string	// Короткая ссылка
}

var storage map[string]*record

//Фабричный метод для создания записи типа record
func CreateRecord() *record {
	return &record{}
}

//Фабричный метод для создания словаря для хранения записей record, ключом является короткая ссылка
func CreateStorage() map[string]*record {
	return make(map[string]*record)
}

func (r *record) Shorten(url string) string {
	base, relLink := parseUrl(url)
	if len(base) == EmptyValue {
		return ""
	}

	r.base  	= base
	r.rawLink 	= relLink
	alphabetCapacity := len(alphabet)
	
	result := []byte{}
	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < LinkLength; i += 1 {
		key := random.Intn(alphabetCapacity)
		result = append(result, alphabet[key])
	}

	r.shortLink = string(result)
	return r.shortLink
}

func (r *record) Resolve(url string) string {
	if EqualValues == strings.Compare(r.shortLink, url) {
		return r.rawLink
	}

	return ""
}

//Парсим url на составляющие
func parseUrl(url string) (string, string) {
	rexp := regexp.MustCompile(`(.+\.[a-zA-Z]{2,4})`)
	base := rexp.FindString(url)
	relLink := rexp.ReplaceAllString(url, "")

	if relLink[FirstElement] == '/' {
		relLink = relLink[SecondElement:]
	}

	return base, relLink
}
