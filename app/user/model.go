package user

// User user model object
type User struct {
	Name     string `json:"name" xml:"name"`
	Username string `json:"username" xml:"username"`
	Email    string `json:"email" xml:"email"`
	Password string `json:"password" xml:"password"`
}
