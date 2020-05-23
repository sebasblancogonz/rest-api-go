package models

import "time"

// User struct
type User struct {
	ID        string    `bson:"id"`
	Name      string    `bson:"name"`
	Surname   string    `bson:"surname"`
	Email     string    `bson:"email"`
	CreatedAt time.Time `bson:"created_at"`
	UpdatedAt time.Time `bson:"updated_at"`
}

// Users list
type Users []User
