package models

import (
	"errors"
	"fmt"
)

// User contains all we need to know about users ;)
type User struct {
	ID        int
	FirstName string
	LastName  string
}

var (
	users  []*User
	nextID = 1
)

// GetUsers returns all the (pointers to the) users
func GetUsers() []*User {
	return users
}

// AddUser adds the (pointer of the) user to our collection
func AddUser(u User) (User, error) {
	if u.ID != 0 {
		return User{}, errors.New("new user must not include an id or it mues bei set to zero")
	}
	u.ID = nextID
	nextID++
	users = append(users, &u)
	return u, nil
}

// GetUserByID returns user or new object
func GetUserByID(id int) (User, error) {
	for _, u := range users {
		if u.ID == id {
			return *u, nil
		}
	}
	return User{}, fmt.Errorf("User with ID '%v' not found", id)
}

// UpdateUser updates existing user if one
func UpdateUser(u User) (User, error) {
	for i, candidate := range users {
		if candidate.ID == u.ID {
			users[i] = &u
			return u, nil
		}
	}
	return User{}, fmt.Errorf("User with ID '%v' not found", u.ID)
}

// RemoveUserByID removes user with given ID
func RemoveUserByID (id int) error {
	for i, u := range(users) {
		if (u.ID == id) {
			users = append(users[:i], users[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("User with ID '%v' not found", id)
}
