package factory

import (
	"bookingtogo/internal/database"
	"bookingtogo/internal/repository"
	"gorm.io/gorm"
)

type Factory struct {
	Db 	*gorm.DB
	CustomerRepository                             repository.Customer
	FamilyListRepository                           repository.FamilyList
	NationalityRepository                           repository.Nationality
	
}

func NewFactory() *Factory {
	f := &Factory{}
	f.SetupDb()
	f.SetupRepository()

	return f
}

func (f *Factory) SetupDb() {
	db, err := database.Connection("SAMPLE1")
	if err != nil {
		panic("Failed setup db, connection is undefined")
	}
	f.Db = db
}

func (f *Factory) SetupRepository() {
	if f.Db == nil {
		panic("Failed setup repository, db is undefined")
	}
	f.CustomerRepository = repository.NewCustomer(f.Db)
	f.FamilyListRepository = repository.NewFamilyList(f.Db)
	f.NationalityRepository = repository.NewNationality(f.Db)
	
}