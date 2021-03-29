package models

import (
	"time"
)

// Counter is a model struct for counter.
type Counter struct {
	ID        uint      `gorm:"primarykey" json:"id"`
	Name      string    `gorm:"type:varchar(100);unique_index;not null;unique" json:"name"`
	Count     uint      `json:"count"`
	Counts    []Count   `json:"-"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// TableName defines table name for Counter.
func (b *Counter) TableName() string {
	return "counters"
}

// Count is a model struct for count.
type Count struct {
	ID        uint      `gorm:"primarykey" json:"id"`
	CounterID uint      `json:"counter_id"`
	CreatedAt time.Time `json:"created_at"`
}

// TableName defines table name for Counter.
func (b *Count) TableName() string {
	return "counts"
}

// CountGroupResult is a grouped result of Count.
type CountGroupResult struct {
	Count int64     `json:"count"`
	Date  time.Time `json:"date"`
}
