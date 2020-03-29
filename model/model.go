package model

import (
	"time"
)

type User struct {
	Nickname string	`json:"nickname"`
	Phone string `json:"phone"`
	PhoneDevice string `json:"phoneDevice"`
	CreatedAt time.Time `json:"createdAt"`
}

type UserInput struct {
	Nickname string	`json:"nickname"`
	Phone string `json:"phone"`
	PhoneDevice string `json:"phoneDevice"`
}

type Users map[int]*User
