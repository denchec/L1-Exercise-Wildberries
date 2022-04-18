package main

import "fmt"

type client struct{} // Структура нашего клиента, он принимает файлы только ExlInterface

func (c *client) ClientFunc(com ExlInterface) {
	fmt.Println("Clients files")
	com.ExlDataFunc()
}

type ExlInterface interface { // Интерфейс работает с файлами Exl и нашего клиента все устраивает
	ExlDataFunc()
}
type ExlStruct struct{}

func (m *ExlStruct) ExlDataFunc() {
	fmt.Println("Exl data files")
}

/* Этот интерфейс использует Json файлы, но ему нужно, чтобы клиент смог получить и прочитать их
Для этого мы и используем адаптер */

type JsonInterface interface {
	JsonDataFunc()
}
type JsonStruct struct{}

func (w *JsonStruct) JsonDataFunc() {
	fmt.Println("Json data files")
}

type Adapter struct {
	JsonStruct *JsonStruct
}

// Наш адаптер позволяет клиенту обращаться, через него, к методам JsonStruct при этом используя Exl интерфейс

func (w *Adapter) ExlDataFunc() {
	fmt.Println("Adapter converts ExlInterface to JsonInterface")
	w.JsonStruct.JsonDataFunc()
}

func main() {
	client := &client{}
	ExlStr := &ExlStruct{}
	fmt.Println("Client use ExlInterface")
	client.ClientFunc(ExlStr)

	JsonStr := &JsonStruct{}
	Adapter := &Adapter{JsonStruct: JsonStr}
	fmt.Println("Client use Adapted JsonInterface")
	client.ClientFunc(Adapter)
}
