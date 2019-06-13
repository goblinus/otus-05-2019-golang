package shorterner

import (
	"testing"
	"strings"
	_ "fmt"
)

const base = "http://www.mail.ru"
const relLink 	= "some/link/to/resource/"
const fullLink 	= "http://www.mail.ru/some/link/to/resource/"

func TestShorterner(t *testing.T) {
	record := CreateRecord()
	link := record.Shorten(fullLink)

	if EqualValues != strings.Compare(link, record.shortLink) {
		t.Errorf("Ссылки (та, которая возвращена методом Shorten и record.shortLink) не идентичны")
	}

	if EqualValues != strings.Compare(base, record.base) {
		t.Errorf("Основа (домен) не соответствует образцу")
	}

	if EqualValues != strings.Compare(relLink, record.rawLink) {
		t.Errorf("Исходная ссылка не соответствует образцу")
	}
}
