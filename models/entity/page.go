package entity

type Page struct {
	Content  interface{} `json:"content"`
	Pageable Pageable    `json:"pageable"`
}
