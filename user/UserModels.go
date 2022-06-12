package user

type userCredentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

type userKeys struct {
	UserId  int    `json:"user_id"`
	UserKey string `json:"user_key"`
}
