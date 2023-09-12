package customer

import (
	"bookingtogo/internal/abstraction"
	"bookingtogo/internal/dto"
	"bookingtogo/internal/factory"
	"bookingtogo/internal/model"
	"bookingtogo/internal/repository"
	"bookingtogo/pkg/trxmanager"
	"net/http"

	"gorm.io/gorm"
)

type service struct {
	Repository                   repository.Customer
	FamilyListRepository   		 repository.FamilyList
	Db                           *gorm.DB
}

type Service interface {
	Find(w http.ResponseWriter, r *http.Request, payload *dto.CustomerGetRequest) (*dto.CustomerGetResponse, error)
}

func NewService(f *factory.Factory) *service {
	repository := f.CustomerRepository
	familyListRepository := f.FamilyListRepository

	db := f.Db
	return &service{
		Repository:                   repository,
		FamilyListRepository:   	  familyListRepository,
		Db: db,
	}
}

func (s *service) Find(w http.ResponseWriter, r *http.Request, payload *dto.CustomerGetRequest) (*dto.CustomerGetResponse, error) {
	data, info, err := s.Repository.Find(&abstraction.Context{ResponseWriter: w, Request: r}, w, r,  &payload.Pagination)
	if err != nil {
		return &dto.CustomerGetResponse{}, err
	}
	result := &dto.CustomerGetResponse{
		Datas:          *data,
		PaginationInfo: *info,
	}
	return result, nil
}

func (s *service) FindByID(w http.ResponseWriter, r *http.Request, payload *dto.CustomerGetByIDRequest) (*dto.CustomerGetByIDResponse, error) {
	data, err := s.Repository.FindByID(&abstraction.Context{ResponseWriter: w, Request: r}, w, r,  &payload.ID)
	if err != nil {
		return &dto.CustomerGetByIDResponse{}, err
	}
	result := &dto.CustomerGetByIDResponse{
		CustomerEntityModel:          *data,
	}
	return result, nil
}

func (s *service) FindCustom(w http.ResponseWriter, r *http.Request, payload *dto.CustomerGetRequest) (*dto.CustomerGetCustomResponse, error) {
	data, _, err := s.Repository.Find(&abstraction.Context{ResponseWriter: w, Request: r}, w, r,  &payload.Pagination)
	if err != nil {
		return &dto.CustomerGetCustomResponse{}, err
	}

	var custom model.CustomerCustomEntity
	for _, c := range *data {
		nationalityName := c.Nationality.NationalityName + " " + c.Nationality.Nationalitycode

		custom.Keluarga = c.FamilyList
		custom.CstName =  c.CstName
		custom.CstDob = c.CstDob
		custom.CstEmail = 	c.CstEmail
		custom.CstPhoneNum = c.CstPhoneNum
		custom.NationalityID = &nationalityName
	}
	result := &dto.CustomerGetCustomResponse{
		Datas:          custom,
	}
	return result, nil
}

func (s *service) Create(w http.ResponseWriter, r *http.Request, payload *dto.CustomerCreateRequest) (*dto.CustomerCreateResponse, error) {
	var cust model.CustomerEntityModel
	
	ctx := &abstraction.Context{
		ResponseWriter: w,
		Request:        r,
		Auth:           &abstraction.AuthContext{}, 
		Trx:            &abstraction.TrxContext{},  
	}
	var familyListArray []model.FamilyListEntityModel
	err := trxmanager.New(s.Db).WithTrx(ctx, func(ctx *abstraction.Context) error {
		cust.Context = ctx
		cust.CustomerEntity = payload.CustomerEntity
		// cust.CstDob = dateCust
		result, err := s.Repository.Create(&abstraction.Context{ResponseWriter: w, Request: r}, w, r, &cust)
		if err != nil {
			return err
		}
		cust = *result

		for _, fml := range payload.FamilyList {
			var fmly model.FamilyListEntityModel
			fmly.Context = ctx
			fmly.CustomerID = &result.ID
			fmly.FlName = fml.FlName
			fmly.FlRelation = fml.FlRelation
			fmly.FlDob = fml.FlDob

			family, err := s.FamilyListRepository.Create(&abstraction.Context{ResponseWriter: w, Request: r}, w, r, &fmly)
			if err != nil {
				return err
			}
			familyListArray = append(familyListArray,*family)
		}

		cust.FamilyList = familyListArray
		return nil
	})

	if err != nil {
		return &dto.CustomerCreateResponse{}, err
	}

	result := &dto.CustomerCreateResponse{
		CustomerEntityModel: cust,
	}
	return result, nil
}




