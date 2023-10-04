package entity

type Pageable struct {
	Page   int `json:"page"`
	Size   int `json:"size"`
	OffSet int `json:"offset"`
}
