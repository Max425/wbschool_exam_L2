package main

import (
	"fmt"
	"log"
)

/*
	Реализовать паттерн «фасад».
Объяснить применимость паттерна, его плюсы и минусы,а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Facade_pattern

Фасад — это структурный паттерн проектирования, который предоставляет простой интерфейс к сложной системе классов,
библиотеке или фреймворку.

Применимость
	Когда вам нужно представить простой или урезанный интерфейс к сложной подсистеме.
	Когда вы хотите разложить подсистему на отдельные слои.

Преимущества
	Изолирует клиентов от компонентов сложной подсистемы.

Недостатки
	Фасад рискует стать божественным объектом, привязанным ко всем классам программы.
*/

type Account struct {
	name string
}

func newAccount(accountName string) *Account {
	return &Account{
		name: accountName,
	}
}

func (a *Account) checkAccount(accountName string) error {
	if a.name != accountName {
		return fmt.Errorf("account Name is incorrect")
	}
	fmt.Println("Account Verified")
	return nil
}

type Wallet struct {
	balance int
}

func (w *Wallet) creditBalance(amount int) {
	w.balance += amount
	fmt.Println("Wallet balance added successfully")
	return
}

type Notification struct {
}

func (n *Notification) sendWalletCreditNotification() {
	fmt.Println("Sending wallet credit notification")
}

type WalletFacade struct {
	account      *Account
	wallet       *Wallet
	notification *Notification
}

func newWalletFacade(accountID string) *WalletFacade {
	walletFacade := &WalletFacade{
		account:      newAccount(accountID),
		wallet:       &Wallet{},
		notification: &Notification{},
	}
	return walletFacade
}

func (w *WalletFacade) addMoneyToWallet(accountID string, amount int) error {
	fmt.Println("Starting add money to wallet")
	err := w.account.checkAccount(accountID)
	if err != nil {
		return err
	}
	w.wallet.creditBalance(amount)
	w.notification.sendWalletCreditNotification()
	return nil
}

func main() {
	fmt.Println()
	walletFacade := newWalletFacade("abc")
	fmt.Println()

	err := walletFacade.addMoneyToWallet("abc", 10)
	if err != nil {
		log.Fatalf("Error: %s\n", err.Error())
	}
}
