package main

import (
	"context"
	"fmt"
	"time"
)

/*
=== Or channel ===

Реализовать функцию, которая будет объединять один или более done каналов в single канал если один из его составляющих каналов закроется.
Одним из вариантов было бы очевидно написать выражение при помощи select, которое бы реализовывало эту связь,
однако иногда неизестно общее число done каналов, с которыми вы работаете в рантайме.
В этом случае удобнее использовать вызов единственной функции, которая, приняв на вход один или более or каналов, реализовывала весь функционал.
*/

func or(channels ...<-chan interface{}) <-chan interface{} {
	ctx, cancel := context.WithCancel(context.Background())
	// Создаем новый канал, который будет использоваться для объединения входных каналов
	myChan := make(chan interface{})
	for i, ch := range channels {
		ch := ch // для go 1.22 эта строчка не нужна)
		i := i
		go func() { // Для каждого канала создаем новую горутину
			select {
			case <-ch: // Ждем данных из текущего канала
				fmt.Printf("%d gorutine end work\n", i)
				close(myChan) // При получении данных закрываем общий канал
				cancel()
			case <-ctx.Done():
				fmt.Printf("%d gorutine close\n", i)
				break
			}
		}()
	}
	return myChan
}

func main() {
	sig := func(after time.Duration) <-chan interface{} {
		c := make(chan interface{})
		go func() {
			defer close(c)
			time.Sleep(after)
		}()

		return c
	}

	start := time.Now()
	<-or(
		sig(2*time.Hour),
		sig(5*time.Minute),
		sig(1*time.Second),
		sig(1*time.Hour),
		sig(1*time.Minute),
	)

	fmt.Printf("Done after %v", time.Since(start))
}
