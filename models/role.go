package models

type Role int8

const (
	Public   Role = 1
	Internal Role = 2
	Private  Role = 3
)
