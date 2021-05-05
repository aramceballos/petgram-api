package posts

type Post struct {
	ID          int    `json:"id"`
	UserID      int    `json:"user_id"`
	CategoryID  int    `json:"category_id"`
	PostDate    string `json:"post_date"`
	ImageURL    string `json:"image_url"`
	Description string `json:"description"`
}
