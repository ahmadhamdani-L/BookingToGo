package repository

import (
	"bookingtogo/internal/abstraction"
	"bookingtogo/internal/model"
	"net/http"
	"gorm.io/gorm"
)

type Nationality interface {
	Find(ctx *abstraction.Context, w http.ResponseWriter, rp *http.Request, p *abstraction.Pagination) (*[]model.NationalityEntityModel, *abstraction.PaginationInfo, error)
    Create(ctx *abstraction.Context, w http.ResponseWriter, rp *http.Request, e *model.NationalityEntityModel) (*model.NationalityEntityModel, error)
}

type nationality struct {
	abstraction.Repository
}

func NewNationality(db *gorm.DB) *nationality {
	return &nationality{
		abstraction.Repository{
			Db: db,
		},
	}
}


func (r *nationality) Find(ctx *abstraction.Context, w http.ResponseWriter, rp *http.Request, p *abstraction.Pagination) (*[]model.NationalityEntityModel, *abstraction.PaginationInfo, error) {
    conn := r.CheckTrx(ctx)

    var datas []model.NationalityEntityModel
    var info abstraction.PaginationInfo

    query := conn.Model(&model.NationalityEntityModel{})
   
    // Pagination
    if p.Page == nil {
        page := 1
        p.Page = &page
    }
    if p.PageSize == nil {
        pageSize := 10
        p.PageSize = &pageSize
    }
    info = abstraction.PaginationInfo{
        Pagination: p,
    }
    limit := *p.PageSize
    offset := limit * (*p.Page - 1)
    var totalData int64
    query = query.Count(&totalData).Limit(limit).Offset(offset)

    if err := query.Find(&datas).Error; err != nil {
        return &datas, &info, err
    }

    info.Count = int(totalData)
    info.MoreRecords = false
    if int(totalData) > *p.PageSize {
        info.MoreRecords = true
    }

   
    return &datas, &info, nil
}
func (r *nationality) Create(ctx *abstraction.Context, w http.ResponseWriter, rp *http.Request, e *model.NationalityEntityModel) (*model.NationalityEntityModel, error) {
	conn := r.CheckTrx(ctx)

	if err := conn.Create(e).Error; err != nil {
		return nil, err
	}
	if err := conn.Model(e).First(e).Error; err != nil {
		return nil, err
	}

	return e, nil
}

