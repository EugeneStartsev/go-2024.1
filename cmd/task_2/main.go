package main

import (
	"time"
)

// Здесь надо реализовать структуру User
type User struct {
	Name       string
	Login      string
	Password   string
	LastSeenAt time.Time
}

// Здесь надо реализовать структуру Admin, включающую в себя User
type Admin struct {
	User
	Access uint8
}

// Здесь должна быть функция изменения имени пользователя
func renameUser(user interface{}, newName string) {
	switch val := user.(type) {
	case *User:
		val.Name = newName
	case *Admin:
		if val.Access < 3 {
			val.Name = newName
		}
	}
}

func main() {
}
