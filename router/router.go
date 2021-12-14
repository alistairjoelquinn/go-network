package router

import (
	"github.com/alistairjoelquinn/go-network/handler"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	auth := app.Group("/auth")
	auth.Get("/user/id.json", handler.CheckUserStatus)
	auth.Post("/register", handler.CreateNewUser)
	auth.Post("/login", handler.LogUserIn)
	auth.Get("/logout", handler.LogUserOut)
	auth.Get("/password-reset/email-check", handler.CheckEmailForReset)
	auth.Get("/password-reset/verify-code", handler.VerifyAndResetUsersPassword)

	user := app.Group(("/user"))
	user.Get("/get-data", handler.GetUserData)
	user.Post("/upload", handler.UploaderUserImage)
	user.Get("/set-bio", handler.SetUserBio)
	user.Get("/recent-users", handler.GetRecentUsers)
	user.Get("/user-search/:q", handler.SearchForUsers)
	user.Get("/other-user/:id", handler.GetOtherUser)

	friendship := app.Group(("/friendship"))
	friendship.Get("/get-initial-status/:id", handler.GetInitialFrienshipStatus)
	friendship.Post("/add-friend/:id", handler.AddFriend)
	friendship.Post("/accept-friend/:id", handler.AcceptFriend)
	friendship.Post("/end-friendship/:id", handler.EndFriendship)
	friendship.Get("/friends-list", handler.GetFriendsList)
}
