package main

import (
	"fmt"
	"os"

	"github.com/beevik/ntp"
)

/*
=== Базовая задача ===

Создать программу печатающую точное время с использованием NTP библиотеки.Инициализировать как go module.
Использовать библиотеку https://github.com/beevik/ntp.
Написать программу печатающую текущее время / точное время с использованием этой библиотеки.

Программа должна быть оформлена с использованием как go module.
Программа должна корректно обрабатывать ошибки библиотеки: распечатывать их в STDERR и возвращать ненулевой код выхода в OS.
Программа должна проходить проверки go vet и golint.
*/
const ntpServer = "0.beevik-ntp.pool.ntp.org"

func getTime(ntpServer string) error {
	// Запрашиваем точное время с NTP-сервера
	time, err := ntp.Time(ntpServer)
	// Проверяем, есть ли ошибка при запросе времени
	if err != nil {
		return err
	}

	fmt.Println(time)

	return nil
}

func main() {
	err := getTime(ntpServer)
	if err != nil {
		// В случае ошибки выводим её в STDERR и завершаем программу с ненулевым кодом возврата
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}
}
