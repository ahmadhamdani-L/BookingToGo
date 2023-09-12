package model

import (
	"bookingtogo/internal/abstraction"
)

type CustomerEntity struct {
	NationalityID *int     `json:"kewarganegaraan" validate:"required" example:"1"`
	CstName       *string  `json:"nama" validate:"required" example:"dani"`
	CstEmail      *string  `json:"email" validate:"required" example:"dani@gmail.com"`
	CstDob        *string  `json:"tanggal_lahir" validate:"required" example:"dani@gmail.com"`
	CstPhoneNum   *float64 `json:"telepon" validate:"required" example:"08967677"`
}

type CustomerFilter struct {
	
}

type CustomerEntityModel struct {
	ID int `json:"id" gorm:"primaryKey;autoIncrement;"`

	// entity
	CustomerEntity

	Nationality NationalityEntityModel  `json:"nationality" gorm:"foreignKey:NationalityID"`
	FamilyList  []FamilyListEntityModel `json:"family_list" gorm:"foreignKey:CustomerID"`

	// context
	Context *abstraction.Context `json:"-" gorm:"-"`
}

type CustomerCustomEntity struct {
	Keluarga	  []FamilyListEntityModel `json:"keluarga" gorm:"-"`
	CstName       *string  `json:"nama" validate:"required" example:"dani"`
	CstDob        *string  `json:"tanggal_lahir" validate:"required" example:"dani@gmail.com"`
	CstPhoneNum   *float64 `json:"telepon" validate:"required" example:"08967677"`
	NationalityID *string  `json:"kewarganegaraan" validate:"required" example:"1"`
	CstEmail      *string  `json:"email" validate:"required" example:"dani@gmail.com"`
	
}

type CustomerFilterModel struct {
	// filter
	CustomerFilter
}

func (CustomerEntityModel) TableName() string {
	return "customer"
}
