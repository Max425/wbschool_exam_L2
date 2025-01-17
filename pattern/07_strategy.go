package main

import (
	"fmt"
	"strings"
)

/*
	Реализовать паттерн «стратегия».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Strategy_pattern

Стратегия — это поведенческий паттерн проектирования, который определяет семейство схожих алгоритмов и помещает
каждый из них в собственный класс, после чего алгоритмы можно взаимозаменять прямо во время исполнения программы.

Применимость
	Когда вам нужно использовать разные вариации какого-то алгоритма внутри одного объекта.
	Когда у вас есть множество похожих классов, отличающихся только некоторым поведением.
	Когда вы не хотите обнажать детали реализации алгоритмов для других классов.
	Когда различные вариации алгоритмов реализованы в виде развесистого условного оператора.
		Каждая ветка такого оператора представляет собой вариацию алгоритма.

Преимущества
	Горячая замена алгоритмов на лету.
	Изолирует код и данные алгоритмов от остальных классов.
	Уход от наследования к делегированию.
	Реализует принцип открытости/закрытости.

Недостатки
	Усложняет программу за счёт дополнительных классов.
	Клиент должен знать, в чём состоит разница между стратегиями, чтобы выбрать подходящую.
*/

type Strategy interface {
	Process(string) string
}

// Метод, который вызывает необходимую стратегию
func RunStrategy(str string, op Strategy) {
	fmt.Println(op.Process(str))
}

type Upper struct{}

func (Upper) Process(str string) string {
	return strings.ToUpper(str)
}

type Lower struct{}

func (Lower) Process(str string) string {
	return strings.ToLower(str)
}

func main() {
	str := "HellO wORld"
	RunStrategy(str, Upper{})
	RunStrategy(str, Lower{})
}
