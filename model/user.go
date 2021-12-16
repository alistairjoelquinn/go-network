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

type User struct {
	ID    string  `json:"id"`
	First string  `json:"first"`
	Last  string  `json:"last"`
	Image *string `json:"image"`
	Bio   *string `json:"bio"`
}

type CheckEmail struct {
	Email string `json:"email"`
}

type CheckPasswordCode struct {
	Email           string `json:"email"`
	SecretCodeTyped string `json:"secretCodeTyped"`
	NewPassword     string `json:"newPassword"`
}
