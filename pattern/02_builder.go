package main

import "fmt"

/*
	Реализовать паттерн «строитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Builder_pattern

Строитель — это порождающий паттерн проектирования, который позволяет создавать сложные объекты пошагово. Строитель даёт
возможность использовать один и тот же код строительства для получения разных представлений объектов.

Применимость
	Когда вы хотите избавиться от «сложного конструктора».
	Когда ваш код должен создавать разные представления какого-то объекта. Например, деревянные и железобетонные дома.
	Когда вам нужно собирать сложные составные объекты

Преимущества
	Позволяет создавать объекты пошагово.
	Позволяет использовать один и тот же код для создания различных объектов.
	Изолирует сложный код сборки объекта от его основной бизнес-логики.

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
	if builderType == "wood" {
		return newWoodBuilder()
	}

	if builderType == "stone" {
		return newStoneBuilder()
	}
	return nil
}

type WoodBuilder struct {
	windowType string
	doorType   string
	floor      int
}

func newWoodBuilder() *WoodBuilder {
	return &WoodBuilder{}
}

func (b *WoodBuilder) setWindowType() {
	b.windowType = "Wooden Window"
}

func (b *WoodBuilder) setDoorType() {
	b.doorType = "Wooden Door"
}

func (b *WoodBuilder) setNumFloor() {
	b.floor = 2
}

func (b *WoodBuilder) getHouse() House {
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
	woodBuilder := getBuilder("wood")
	stoneBuilder := getBuilder("stone")

	director := newDirector(woodBuilder)
	woodHouse := director.buildHouse()

	fmt.Printf("Wood House Door Type: %s\n", woodHouse.doorType)
	fmt.Printf("Wood House Window Type: %s\n", woodHouse.windowType)
	fmt.Printf("Wood House Num Floor: %d\n", woodHouse.floor)

	director.setBuilder(stoneBuilder)
	stoneHouse := director.buildHouse()

	fmt.Printf("\nStone House Door Type: %s\n", stoneHouse.doorType)
	fmt.Printf("Stone House Window Type: %s\n", stoneHouse.windowType)
	fmt.Printf("Stone House Num Floor: %d\n", stoneHouse.floor)

}
