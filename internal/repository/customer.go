package repository

import (
	"bookingtogo/internal/abstraction"
	"bookingtogo/internal/model"
	"net/http"
	"gorm.io/gorm"
)

type Customer interface {
	Find(ctx *abstraction.Context, w http.ResponseWriter, rp *http.Request, p *abstraction.Pagination) (*[]model.CustomerEntityModel, *abstraction.PaginationInfo, error)
    Create(ctx *abstraction.Context, w http.ResponseWriter, rp *http.Request, e *model.CustomerEntityModel) (*model.CustomerEntityModel, error)
    FindByID(ctx *abstraction.Context, w http.ResponseWriter, rp *http.Request, id *int) (*model.CustomerEntityModel, error)
}

type customer struct {
	abstraction.Repository
}

func NewCustomer(db *gorm.DB) *customer {
	return &customer{
		abstraction.Repository{
			Db: db,
		},
	}
}


func (r *customer) Find(ctx *abstraction.Context, w http.ResponseWriter, rp *http.Request, p *abstraction.Pagination) (*[]model.CustomerEntityModel, *abstraction.PaginationInfo, error) {
    conn := r.CheckTrx(ctx)

    var datas []model.CustomerEntityModel
    var info abstraction.PaginationInfo

    query := conn.Model(&model.CustomerEntityModel{})
   
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

    if err := query.Preload("FamilyList").Preload("Nationality").Find(&datas).Error; err != nil {
        return &datas, &info, err
    }

    info.Count = int(totalData)
    info.MoreRecords = false
    if int(totalData) > *p.PageSize {
        info.MoreRecords = true
    }

    // Mengembalikan nilai ctx yang telah ada, serta w, rp, datas, info, dan error
    return &datas, &info, nil
}
func (r *customer) Create(ctx *abstraction.Context, w http.ResponseWriter, rp *http.Request, e *model.CustomerEntityModel) (*model.CustomerEntityModel, error) {
	conn := r.CheckTrx(ctx)

	if err := conn.Create(e).Error; err != nil {
		return nil, err
	}
	if err := conn.Model(e).Preload("FamilyList").Preload("Nationality").First(e).Error; err != nil {
		return nil, err
	}

	return e, nil
}

func (r *customer) FindByID(ctx *abstraction.Context, w http.ResponseWriter, rp *http.Request, id *int) (*model.CustomerEntityModel, error) {
	conn := r.CheckTrx(ctx)

	var data model.CustomerEntityModel
	if err := conn.Where("id = ?", &id).Preload("TrialBalance").Preload("Company").Preload("CustomerDetail").Preload("UserCreated").Preload("UserModified").First(&data).Error; err != nil {
		return &data, err
	}
	
	return &data, nil
}

// func (r *customer) Create(ctx *abstraction.Context, e *model.CustomerEntityModel) (*model.CustomerEntityModel, error) {
// 	conn := r.CheckTrx(ctx)

// 	if err := conn.Create(e).WithContext(ctx.Request().Context()).Error; err != nil {
// 		return nil, err
// 	}
// 	if err := conn.Model(e).Preload("UserCreated").Preload("UserModified").First(e).WithContext(ctx.Request().Context()).Error; err != nil {
// 		return nil, err
// 	}
// 	e.UserCreatedString = e.UserCreated.Name
// 	e.UserModifiedString = &e.UserModified.Name
// 	return e, nil
// }

// func (r *customer) Update(ctx *abstraction.Context, id *int, e *model.CustomerEntityModel) (*model.CustomerEntityModel, error) {
// 	conn := r.CheckTrx(ctx)

// 	if err := conn.Model(e).Where("id = ?", &id).Updates(e).Preload("Company").WithContext(ctx.Request().Context()).Error; err != nil {
// 		return nil, err
// 	}
// 	if err := conn.Model(e).Where("id = ?", &id).Preload("Company").Preload("UserCreated").Preload("UserModified").First(e).WithContext(ctx.Request().Context()).Error; err != nil {
// 		return nil, err
// 	}
// 	e.UserCreatedString = e.UserCreated.Name
// 	e.UserModifiedString = &e.UserModified.Name
// 	return e, nil

// }
// func (r *customer) Delete(ctx *abstraction.Context, id *int, e *model.CustomerEntityModel) (*model.CustomerEntityModel, error) {
// 	conn := r.CheckTrx(ctx)

// 	if err := conn.Model(e).Where("id =?", id).Update("status", 4).WithContext(ctx.Request().Context()).Error; err != nil {
// 		return nil, err
// 	}
// 	e.UserCreatedString = e.UserCreated.Name
// 	e.UserModifiedString = &e.UserModified.Name
// 	return e, nil
// }