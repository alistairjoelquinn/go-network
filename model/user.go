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

type UpdatedBio struct {
	Bio string `json:"bio"`
}

type RecentUsers struct {
	ID    string  `json:"id"`
	First string  `json:"first"`
	Last  string  `json:"last"`
	Image *string `json:"image"`
}

type OtherUserID struct {
	ID string `json:"id"`
}

type OtherUser struct {
	UserId string  `json:"userId"`
	First  string  `json:"first"`
	Last   string  `json:"last"`
	Image  *string `json:"image"`
	Bio    *string `json:"bio"`
}

type RequestsFriends struct {
	ID           string  `json:"id"`
	First        string  `json:"first"`
	Last         string  `json:"last"`
	Image        *string `json:"image"`
	Accepted     bool    `json:"accepted"`
	FriendshipId string  `json:"friendship_id"`
}
