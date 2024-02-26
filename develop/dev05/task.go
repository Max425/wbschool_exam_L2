package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"regexp"
	"strings"
)

/*
=== Утилита grep ===

Реализовать утилиту фильтрации (man grep)

Поддержать флаги:
-A - "after" печатать +N строк после совпадения
-B - "before" печатать +N строк до совпадения
-C - "context" (A+B) печатать ±N строк вокруг совпадения
-c - "count" (количество строк)
-i - "ignore-case" (игнорировать регистр)
-v - "invert" (вместо совпадения, исключать)
-F - "fixed", точное совпадение со строкой, не паттерн
-n - "line num", печатать номер строки

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

type Flags struct {
	After, Before, Context                    int
	Count, IgnoreCase, Invert, Fixed, LineNum bool
	Pattern, File                             string
}

func ParserFlags() (*Flags, error) {
	flags := &Flags{}
	flag.IntVar(&flags.After, "A", 0, "печатать +N строк после совпадения")
	flag.IntVar(&flags.Before, "B", 0, "печатать +N строк до совпадения")
	flag.IntVar(&flags.Context, "C", 0, "печатать ±N строк вокруг совпадения")
	flag.BoolVar(&flags.Count, "c", false, "количество строк")
	flag.BoolVar(&flags.IgnoreCase, "i", false, "игнорировать регистр")
	flag.BoolVar(&flags.Invert, "v", false, "вместо совпадения, исключать")
	flag.BoolVar(&flags.Fixed, "F", false, "точное совпадение со строкой, не паттерн")
	flag.BoolVar(&flags.LineNum, "n", false, "напечатать номер строки")
	flag.Parse()

	flags.Pattern = flag.Arg(0)
	flags.File = flag.Arg(1)

	if flags.Pattern == "" || flags.File == "" {
		return nil, fmt.Errorf("необходимо указать шаблон и файл")
	}

	return flags, nil
}

func main() {
	flags, err := ParserFlags()
	if err != nil {
		fmt.Println("Ошибка:", err)
		os.Exit(1)
	}

	file, err := os.Open(flags.File)
	if err != nil {
		fmt.Println("Ошибка открытия файла:", err)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineNum := 0
	for scanner.Scan() {
		line := scanner.Text()
		lineNum++

		if match(flags, line) {
			if flags.LineNum {
				fmt.Printf("%d:", lineNum)
			}
			fmt.Println(line)
		}
	}
	if err = scanner.Err(); err != nil {
		fmt.Println("Ошибка чтения файла:", err)
	}
}

func match(flags *Flags, line string) bool {
	pattern := flags.Pattern

	if flags.Fixed {
		pattern = regexp.QuoteMeta(pattern)
	}

	if flags.IgnoreCase {
		line = strings.ToLower(line)
		pattern = strings.ToLower(pattern)
	}

	matched := strings.Contains(line, pattern)

	if flags.Invert {
		matched = !matched
	}

	return matched
}
