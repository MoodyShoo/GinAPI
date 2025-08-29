package models

import "time"

type Gender uint

const (
	Male = iota
	Female
)

type User struct {
	Id           uint
	Login        string
	Name         string
	Gender       Gender
	Age          uint32
	Contacts     Contacts
	Avatar       []byte
	RegisterTime time.Time
	IsActive     bool
}
