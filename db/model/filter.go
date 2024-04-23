package model

type FilterQueryInterface interface {
	GetLimit() int64
	GetPage() int64
	GetPageSize() int64
	GetSort() TypeSort
	GetSortBy() string
}

type TypeSort int

const (
	NONE TypeSort = iota
	ASC
	DESC
)

type FilterQuery struct {
	Limit    int64    `json:"limit"`
	Page     int64    `json:"page"`
	PageSize int64    `json:"pageSize"`
	Sorted   TypeSort `json:"sorted"`
	SortedBy string   `json:"sortedBy"`
}
