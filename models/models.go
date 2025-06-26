package models

import "gorm.io/gorm"


// campaign model represents a phishing campaign
type Campaign struct {
	gorm.Model
	Name        string `gorm:"not null"`
	Description string `gorm:"not null"`
	Target []Target
}


// Target users or victims of the phishing campaign
type Target struct {
	gorm.Model
	Name	 string `gorm:"not null"`
	Email	 string `gorm:"not null"`
	Phone	 string `gorm:"not null"`
	Clicked   bool   `gorm:"default:false"`
	Submitted bool   `gorm:"default:false"`
	CampaignID uint
}