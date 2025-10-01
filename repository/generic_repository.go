package repository

import (
	"book-car/pkg/pagination"

	"gorm.io/gorm"
)

type Repo[T any] struct {
	db      *gorm.DB
	preload []string
}

func NewRepo[T any](db *gorm.DB, preload ...string) *Repo[T] {
	return &Repo[T]{db: db, preload: preload}
}

func (r *Repo[T]) withPreload(tx *gorm.DB) *gorm.DB {
	for _, p := range r.preload {
		tx = tx.Preload(p)
	}

	return tx
}

func (r *Repo[T]) Create(model *T) (*T, error) {
	result := r.db.Create(model)

	if result.Error != nil {
		return nil, result.Error
	}

	return model, nil
}

func (r *Repo[T]) FindOneByID(ID string) (*T, error) {
	var model *T
	result := r.db.Where("id = ?", ID).Find(model)

	if result.Error != nil {
		return nil, result.Error
	}

	return model, nil
}

func (r *Repo[T]) FindAll(pagination *pagination.Pagination) (*[]T, error) {
	var models *[]T
	result := r.db.Scopes(pagination.Scope()).Find(models)

	if result.Error != nil {
		return nil, result.Error
	}

	return models, nil
}

func (r *Repo[T]) Update(active *T, patch *T) (*T, error) {
	result := r.db.Model(active).Updates(patch)
	if result.Error != nil {
		return nil, result.Error
	}

	return active, nil
}

func (r *Repo[T]) Delete(model *T) error {
	result := r.db.Delete(model)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
