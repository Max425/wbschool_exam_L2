package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"os/signal"
	"time"
)

/*
=== Утилита telnet ===

Реализовать примитивный telnet клиент:
Примеры вызовов:
go-telnet --timeout=10s host port go-telnet mysite.ru 8080 go-telnet --timeout=3s 1.1.1.1 123

Программа должна подключаться к указанному хосту (ip или доменное имя) и порту по протоколу TCP.
После подключения STDIN программы должен записываться в сокет, а данные полученные и сокета должны выводиться в STDOUT
Опционально в программу можно передать таймаут на подключение к серверу (через аргумент --timeout, по умолчанию 10s).

При нажатии Ctrl+D программа должна закрывать сокет и завершаться. Если сокет закрывается со стороны сервера, программа должна также завершаться.
При подключении к несуществующему сервер, программа должна завершаться через timeout.
*/

type Telnet struct {
	timeout    time.Duration
	host, port string
}

func NewTelnet() *Telnet {
	var t time.Duration
	flag.DurationVar(&t, "timeout", 10*time.Second, "timeout")
	flag.Parse()
	args := flag.Args()

	return &Telnet{
		timeout: t,
		host:    args[0],
		port:    args[1],
	}
}

func (t *Telnet) Run() error {
	endSignal := make(chan os.Signal)
	conn, err := net.DialTimeout("tcp", fmt.Sprintf("%s:%s", t.host, t.port), t.timeout)
	if err != nil {
		return err
	}
	defer conn.Close()
	signal.Notify(endSignal, os.Interrupt)

	go io.Copy(os.Stdout, conn)
	go io.Copy(conn, os.Stdin)

	select {
	case c := <-endSignal:
		log.Println("Catch signal:", c)
		return nil
	}
}

func main() {
	telnet := NewTelnet()
	err := telnet.Run()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
