package main

import "fmt"

/*
	Реализовать паттерн «посетитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Visitor_pattern

Посетитель — это поведенческий паттерн проектирования, который позволяет добавлять в программу новые операции,
не изменяя классы объектов, над которыми эти операции могут выполняться.

Применимость
	Когда вам нужно выполнить какую-то операцию над всеми элементами сложной структуры объектов, например, деревом.
	Когда над объектами сложной структуры объектов надо выполнять некоторые не связанные между собой операции,
		но вы не хотите «засорять» классы такими операциями.
	Когда новое поведение имеет смысл только для некоторых классов из существующей иерархии.


Преимущества
	Упрощает добавление операций, работающих со сложными структурами объектов.
	Объединяет родственные операции в одном классе.
	Посетитель может накапливать состояние при обходе структуры элементов.

Недостатки
	Паттерн не оправдан, если иерархия элементов часто меняется.
	Может привести к нарушению инкапсуляции элементов.
*/

type Shape interface {
	GetType() string
	Accept(Visitor)
}

type Circle struct {
	radius int
}

func (c *Circle) GetType() string {
	return "Circle"
}

func (c *Circle) Accept(v Visitor) {
	v.VisitForCircle(c)
}

type Rectangle struct {
	height int
	length int
}

func (r *Rectangle) GetType() string {
	return "Rectangle"
}

func (r *Rectangle) Accept(v Visitor) {
	v.VisitForRectangle(r)
}

type Visitor interface {
	VisitForRectangle(*Rectangle)
	VisitForCircle(*Circle)
}

type AreaCalculate struct{}

func (a *AreaCalculate) VisitForCircle(c *Circle) {
	fmt.Println("Calculate area for", c.GetType())
}

func (a *AreaCalculate) VisitForRectangle(r *Rectangle) {
	fmt.Println("Calculate area for", r.GetType())
}

func main() {
	circle := &Circle{}
	react := &Rectangle{}
	circle.Accept(&AreaCalculate{})
	react.Accept(&AreaCalculate{})
}
