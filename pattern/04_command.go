package main

import "fmt"

/*
	Реализовать паттерн «комманда».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Command_pattern

Команда — это поведенческий паттерн проектирования, который превращает запросы в объекты, позволяя передавать их
как аргументы при вызове методов, ставить запросы в очередь, логировать их, а также поддерживать отмену операций.

Применимость
	Когда вы хотите параметризовать объекты выполняемым действием.
	Когда вы хотите ставить операции в очередь, выполнять их по расписанию или передавать по сети.
	Когда вам нужна операция отмены.

Преимущества
	Убирает прямую зависимость между объектами, вызывающими операции, и объектами, которые их непосредственно выполняют.
	Позволяет реализовать простую отмену и повтор операций.
	Позволяет реализовать отложенный запуск операций.
	Позволяет собирать сложные команды из простых.
	Принцип единой ответственности: Классы, вызывающие операции, можно отделить от классов, выполняющих эти операции.
	Принцип открытости/закрытости: Можно вводить в приложение новые команды, не ломая существующий код.

Недостатки
	Усложняет код программы из-за введения множества дополнительных классов.
*/

type Button struct {
	command Command
}

func (b *Button) press() {
	b.command.execute()
}

type Command interface {
	execute()
}

type OnCommand struct {
	device Device
}

func (c *OnCommand) execute() {
	c.device.on()
}

type OffCommand struct {
	device Device
}

func (c *OffCommand) execute() {
	c.device.off()
}

type Device interface {
	on()
	off()
}

type Tv struct {
	isRunning bool
}

func (t *Tv) on() {
	t.isRunning = true
	fmt.Println("Turning tv on")
}

func (t *Tv) off() {
	t.isRunning = false
	fmt.Println("Turning tv off")
}

func main() {
	tv := &Tv{}

	onCommand := &OnCommand{
		device: tv,
	}

	offCommand := &OffCommand{
		device: tv,
	}

	onButton := &Button{
		command: onCommand,
	}
	onButton.press()

	offButton := &Button{
		command: offCommand,
	}
	offButton.press()
}
