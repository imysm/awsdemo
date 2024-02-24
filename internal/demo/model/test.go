package model

import (
	"time"
)

type Test struct {
	ID            int       `gorm:"column:id;primaryKey"`
	Name          string    `gorm:"column:name;type:varchar(100);NOT NULL"`
	Message       string    `gorm:"column:message;type:varchar(300)"`
	CreatedAt     time.Time `gorm:"column:created_at;NOT NULL"`
	CreatedBy     string    `gorm:"column:created_by;type:varchar(100);default:'system';NOT NULL"`
	LastUpdatedAt time.Time `gorm:"column:last_updated_at;NOT NULL"`
	LastUpdatedBy string    `gorm:"column:last_updated_by;type:varchar(100);default:'system';NOT NULL"`
}

func (s *Test) TableName() string {
	return "demo_test_info"
}
