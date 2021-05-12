package entities

type User struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginInput struct {
	Identity string `json:"identity"`
	Password string `json:"password"`
}
