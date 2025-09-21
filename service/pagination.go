package service

type paginate struct {
	limit int
	page  int
}

func Pagination(limit int, page int) *paginate {
	return &paginate{limit: limit, page: page}
}
