package types

import "fmt"

// Notifier interface
type Notifier interface {
	Notify()
}

func sendNotification(notifier Notifier) {
	notifier.Notify()
}

// SayHello say hello interface
type SayHello interface {
	SayHello(name string)
}

func sayHello(value SayHello, name string) {
	value.SayHello(name)
}

// Animal struct
type Animal struct {
	name string
}

// SayHello from animal
func (animal *Animal) SayHello(name string) {
	fmt.Println("Hello ", name)
}

// SendNotification send notification
func SendNotification() {
	bill := User{"Bill", "bill@email.com", "123", true}
	sendNotification(bill)

	lisa := &User{"Lisa", "lisa@email.com", "123", true}
	sendNotification(lisa)

	bill.ChangeEmail("bill@newdomain.com")
	sendNotification(bill)

	lisa.ChangeEmail("lisa@newdomain.com")
	sendNotification(lisa)
}

// ImplementByPointer implement interface using pointer parameter
func ImplementByPointer() {
	kucing := Animal{"Kucing"}
	kucing.SayHello("Anjing")

	sayHello(&kucing, "Anjing")
}
