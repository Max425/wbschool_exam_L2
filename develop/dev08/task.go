package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
)

/*
=== Взаимодействие с ОС ===

Необходимо реализовать свой собственный UNIX-шелл-утилиту с поддержкой ряда простейших команд:
- cd <args> - смена директории (в качестве аргумента могут быть то-то и то)
- pwd - показать путь до текущего каталога
- echo <args> - вывод аргумента в STDOUT
- kill <args> - "убить" процесс, переданный в качесте аргумента (пример: такой-то пример)
- ps - выводит общую информацию по запущенным процессам в формате *такой-то формат*

Так же требуется поддерживать функционал fork/exec-команд

Дополнительно необходимо поддерживать конвейер на пайпах (linux pipes, пример cmd1 | cmd2 | .... | cmdN).
Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {
	reader := bufio.NewReader(os.Stdin)

LOOP:
	for {
		dir, _ := os.Getwd()
		fmt.Printf("%s$ ", dir)
		input, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Fprintln(os.Stderr, "Ошибка чтения ввода:", err)
			continue
		}
		input = strings.TrimSpace(input)
		if input == "" {
			continue
		}
		args := strings.Fields(input)
		switch args[0] {
		case "":
			continue
		case "exit":
			break LOOP
		case "cd":
			if len(args) < 2 {
				fmt.Fprintln(os.Stderr, "Неверное количество аргументов для cd")
				continue
			}
			if err = os.Chdir(args[1]); err != nil {
				fmt.Fprintln(os.Stderr, "Ошибка при смене директории:", err)
			}
		default:
			cmd := exec.Command(args[0], args[1:]...)
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			if err = cmd.Run(); err != nil {
				fmt.Fprintln(os.Stderr, "Ошибка при выполнении команды:", err)
			}
		}
	}
}
