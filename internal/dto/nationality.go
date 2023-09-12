package dto

import (
	"bookingtogo/internal/abstraction"
	"bookingtogo/internal/model"
)

// Get
type NationalityGetRequest struct {
	abstraction.Pagination
	model.NationalityFilterModel
}
type NationalityGetResponse struct {
	Datas          []model.NationalityEntityModel
	PaginationInfo abstraction.PaginationInfo
}

// GetByID
type NationalityGetByIDRequest struct {
	ID int `param:"id" validate:"required,numeric"`
}
type NationalityGetByIDResponse struct {
	model.NationalityEntityModel
}

// Create
type NationalityCreateRequest struct {
	model.NationalityEntity
}
type NationalityCreateResponse struct {
	model.NationalityEntityModel
}

// Update
type NationalityUpdateRequest struct {
	ID int `param:"id" validate:"required,numeric"`
	model.NationalityEntity
}
type NationalityUpdateResponse struct {
	model.NationalityEntityModel
}

// Delete
type NationalityDeleteRequest struct {
	ID int `param:"id" validate:"required,numeric"`
}
type NationalityDeleteResponse struct {
	model.NationalityEntityModel
}
