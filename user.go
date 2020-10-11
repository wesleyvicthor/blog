package main

import (
	"bufio"
	"os"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Name     string
	Username string
	Password []byte
}

func (u *User) CheckPassw(passw string) error {
	return bcrypt.CompareHashAndPassword(u.Password, []byte(passw))
}

func Users() map[string]*User {
	f, _ := os.Open("data/Users.txt")

	scan := bufio.NewScanner(f)

	var users = make(map[string]*User)
	for scan.Scan() {
		cols := strings.Split(scan.Text(), "\t")
		passw, _ := bcrypt.GenerateFromPassword([]byte(cols[2]), bcrypt.DefaultCost)
		users[cols[1]] = &User{cols[0], cols[1], passw}
	}

	return users
}
