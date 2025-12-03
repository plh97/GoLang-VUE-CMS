package repository

import (
	"context"
	v1 "go-nunu/api/v1"
	"go-nunu/internal/model"

	"gorm.io/gorm"
)

type PermissionRepository interface {
	GetPermissionList(ctx context.Context, req v1.GetPermissionListRequest) ([]model.Permission, error)
	GetPermissionCount(ctx context.Context, req v1.GetPermissionListRequest) (int, error)
	CreatePermission(ctx context.Context, permission *model.Permission) (*model.Permission, error)
}

func NewPermissionRepository(
	repository *Repository,
) PermissionRepository {
	return &permissionRepository{
		Repository: repository,
	}
}

type permissionRepository struct {
	*Repository
}

func (r *permissionRepository) Get(ctx context.Context, param v1.GetPermissionListRequest) *gorm.DB {
	var permissions []model.Permission
	db := r.db.WithContext(ctx).Model(&permissions)
	if param.PageRequest.CurrentPage > 0 {
		db = db.Scopes(model.Paginate(param.PageRequest))
	}
	if param.Name != "" {
		db = db.Where("name LIKE ?", "%"+param.Name+"%")
	}
	if param.ID != 0 {
		db = db.Where("id = ?", param.ID)
	}
	return db
}

func (r *permissionRepository) GetPermissionList(ctx context.Context, req v1.GetPermissionListRequest) ([]model.Permission, error) {
	var permissionList []model.Permission
	db := r.Get(ctx, req)
	err := db.Find(&permissionList).Error
	if err != nil {
		return nil, err
	}
	return permissionList, nil
}

func (r *permissionRepository) GetPermissionCount(ctx context.Context, req v1.GetPermissionListRequest) (int, error) {
	var count int64
	db := r.Get(ctx, v1.GetPermissionListRequest{
		Name: req.Name,
		ID:   req.ID,
	})
	err := db.Count(&count).Error
	if err != nil {
		return 0, err
	}
	return int(count), nil
}

func (r *permissionRepository) CreatePermission(ctx context.Context, permission *model.Permission) (*model.Permission, error) {
	err := r.db.WithContext(ctx).Create(permission).Error
	return permission, err
}
