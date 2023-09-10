package models

type Warning struct {
	ID     uint   `gorm:"primaryKey" json:"-"`
	Water  uint   `gorm:"not_null" json:"-"`
	Wind   uint   `gorm:"not_null" json:"-"`
	Status string `gorm:"not_null" json:"-"`
}
