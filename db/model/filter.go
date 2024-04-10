package model

type FilterQueryInterface interface {
	GetLimit() int64
	GetPage() int64
	GetPageSize() int64
	GetSorted() TypeSort
	GetSortedBy() string
}

type TypeSort int

const (
	NONE TypeSort = iota
	ASC
	DESC
)

type FilterQuery struct {
	Limit    int64
	Page     int64
	PageSize int64
	Sorted   TypeSort
	SortedBy string
}
