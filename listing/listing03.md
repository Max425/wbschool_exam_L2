Что выведет программа? Объяснить вывод программы. Объяснить внутреннее устройство интерфейсов и их отличие от пустых интерфейсов.

```go
package main

import (
	"fmt"
	"os"
)

func Foo() error {
	var err *os.PathError = nil
	return err
}

func main() {
	err := Foo()
	fmt.Println(err)
	fmt.Println(err == nil)
}
```

Ответ:
```
<nil>
false

fmt.Println(err == nil)
эта операция будет true тогда, и только тогда, когда оба поля в интерфейсе будут нил 
в интерфейсе 2 поля: tab и data : информация о конкретном типа и ссылка на данные (соответственно)

type iface struct {
    tab  *itab // Иформация об интерфейсе
    data unsafe.Pointer // Хранимые данные (информация о значении)
}
```
