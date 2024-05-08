package godtos

// PaginationInfoDto struct for paginated result response
type PaginationInfoDto struct {
	//Page - page number
	Page uint64 `json:"page" doc:"Page number" default:"1" minimum:"1"`
	//Size - page size
	Size uint64 `json:"size" doc:"Page size" minimum:"1" default:"10"`
	//TotalCount - total number of items
	TotalCount uint64 `json:"totalCount" minimum:"0" doc:"Total number of items"`
	//TotalPages - total number of pages
	TotalPages uint64 `json:"totalPages" minimum:"0" doc:"Total number of pages"`
}

// PaginatedItemsDto struct with slice of items and pagination info
type PaginatedItemsDto[T any] struct {
	//Pagination info about pagination
	Pagination PaginationInfoDto `json:"pagination" doc:"Pagination info"`
	//Items slice of items
	Items []T `json:"items" doc:"Slice of items"`
}

func pageCount(totalItems, pageSize uint64) uint64 {
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
func NewPaginatedItemsDto[T any](page, size, totalCount uint64, items []T) PaginatedItemsDto[T] {
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
