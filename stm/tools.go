package stm

import(
	"strconv"
	"strings"
)

func isDigit(symbol string) bool {
	_, err := strconv.Atoi(symbol)
	if err == nil {
		return true
	}

	return false
}

func getSymbolType(symbol string) int {
	var result int;
	
	result = PrintableSymbol
	if isDigit(symbol) {
		result = DigitSymbol
	} else if EqualValues == strings.Compare(symbol, `\`) {
		result = BackSlashSymbol
	} 

	return result
}

func repeat(symbol string, quantity int) string {
	var result string
	for i := 0; i < quantity; i++ {
		result += symbol
	}

	return result
}