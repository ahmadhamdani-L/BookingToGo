package dto

import (
	"bookingtogo/internal/abstraction"
	"bookingtogo/internal/model"
)

// Get
type CustomerGetRequest struct {
	abstraction.Pagination
	model.CustomerFilterModel
}
type CustomerGetResponse struct {
	Datas          []model.CustomerEntityModel
	PaginationInfo abstraction.PaginationInfo
}

// Get Custom
type CustomerGetCustomResponse struct {
	Datas          model.CustomerCustomEntity
}


// GetByID
type CustomerGetByIDRequest struct {
	ID int `param:"id" validate:"required,numeric"`
}
type CustomerGetByIDResponse struct {
	model.CustomerEntityModel
}


// Create
type CustomerCreateRequest struct {
	model.CustomerEntity
	FamilyList []model.FamilyListEntity
}
type CustomerCreateResponse struct {
	model.CustomerEntityModel
}


// Update
type CustomerUpdateRequest struct {
	ID int `param:"id" validate:"required,numeric"`
	model.CustomerEntity
}
type CustomerUpdateResponse struct {
	model.CustomerEntityModel
}


// Delete
type CustomerDeleteRequest struct {
	ID int `param:"id" validate:"required,numeric"`
}
type CustomerDeleteResponse struct {
	model.CustomerEntityModel
}
