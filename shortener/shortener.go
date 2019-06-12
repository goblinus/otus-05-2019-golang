package shorterner

const alphabet = "123456789bcdfghjkmnpqrstvwxyzBCDFGHJKLMNPQRSTVWXYZ" 

type Shortener interface {
	Shorten(url string) string
	Resolve(url string) string
}

type data struct {
	var storage map[string] string
	var alphabet string
}

func (d *data) Shorten(url string) string {

}

func (d *data) Resolve(url string) string {

}

func New() *data {
	source := alphabet
	result := make(map[string] string)
	
	return &result
}