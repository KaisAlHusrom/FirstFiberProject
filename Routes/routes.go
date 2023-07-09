package Routes

import (
	Controllers "github.com/KaisAlHusrom/FirstFiberProject/Controllers"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Post("/users/addUser", Controllers.addUser)
}
