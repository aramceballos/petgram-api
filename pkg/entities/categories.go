package entities

type Category struct {
	tableName struct{} `json:"-" pg:"petgram.categories"`
	ID        int      `json:"id"`
	Category  string   `json:"category"`
	ImageURL  string   `json:"image_url"`
}
