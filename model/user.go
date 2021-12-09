package model

type NewUser struct {
	First    string `json:"first"`
	Last     string `json:"last"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LogUserIn struct {
	ID             string `json:"id"`
	HashedPassword string `json:"hashed_password"`
}
