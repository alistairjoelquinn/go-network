package model

type NewUser struct {
	ID       string `json:"id"`
	First    string `json:"first"`
	Last     string `json:"last"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
