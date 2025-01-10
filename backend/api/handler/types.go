package handler

type GetSessionRes struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Picture    string `json:"picture"`
	Seed       string `json:"seed"`
	IsArchived string `json:"is_archived"`
}
