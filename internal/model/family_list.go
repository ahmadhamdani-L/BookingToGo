package model

import (
	"bookingtogo/internal/abstraction"
)

type FamilyListEntity struct {
	CustomerID *int    `json:"customer_id" validate:"required" example:"1"`
	FlRelation *string `json:"hubungan" validate:"required" example:"dani"`
	FlName     *string `json:"nama" validate:"required" example:"dani@gmail.com"`
	FlDob      *string `json:"tanggal_lahir" validate:"required" example:"dani@gmail.com"`
	
}

type FamilyListFilter struct {
}

type FamilyListEntityModel struct {
	ID int `json:"id" gorm:"primaryKey;autoIncrement;"`

	// entity
	FamilyListEntity

	// Customer CustomerEntityModel `json:"customer" gorm:"foreignKey:CustomerID"`

	// context
	Context *abstraction.Context `json:"-" gorm:"-"`
}

type FamilyListCustomEntity struct {
	FlRelation *string `json:"hubungan" validate:"required" example:"dani"`
	FlName     *string `json:"nama" validate:"required" example:"dani@gmail.com"`
	FlDob      *string `json:"tanggal_lahir" validate:"required" example:"dani@gmail.com"`
}
type FamilyListFilterModel struct {

	// filter
	FamilyListFilter
}

func (FamilyListEntityModel) TableName() string {
	return "family_list"
}
