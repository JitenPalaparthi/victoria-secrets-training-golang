package models

type Contact struct {
	Id           int    `json:"id" gorm:"primaryKey"`
	Name         string `json:"name" gorm:"not null;index;text"`
	Email        string `json:"email"  gorm:"not null"`
	Mobile       string `json:"mobile"  gorm:"not null"`
	Status       string `json:"status"  gorm:"not null"`
	LastModified int64  `json:"lastModified" gorm:"column:lastModified;not null"`
}
