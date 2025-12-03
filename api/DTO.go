package api

type PageRequest struct {
	CurrentPage int `json:"current_page" form:"current_page"`
	PageSize    int `json:"page_size" form:"page_size"`
}

type PageResponse struct {
	Total int `json:"total"`
	// CurrentPage int `json:"current_page"`
}
