package dto

import (
	"bookingtogo/internal/abstraction"
	"bookingtogo/internal/model"
)

// Get
type FamilyListGetRequest struct {
	abstraction.Pagination
	model.FamilyListFilterModel
}
type FamilyListGetResponse struct {
	Datas          []model.FamilyListEntityModel
	PaginationInfo abstraction.PaginationInfo
}

// GetByID
type FamilyListGetByIDRequest struct {
	ID int `param:"id" validate:"required,numeric"`
}
type FamilyListGetByIDResponse struct {
	model.FamilyListEntityModel
}

// Create
type FamilyListCreateRequest struct {
	model.FamilyListEntity
}
type FamilyListCreateResponse struct {
	model.FamilyListEntityModel
}

// Update
type FamilyListUpdateRequest struct {
	ID int `param:"id" validate:"required,numeric"`
	model.FamilyListEntity
}
type FamilyListUpdateResponse struct {
	model.FamilyListEntityModel
}

// Delete
type FamilyListDeleteRequest struct {
	ID int `param:"id" validate:"required,numeric"`
}
type FamilyListDeleteResponse struct {
	model.FamilyListEntityModel
}
