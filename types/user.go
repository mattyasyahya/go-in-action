package types

import "fmt"

// User defined type
type User struct {
	name       string
	email      string
	ext        string
	privileged bool
}

// Notify user email
func (user User) Notify() {
	fmt.Println("Send email to", user.email)
}

// ChangeEmail change user email
func (user *User) ChangeEmail(email string) {
	user.email = email
}
