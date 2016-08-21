package types

import "fmt"

// Admin defined type
type Admin struct {
	user  User
	level string
}

// Duration alias type
type Duration int64

// CreateUser create user struct
func CreateUser() {
	var user User
	fmt.Println(user)

	eko := User{
		name:       "Eko Kurniawan",
		email:      "echo.khannedy@gmail.com",
		ext:        "123",
		privileged: true,
	}
	fmt.Println(eko)

	fred := Admin{
		user: User{
			name:       "Eko Kurniawan",
			email:      "echo.khannedy@gmail.com",
			ext:        "123",
			privileged: true,
		},
		level: "super",
	}
	fmt.Println(fred)
}

// Method invoke
func Method() {
	bill := User{"Bill", "bill@email.com", "123", true}
	bill.Notify()

	lisa := &User{"Lisa", "lisa@email.com", "123", true}
	lisa.Notify()

	bill.ChangeEmail("bill@newdomain.com")
	bill.Notify()

	lisa.ChangeEmail("lisa@newdomain.com")
	lisa.Notify()
}
