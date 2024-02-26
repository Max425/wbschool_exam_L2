package main

import (
	"sort"
	"strings"
)

/*
=== Поиск анаграмм по словарю ===

Напишите функцию поиска всех множеств анаграмм по словарю.
Например:
'пятак', 'пятка' и 'тяпка' - принадлежат одному множеству,
'листок', 'слиток' и 'столик' - другому.

Входные данные для функции: ссылка на массив - каждый элемент которого - слово на русском языке в кодировке utf8.
Выходные данные: Ссылка на мапу множеств анаграмм.
Ключ - первое встретившееся в словаре слово из множества
Значение - ссылка на массив, каждый элемент которого, слово из множества. Массив должен быть отсортирован по возрастанию.
Множества из одного элемента не должны попасть в результат.
Все слова должны быть приведены к нижнему регистру.
В результате каждое слово должно встречаться только один раз.

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func FindAnagrams(words []string) map[string][]string {
	anagramSets := make(map[string][]string)

	for _, word := range words {
		// Приводим слово к нижнему регистру и сортируем его буквы
		sortedWord := sortString(strings.ToLower(word))
		// Добавляем слово в соответствующее множество анаграмм
		anagramSets[sortedWord] = append(anagramSets[sortedWord], word)
	}
	// Удаляем множества из одного элемента
	for key, value := range anagramSets {
		if len(value) == 1 {
			delete(anagramSets, key)
		}
	}

	return anagramSets
}

func sortString(s string) string {
	runes := []rune(s)
	sort.Slice(runes, func(i, j int) bool { return runes[i] < runes[j] })
	return string(runes)
}
