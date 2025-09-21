package pagination

import "gorm.io/gorm"

type Pagination struct {
	Page int
	Size int
}

func New(page, size int) Pagination {
	if page < 1 {
		page = 1
	}
	if size <= 0 || size > 100 {
		size = 10
	}
	return Pagination{Page: page, Size: size}
}

func (p *Pagination) Scope() func(*gorm.DB) *gorm.DB {
	offset := (p.Page - 1) * p.Size
	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(offset).Limit(p.Size)
	}
}
