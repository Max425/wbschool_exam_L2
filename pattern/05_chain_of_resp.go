package main

import "fmt"

/*
	Реализовать паттерн «цепочка вызовов».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Chain-of-responsibility_pattern

Цепочка обязанностей — это поведенческий паттерн проектирования, который позволяет передавать запросы
последовательно по цепочке обработчиков. Каждый последующий обработчик решает, может ли он обработать
запрос сам и стоит ли передавать запрос дальше по цепи.

Применимость
	Когда программа должна обрабатывать разнообразные запросы несколькими способами, но заранее неизвестно,
		какие конкретно запросы будут приходить и какие обработчики для них понадобятся.
	Когда важно, чтобы обработчики выполнялись один за другим в строгом порядке.
	Когда набор объектов, способных обработать запрос, должен задаваться динамически.

Преимущества
	Контроль порядка обработки запросов
	Уменьшает зависимость между клиентом и обработчиками.
	Принцип единой ответственности: Можете отделить классы, вызывающие операции, от классов, выполняющих операции.
	Принцип открытости/закрытости: Можно вводить в приложение новые обработчики, не ломая существующий код.

Недостатки
	Запрос может остаться никем не обработанным.
*/

// Информация о клиенте
type Patient struct {
	name              string
	registrationDone  bool
	doctorCheckUpDone bool
	medicineDone      bool
	paymentDone       bool
}

// Интерфейс для выполнения операций в разных отделах
type Department interface {
	execute(*Patient)
	setNext(Department)
}

// Регистратура
type Reception struct {
	next Department
}

func (r *Reception) execute(p *Patient) {
	if p.registrationDone {
		fmt.Println("Patient registration already done")
		r.next.execute(p)
		return
	} else {
		fmt.Println("Reception registering patient")
		p.registrationDone = true
		r.next.execute(p)
	}
}

func (r *Reception) setNext(next Department) {
	r.next = next
}

// Врач
type Doctor struct {
	next Department
}

func (d *Doctor) execute(p *Patient) {
	if p.doctorCheckUpDone {
		fmt.Println("Doctor checkup already done")
		d.next.execute(p)
		return
	} else {
		fmt.Println("Doctor checking patient")
		p.doctorCheckUpDone = true
		d.next.execute(p)
	}
}

func (d *Doctor) setNext(next Department) {
	d.next = next
}

// Медицинский отдел
type Medical struct {
	next Department
}

func (m *Medical) execute(p *Patient) {
	if p.medicineDone {
		fmt.Println("Medicine already given to patient")
		m.next.execute(p)
		return
	} else {
		fmt.Println("Medical giving medicine to patient")
		p.medicineDone = true
		m.next.execute(p)
	}
}

func (m *Medical) setNext(next Department) {
	m.next = next
}

// Касса
type Cashier struct {
	next Department
}

func (c *Cashier) execute(p *Patient) {
	if p.paymentDone {
		fmt.Println("Payment Done")
		return
	} else {
		fmt.Println("Patient paying to the cashier")
		p.paymentDone = true
	}
}

func (c *Cashier) setNext(next Department) {
	c.next = next
}

func main() {
	cashier := &Cashier{}

	// Указываем следующий отдел (касса)
	medical := &Medical{}
	medical.setNext(cashier)

	// Указываем следующий отдел (медицинский)
	doctor := &Doctor{}
	doctor.setNext(medical)

	// Указываем следующий отдел (доктор)
	reception := &Reception{}
	reception.setNext(doctor)

	patient := &Patient{name: "Dave"}
	// Пациент пришёл в регистратуру
	reception.execute(patient)
}
