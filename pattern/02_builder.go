package main

import "fmt"

/*
	Реализовать паттерн «строитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Builder_pattern

Строитель — это порождающий паттерн проектирования, который позволяет создавать сложные объекты пошагово. Строитель даёт
возможность использовать один и тот же код строительства для получения разных представлений объектов.

Применимость
	Когда вы хотите избавиться от «телескопического конструктора».
	Когда ваш код должен создавать разные представления какого-то объекта. Например, деревянные и железобетонные дома.
	Когда вам нужно собирать сложные составные объекты

Преимущества
	Позволяет создавать продукты пошагово.
	Позволяет использовать один и тот же код для создания различных продуктов.
	Изолирует сложный код сборки продукта от его основной бизнес-логики.

Недостатки
	Усложняет код программы из-за введения дополнительных классов.
	Клиент будет привязан к конкретным классам строителей, так как в интерфейсе директора может не быть метода получения результата.
*/

type Builder interface {
	setWindowType()
	setDoorType()
	setNumFloor()
	getHouse() House
}

func getBuilder(builderType string) Builder {
	if builderType == "normal" {
		return newNormalBuilder()
	}

	if builderType == "stone" {
		return newStoneBuilder()
	}
	return nil
}

type NormalBuilder struct {
	windowType string
	doorType   string
	floor      int
}

func newNormalBuilder() *NormalBuilder {
	return &NormalBuilder{}
}

func (b *NormalBuilder) setWindowType() {
	b.windowType = "Wooden Window"
}

func (b *NormalBuilder) setDoorType() {
	b.doorType = "Wooden Door"
}

func (b *NormalBuilder) setNumFloor() {
	b.floor = 2
}

func (b *NormalBuilder) getHouse() House {
	return House{
		doorType:   b.doorType,
		windowType: b.windowType,
		floor:      b.floor,
	}
}

type StoneBuilder struct {
	windowType string
	doorType   string
	floor      int
}

func newStoneBuilder() *StoneBuilder {
	return &StoneBuilder{}
}

func (b *StoneBuilder) setWindowType() {
	b.windowType = "Stone Window"
}

func (b *StoneBuilder) setDoorType() {
	b.doorType = "Stone Door"
}

func (b *StoneBuilder) setNumFloor() {
	b.floor = 1
}

func (b *StoneBuilder) getHouse() House {
	return House{
		doorType:   b.doorType,
		windowType: b.windowType,
		floor:      b.floor,
	}
}

type House struct {
	windowType string
	doorType   string
	floor      int
}

type Director struct {
	builder Builder
}

func newDirector(b Builder) *Director {
	return &Director{
		builder: b,
	}
}

func (d *Director) setBuilder(b Builder) {
	d.builder = b
}

func (d *Director) buildHouse() House {
	d.builder.setDoorType()
	d.builder.setWindowType()
	d.builder.setNumFloor()
	return d.builder.getHouse()
}

func main() {
	normalBuilder := getBuilder("normal")
	stoneBuilder := getBuilder("stone")

	director := newDirector(normalBuilder)
	normalHouse := director.buildHouse()

	fmt.Printf("Normal House Door Type: %s\n", normalHouse.doorType)
	fmt.Printf("Normal House Window Type: %s\n", normalHouse.windowType)
	fmt.Printf("Normal House Num Floor: %d\n", normalHouse.floor)

	director.setBuilder(stoneBuilder)
	stoneHouse := director.buildHouse()

	fmt.Printf("\nStone House Door Type: %s\n", stoneHouse.doorType)
	fmt.Printf("Stone House Window Type: %s\n", stoneHouse.windowType)
	fmt.Printf("Stone House Num Floor: %d\n", stoneHouse.floor)

}
