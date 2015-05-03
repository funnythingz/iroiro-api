package tables

import (
	"time"
)

type Iro struct {
	Id        int
	ColorName string
	ColorCode string
	Content   string
	ReIroId   int
	CreatedAt time.Time
	UpdatedAt time.Time
}
