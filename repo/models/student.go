package models

import "gorm.io/gorm"

type Student struct {
	gorm.Model
	Username        string `gorm:"size 255; not null" json:"username"`
	AdmissionNumber string `gorm:"size: 255;not null;unique" json:"admissionnumber"`
	Stream          string `gorm:"size: 255;not null" json:"stream"`
	Boardingstatus  bool   `gorm:"not null" json:"boardingstatus"`
	Hostelname      string `gorm:"size:255;" json:"hostelname,omitempty"`
	ParentID        uint   `gorm:"not null" json:"parentid"`
	Parent          Parent `gorm:"foreignKey:ParentID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"parent"`
}
