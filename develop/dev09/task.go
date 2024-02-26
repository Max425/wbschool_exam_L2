package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)

/*
=== Утилита wget ===

# Реализовать утилиту wget с возможностью скачивать сайты целиком

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {
	var urlFlag string
	flag.StringVar(&urlFlag, "url", "", "URL")
	flag.Parse()

	if urlFlag == "" {
		fmt.Fprintln(os.Stderr, "Usage: go run task.go -url <URL>")
		os.Exit(1)
	}

	parseUrl, err := url.Parse(urlFlag)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Ошибка парсинга URL:", err)
		os.Exit(1)
	}

	resp, err := http.Get(urlFlag)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Ошибка скачивания:", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Ошибка чтения тела:", err)
		os.Exit(1)
	}

	outputFile, err := os.Create(parseUrl.Host + ".html")
	if err != nil {
		fmt.Fprintln(os.Stderr, "Не получилось создать файл для результата:", err)
		os.Exit(1)
	}
	defer outputFile.Close()

	_, err = outputFile.Write(body)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Ошибка записи в файл:", err)
		os.Exit(1)
	}

	fmt.Println("Скачен файл:", parseUrl.Host+".html")
}
