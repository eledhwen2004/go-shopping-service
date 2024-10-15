package api

import (
	"fmt"

	"shopping-clone/config"

	"github.com/gofiber/fiber/v2"
)

var (
	App        *fiber.App   = fiber.New()
	UserApi    fiber.Router = App.Group("/user")
	CommentApi fiber.Router = App.Group("/comment")
	OrderApi   fiber.Router = App.Group("/order")
	ProductApi fiber.Router = App.Group("/product")
	AuthApi    fiber.Router = App.Group("auth")
)

func ListenPort() {
	if err := App.Listen(config.GetListenPort()); err != nil {
		fmt.Printf("Error happened while listening the port: %v", err)
	}
}
