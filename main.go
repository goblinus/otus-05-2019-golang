package main

import(
	"fmt"
	"strings"
	"sort"
	"regexp"
)

const Separator = ` `

type kv struct {
	Key string
	Quantity int
}

func main() {
	source := `Напомним, вечером в пятницу около 1,5 тыс. жителей села Чемодановка вышли на сельский сход, где обсуждался конфликт с цыганами, произошедший 13 июня. Предположительно, селяне вступились за детей, которые поссорились с цыганами. Произошла массовая драка, в результате которой были пострадавшие и один погибший. Позже в полицию были доставлены 174 человека для установления причастности к драке, трое из них были задержаны. По факту произошедшего возбуждены уголовные дела об убийстве и покушении на убийство двух и более лиц, после чего дело было передано в центральный аппарат СКР.
	Кроме того, 15 июня в соседнем с Чемодановкой селе Лопатки неизвестные подожгли деревянный жилой дом и надворные постройки, из-за чего МВД возбудило уголовное дело, фигурантам которого грозит до пяти лет лишения свободы`
	result := []kv{}
	stats  := make(map[string]int)
	buffer := strings.Split(source, Separator)
	
	//Подсчитавыем частоту отдельных слов: предварительно очищаем пробельные символы, точки и запятые
	rexp := regexp.MustCompile(`(\n|\s|\.|\,)`)
	for _, value := range buffer {
		value = rexp.ReplaceAllString(value, "")
		value = strings.ToLower(value)
		if _, ok := stats[value]; false == ok {
			stats[value] = 0
		}

		stats[value]++
	}
	
	//Готовим массив структур для сортировки
	for k, v := range stats {
		result = append(result, kv{Key: k, Quantity: v})
	}

	//Сортируем по убыванию значения kv.Quantity
	sort.Slice(result, func (i, j int) bool {
		return result[i].Quantity > result[j].Quantity
	});

	//Выводим первые 10 самых крупных элементов
	result = result[:10]
	fmt.Printf("%15s   %-10s", "Слово", "Частота упоминаний\n")
	fmt.Printf("------------------------------------\n")
	for _, item := range result {
		fmt.Printf("%15s   %-10d\n", item.Key, item.Quantity)
	}
}
