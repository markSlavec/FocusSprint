package domain

import "time"

// TODO: потом поменять id int на uuid

type Matter struct {
	ID          int
	Name        string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	SubMatter   []*SubMatter
}

type SubMatter struct {
	ID          int
	Name        string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type MatterOfDay struct {
	ID        int
	Date      time.Time
	Matter    []*Matter
	CreatedAt time.Time
	UpdatedAt time.Time
}
