package models

import (
	"time"
)

// Namespace is a model struct for counter.
type Namespace struct {
	ID        uint      `gorm:"primarykey" json:"id"`
	Name      string    `gorm:"type:varchar(100);unique_index;not null;unique" json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// TableName defines table name for Namespace.
func (b *Namespace) TableName() string {
	return "namespaces"
}

// Counter is a model struct for counter.
type Counter struct {
	ID        uint      `gorm:"primarykey" json:"id"`
	Name      string    `gorm:"type:varchar(100);not null" json:"name"`
	Count     uint      `json:"count"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	// Relations
	NamespaceID uint      `json:"namespace_id"`
	Namespace   Namespace `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"namespace"`
}

// TableName defines table name for Counter.
func (b *Counter) TableName() string {
	return "counters"
}

// Count is a model struct for count.
type Count struct {
	ID        uint      `gorm:"primarykey" json:"id"`
	CreatedAt time.Time `json:"created_at"`

	// Relations
	CounterID uint    `json:"counter_id"`
	Counter   Counter `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"counter"`
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
