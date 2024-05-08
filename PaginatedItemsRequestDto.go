package godtos

type PaginatedItemsRequestDto struct {
	// Page is a field of type uint64 that represents the number of the page.
	// It is a query parameter and can be used for pagination.
	// The minimum value for the page is 1.
	// The default value is 1.
	Page uint64 `query:"page" doc:"Number of the page" minimum:"1" default:"1"`
	// Size is a field of type uint64 that represents the number of elements on a page.
	// It is a query parameter used for pagination.
	// The default value is 10 and the minimum value is 1.
	// The maximum value is 1001.
	Size uint64 `query:"size" doc:"Elements count on the page" default:"10" minimum:"1" maximum:"1000"`
	// Sorts is a field of type string that represents a list of fields names with + or - at
	// the end separated by ',' for selecting the sort direction in a paginated request.
	Sorts string `query:"sorts" doc:"Sorts list with fields names with + or - at the end separated by ',' for select sort direction"`
	// Search is a field in the PaginatedRequestDTO struct that represents the search keyword.
	// It is used for searching elements based on a specific keyword.
	// The value of this field is provided as the "search" query parameter in a paginated request.
	Filters string `query:"filters" doc:"See filters documentation endpoint"`
}
