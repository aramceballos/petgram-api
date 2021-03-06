package entities

type Post struct {
	ID          int     `json:"id"`
	UserID      int     `json:"-"`
	CategoryID  int     `json:"category_id"`
	PostDate    string  `json:"post_date"`
	ImageURL    string  `json:"image_url"`
	Description string  `json:"description"`
	Name        string  `json:"name"`
	Username    string  `json:"username"`
	Email       string  `json:"email"`
	Likes       []*Like `json:"likes"`
}

type Like struct {
	ID     int `json:"id"`
	UserID int `json:"user_id"`
	PostID int `json:"post_id"`
}
