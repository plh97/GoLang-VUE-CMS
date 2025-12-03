package v1

import (
	"go-nunu/api"
	"go-nunu/internal/model"
)

type GetPermissionListRequest struct {
	api.PageRequest
	Name string `json:"name" form:"name"`
	ID   int    `json:"id" form:"id"`
}

type GetPermissionListResponseData struct {
	api.PageResponse
	List []model.Permission `json:"list"`
}
type GetPermissionListResponse struct {
	Response
	Data GetPermissionListResponseData `json:"data"`
}

type CreatePermissionRequest struct {
	Name   string `json:"name"`
	Key    string `json:"key"`
	Status int    `json:"status"`
	// Permissions []Permission `gorm:"many2many:sys_role_permissions;" json:"permissions"`
}
type CreatePermissionResponse struct {
	Response
	Data model.Permission `json:"data"`
}
