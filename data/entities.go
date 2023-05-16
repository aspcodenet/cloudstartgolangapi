package data

import "time"

type Game struct {
	Id            int
	CreatedAt     time.Time
	Winner        string
	YourSelection string
	MySelection   string
}
