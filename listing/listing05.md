Что выведет программа? Объяснить вывод программы.

```go
package main

type customError struct {
	msg string
}

func (e *customError) Error() string {
	return e.msg
}

func test() *customError {
	{
		// do something
	}
	return nil
}

func main() {
	var err error
	err = test()
	if err != nil {
		println("error")
		return
	}
	println("ok")
}
```
Ответ:
```
error

error это интерфейс, аналогично как в 4, мы записываем в поле тип, но поле дата по прежнему nil
а тип: *main.customError 
```
