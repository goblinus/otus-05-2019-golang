package convert

const Ten = 10
const Zero = 0
const Minus = '-'

func Itoa(i int) string {

	//Инициализируем первоначальные данные
	buffer := make([]rune, 0)
	result := make([]rune, 0)
	alphabet := []rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}

	//Определяем, является ли указанная цифра отрицательной или нет
	isNegative := false
	if i < Zero {
		i = -i
		isNegative = true
	}

	//Переводим каждый из показателей порядка из цифрового представления в форму руны
	for i >= Ten {
		rest := i % 10
		buffer = append(buffer, alphabet[rest])
		i = (i - rest) / 10
	}

	//Обрабатываем последний показатель порядка
	buffer = append(buffer, alphabet[i])

	//Учитываем знак исходной цифры, если есть
	if isNegative {
		buffer = append(buffer, Minus)
	}

	//Реверс полученного результата
	for index := len(buffer) - 1; index >= 0; index -= 1 {
		result = append(result, buffer[index])
	}

	return string(result)
}
