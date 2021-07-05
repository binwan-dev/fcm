package models

type ValidType int8

const (
	Realtime ValidType = 1
	NextRun  ValidType = 2
	Delay    ValidType = 3
)
