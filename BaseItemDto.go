package godtos

import "time"

//BaseItemDto base item dto struct type
type BaseItemDto[IDType comparable] struct {
	//	ID - unique identifier
	ID IDType `doc:"id of this item"`
	//	CreatedAt - item create time
	CreatedAt time.Time `doc:"item creation time"`
}
