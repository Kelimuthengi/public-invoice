package models

import "gorm.io/gorm"

type UnitName string

const (
	UnitTypeBedSitter  UnitName = "BEDSITTER"
	UnitTypeOneBedRoom UnitName = "ONEBEDROOM"
	UnitTypeTwoBedRoom UnitName = "TWOBEDROOM"
)

type HouseUnitName struct {
	gorm.Model
	UnitTypeName string `gorm:"unique; not null" validate:"required,oneof=BEDSITTER ONEBEDROOM TWOBEDROOM"`
}

type HouseUnitTypes struct {
	gorm.Model
	Price           float64       `gorm:"not null" json:"price"`
	Discount        float64       `gorm:"default:0" json:"discount"`
	HouseNo         string        `gorm:"not null; unique;" validate:"required" json:"houseno"`
	HouseUnitNameID uint          `gorm:"not null" json:"house_unit_name_id"`
	HouseUnitName   HouseUnitName `gorm:"foreignKey:HouseUnitNameID; constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"houseId"`
}
