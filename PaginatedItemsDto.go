package godtos

// PaginationInfoDto struct for paginated result response
type PaginationInfoDto struct {
	//Page - page number
	Page int `json:"page" doc:"Page number"`
	//Size - page size
	Size int `json:"size" doc:"Page size"`
	//TotalCount - total number of items
	TotalCount int `json:"totalCount" doc:"Total number of items"`
	//TotalPages - total number of pages
	TotalPages int `json:"totalPages" doc:"Total number of pages"`
}

// PaginatedItemsDto struct with slice of items and pagination info
type PaginatedItemsDto[T any] struct {
	//Pagination info about pagination
	Pagination PaginationInfoDto `json:"pagination" doc:"Pagination info"`
	//Items slice of items
	Items []T `json:"items" doc:"Slice of items"`
}

func pageCount(totalItems, pageSize int) int {
	pages := totalItems / pageSize

	if totalItems%pageSize > 0 {
		pages++
	}
	if pages < 1 {
		return 1
	}
	return pages
}

// NewEmptyPaginatedItemsDto constructor for PaginatedItemsDto[DataType] without data, used for errors
func NewEmptyPaginatedItemsDto[T any]() PaginatedItemsDto[T] {
	return PaginatedItemsDto[T]{
		Pagination: PaginationInfoDto{
			Page:       1,
			Size:       10,
			TotalCount: 0,
			TotalPages: 1,
		},
		Items: []T{},
	}
}

// NewPaginatedItemsDto constructor for PaginatedItemsDto[DataType]
func NewPaginatedItemsDto[T any](page, size, totalCount int, items []T) PaginatedItemsDto[T] {
	if items == nil {
		items = make([]T, 0)
	}
	return PaginatedItemsDto[T]{
		Pagination: PaginationInfoDto{
			Page:       page,
			Size:       size,
			TotalCount: totalCount,
			TotalPages: pageCount(totalCount, size),
		},
		Items: items,
	}
}
