package model

type Book struct {
	ID               int    `json:"id"`
	Name             string `json:"name"`
	Author           string `json:"author"`
	Pages            int    `json:"pages"`
	PercentageOfRead int    `json:"percentageOfRead"`
}
