package main

import (
	"errors"
	"fmt"
	"strings"
)

/*
	Реализовать паттерн «фабричный метод».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Factory_method_pattern

Фабричный метод — это порождающий паттерн проектирования, который определяет общий интерфейс для создания
объектов в суперклассе, позволяя подклассам изменять тип создаваемых объектов.

Применимость
	Заранее не известны точные типы и зависимости объектов
	Возможность расширения внутренних компонентов
	Экономия системных ресурсов, повторно используя существующие объекты

Плюсы
	позволяет сделать код создания объектов более универсальным, не привязываясь к конкретным классам
	Принцип единой ответственности: Можно переместить код создания продукта в одно место программы, что упростит поддержку кода.
	Принцип открытости/закрытости: Можно вводить в программу новые виды продуктов, не нарушая существующий код.

Минусы
	Может привести к созданию больших параллельных иерархий классов, так как для каждого класса продукта надо создать свой подкласс создателя.
	Код может стать более сложным: В небольших программах добавление фабричных методов может привести к избыточности кода и усложнению его структуры
*/

type Person interface{}

type person struct {
	Name     string
	Position string
	Salary   int
}

type director struct {
	person
}

type TeamLead struct {
	person
}

// GetPerson порождающий объекты метод, центр самого паттерна
func GetPerson(worker string) (Person, error) {
	switch strings.ToLower(worker) {
	case "director":
		return &director{}, nil
	case "teamlead":
		return &TeamLead{}, nil
	}
	return nil, errors.New("unknown type person")
}

func main() {
	person, err := GetPerson("director")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("%T", person)
	}
}
