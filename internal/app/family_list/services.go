package family_list

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
	Repository                   repository.FamilyList
	Db                           *gorm.DB
}

type Service interface {
	Find(w http.ResponseWriter, r *http.Request, payload *dto.FamilyListGetRequest) (*dto.FamilyListGetResponse, error)
}

func NewService(f *factory.Factory) *service {
	repository := f.FamilyListRepository

	db := f.Db
	return &service{
		Repository:                   repository,
		Db: db,
	}
}

func (s *service) Find(w http.ResponseWriter, r *http.Request, payload *dto.FamilyListGetRequest) (*dto.FamilyListGetResponse, error) {
	data, info, err := s.Repository.Find(&abstraction.Context{ResponseWriter: w, Request: r}, w, r,  &payload.Pagination)
	if err != nil {
		return &dto.FamilyListGetResponse{}, err
	}
	result := &dto.FamilyListGetResponse{
		Datas:          *data,
		PaginationInfo: *info,
	}
	return result, nil
}

func (s *service) Create(w http.ResponseWriter, r *http.Request, payload *dto.FamilyListCreateRequest) (*dto.FamilyListCreateResponse, error) {
	var fam model.FamilyListEntityModel
	
	ctx := &abstraction.Context{
		ResponseWriter: w,
		Request:        r,
		Auth:           &abstraction.AuthContext{}, 
		Trx:            &abstraction.TrxContext{},  
	}
	err := trxmanager.New(s.Db).WithTrx(ctx, func(ctx *abstraction.Context) error {
		fam.Context = ctx
		fam.FamilyListEntity = payload.FamilyListEntity
		// fam.CstDob = datefam
		result, err := s.Repository.Create(&abstraction.Context{ResponseWriter: w, Request: r}, w, r, &fam)
		if err != nil {
			return err
		}
		fam = *result
		return nil
	})

	if err != nil {
		return &dto.FamilyListCreateResponse{}, err
	}

	result := &dto.FamilyListCreateResponse{
		FamilyListEntityModel: fam,
	}
	return result, nil
}


