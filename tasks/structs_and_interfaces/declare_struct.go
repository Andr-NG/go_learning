package main

import "fmt"

type User struct {
	ID       int
	Email    string
	IsActive bool
}

func PrintUser(u User) {
	fmt.Printf("ID: %d, Email: %s, Active: %v\n", u.ID, u.Email, u.IsActive)
}

func ToggleActiveByValue(u User) {
    // A copy of the struct is made. Changes affect only the copy. Original remains unchanged
	
    u.IsActive = !u.IsActive
    fmt.Printf("After toggling by value: %v\n", u.IsActive)
}

func ToggleActiveByPointer(u *User) {
    // Function receives the memory address. Changes affect the original struct
    
	u.IsActive = !u.IsActive
    fmt.Printf("After toggling by pointer %v\n", u.IsActive)
}