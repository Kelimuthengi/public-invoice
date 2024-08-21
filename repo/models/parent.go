package models

import "gorm.io/gorm"

type Parent struct {
	gorm.Model
	Username        string         `gorm:"size:255;not null" json:"username" `
	Email           string         `gorm:"size:255;not null" json:"email"`
	Address         string         `gorm:"size:255;not null" json:"address"`
	Phonenumber     string         `gorm:"size:255;not null" json:"phonenumber"`
	UserID          uint           `gorm:"not null" json:"userid"`
	User            User           `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"user"`
	Students        []Student      `gorm:"foreignKey:ParentID" json:"students"`
	HouseUnitTypeID uint           `json:"house_id,omitempty"`
	HouseUnitTypes  HouseUnitTypes `gorm:"foreignKey:HouseUnitTypeID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"unittype"`
}
