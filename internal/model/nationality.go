package model

import (
	"bookingtogo/internal/abstraction"
	// "bookingtogo/pkg/util/date"

	// "gorm.io/gorm"
)

type NationalityEntity struct {
	NationalityName string `json:"nationality_name"`
	Nationalitycode string `json:"nationality_code"`
}

type NationalityFilter struct {
	
}

type NationalityEntityModel struct {
	// abstraction
	abstraction.Entity

	// entity
	NationalityEntity

	// relations
	// TrialBalance     TrialBalanceEntityModel       `json:"trial_balance" gorm:"foreignKey:TbID"`
	// Company          CompanyEntityModel            `json:"company" gorm:"foreignKey:CompanyID"`
	// Customer []CustomerEntityModel `json:"customer" gorm:"foreignKey:NationalityID"`
	// UserRelationModel
	// context
	Context *abstraction.Context `json:"-" gorm:"-"`
}

type NationalityFilterModel struct {
	// abstraction
	abstraction.Filter

	// filter
	NationalityFilter
	// CompanyCustomFilter
}

func (NationalityEntityModel) TableName() string {
	return "nationality"
}

