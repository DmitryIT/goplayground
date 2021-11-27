package main

import (
	"errors"
	"fmt"
)

type Role string

var (
	Unknown   Role = ""
	Guest     Role = "guest"
	Member    Role = "member"
	Moderator Role = "moderator"
	Admin     Role = "admin"
	Roles          = []Role{Unknown, Guest, Member, Moderator, Admin}
)

func (r Role) isValid() bool {
	for _, role := range Roles {
		if r == role {
			return true
		}
	}
	return false
}

func CreateUser(r Role) error {
	if !r.isValid() {
		return errors.New("unknown role ")
	}
	fmt.Println("created user with role", r)
	return nil
}
