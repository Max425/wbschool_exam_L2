package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"sort"
	"strconv"
)

/*
=== Утилита sort ===

Отсортировать строки (man sort)
Основное

# Поддержать ключи

-k — указание колонки для сортировки
-n — сортировать по числовому значению
-r — сортировать в обратном порядке
-u — не выводить повторяющиеся строки

# Дополнительное

# Поддержать ключи

-M — сортировать по названию месяца
-b — игнорировать хвостовые пробелы
-c — проверять отсортированы ли данные
-h — сортировать по числовому значению с учётом суффиксов

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

type Flags struct {
	N, R, U bool
	K       int
	Input   string
	Output  string
}

func ParserFlags() (*Flags, error) {
	flags := &Flags{}
	flag.IntVar(&flags.K, "k", 1, "указание колонки для сортировки (слова в строке могут выступать в качестве колонок, по умолчанию разделитель — пробел)")
	flag.BoolVar(&flags.N, "n", false, "сортировать по числовому значению")
	flag.BoolVar(&flags.R, "r", false, "сортировать в обратном порядке")
	flag.BoolVar(&flags.U, "u", false, "не выводить повторяющиеся строки")
	flag.Parse()
	flags.Input = flag.Arg(0)
	flags.Output = flag.Arg(1)
	flags.K--
	return flags, nil
}

func main() {
	flags, err := ParserFlags()
	if err != nil {
		fmt.Println("Ошибка:", err)
		os.Exit(1)
	}
	inputFile, err := os.Open(flags.Input)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer inputFile.Close()

	scanner := bufio.NewScanner(inputFile)
	lines := make([]string, 0)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	lines = Sort(flags, lines)

	outputFile, err := os.Create(flags.Output)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer outputFile.Close()

	writer := bufio.NewWriter(outputFile)
	for _, line := range lines {
		fmt.Fprintln(writer, line)
	}
	writer.Flush()
}

func Sort(flags *Flags, lines []string) []string {
	if flags.N {
		sort.SliceStable(lines, func(i, j int) bool {
			val1, err1 := strconv.Atoi(getColumn(lines[i], flags.K))
			val2, err2 := strconv.Atoi(getColumn(lines[j], flags.K))
			if err1 != nil || err2 != nil {
				return lines[i] < lines[j]
			}
			return val1 < val2
		})
	} else {
		sort.SliceStable(lines, func(i, j int) bool {
			return getColumn(lines[i], flags.K) < getColumn(lines[j], flags.K)
		})
	}

	if flags.R {
		reverseSlice(lines)
	}

	if flags.U {
		lines = unique(lines)
	}
	return lines
}

func getColumn(line string, k int) string {
	columns := splitColumns(line, ' ')
	if k >= len(columns) {
		return ""
	}
	return columns[k]
}

func splitColumns(s string, sep rune) []string {
	var fields []string
	fieldStart := -1
	for i, c := range s {
		if c == sep {
			if fieldStart != -1 {
				fields = append(fields, s[fieldStart:i])
				fieldStart = -1
			}
		} else if fieldStart == -1 {
			fieldStart = i
		}
	}
	if fieldStart != -1 {
		fields = append(fields, s[fieldStart:])
	}
	return fields
}

func reverseSlice(slice []string) {
	for i := 0; i < len(slice)/2; i++ {
		j := len(slice) - i - 1
		slice[i], slice[j] = slice[j], slice[i]
	}
}

func unique(slice []string) []string {
	uniqueMap := make(map[string]bool)
	uniqueSlice := make([]string, 0)
	for _, v := range slice {
		if !uniqueMap[v] {
			uniqueMap[v] = true
			uniqueSlice = append(uniqueSlice, v)
		}
	}
	return uniqueSlice
}
