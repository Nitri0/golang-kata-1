package domain

import "time"

type Magazine struct {
	title       string
	isbn        string
	authors     []string
	publishedAt time.Time
}
