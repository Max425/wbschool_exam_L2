package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
=== Утилита cut ===

Принимает STDIN, разбивает по разделителю (TAB) на колонки, выводит запрошенные

Поддержать флаги:
-f - "fields" - выбрать поля (колонки)
-d - "delimiter" - использовать другой разделитель
-s - "separated" - только строки с разделителем

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

type Flags struct {
	Fields    string
	Delimiter string
	Separated bool
}

func ParseFlags() *Flags {
	flags := &Flags{}
	flag.StringVar(&flags.Fields, "f", "", "выбрать поля (колонки)")
	flag.StringVar(&flags.Delimiter, "d", "\t", "использовать другой разделитель")
	flag.BoolVar(&flags.Separated, "s", false, "только строки с разделителем")
	flag.Parse()
	return flags
}

func main() {
	flags := ParseFlags()

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()

		res, notSkip := Cut(flags, line)
		if notSkip {
			fmt.Println(res)
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Ошибка при чтении ввода:", err)
		os.Exit(1)
	}
}

func Cut(flags *Flags, line string) (string, bool) {
	// Если указан флаг -s и строка не содержит разделитель, пропускаем её
	if flags.Separated && !strings.Contains(line, flags.Delimiter) {
		return "", false
	}

	fields := strings.Split(line, flags.Delimiter)
	output := make([]string, 0)

	// Выбираем указанные поля (колонки)
	if flags.Fields != "" {
		fieldNumbers := strings.Split(flags.Fields, ",")
		for _, fieldNumber := range fieldNumbers {
			index, _ := strconv.Atoi(fieldNumber)
			output = append(output, fields[index-1])
		}
	} else {
		output = fields
	}

	return strings.Join(output, flags.Delimiter), true
}
