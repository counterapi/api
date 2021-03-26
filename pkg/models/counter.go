package models

import (
	"time"
)

type Counter struct {
	ID        uint      `gorm:"primarykey" json:"id"`
	Name      string    `gorm:"type:varchar(100);unique_index;not null;unique" json:"name"`
	Count     int16     `json:"count"`
	Counts    []Count   `json:"-"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (b *Counter) TableName() string {
	return "counters"
}

type Count struct {
	ID        uint      `gorm:"primarykey" json:"id"`
	CounterID uint      `json:"counter_id"`
	CreatedAt time.Time `json:"created_at"`
}

func (b *Count) TableName() string {
	return "counts"
}
