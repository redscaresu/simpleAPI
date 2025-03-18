package models

type Info struct {
	Page       int    `json:"page"`
	PerPage    int    `json:"per_page"`
	Total      int    `json:"total"`
	TotalPages int    `json:"total_pages"`
	Data       []City `json:"data"`
}

type City struct {
	Name    string   `json:"name"`
	Status  []string `json:"status"`
	Weather string   `json:"weather"`
}
