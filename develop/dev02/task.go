package main

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

/*
=== Задача на распаковку ===

Создать Go функцию, осуществляющую примитивную распаковку строки, содержащую повторяющиеся символы / руны, например:
  - "a4bc2d5e" => "aaaabccddddde"
  - "abcd" => "abcd"
  - "45" => "" (некорректная строка)
  - "" => ""

Дополнительное задание: поддержка escape - последовательностей
  - qwe\4\5 => qwe45 (*)
  - qwe\45 => qwe44444 (*)
  - qwe\\5 => qwe\\\\\ (*)

В случае если была передана некорректная строка функция должна возвращать ошибку. Написать unit-тесты.

Функция должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/
func UnpackString(str string) (string, error) {
	// если это число
	if _, err := strconv.Atoi(str); err == nil {
		return "", errors.New("некорректная строка")
	}

	var res strings.Builder // буфер для результата
	var pred rune           // переменная для предыдущего символа
	var isEscape bool       // флаг на эскейп последовательность
	for _, char := range str {
		if !isEscape && unicode.IsDigit(char) {
			num, _ := strconv.Atoi(string(char))
			repeat := strings.Repeat(string(pred), num-1) // делаем репит символа
			res.WriteString(repeat)                       // пишем в буфер
		} else {
			isEscape = string(char) == "\\" && string(pred) != "\\"
			if !isEscape {
				res.WriteRune(char)
			}
			pred = char
		}
	}
	return res.String(), nil
}
