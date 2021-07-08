package utils

type Paged struct {
	PageNumber int
	PageSize   int
	TotalPage  int
	TotalRow   int64
	List       interface{}
}
