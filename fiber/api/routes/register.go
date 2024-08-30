package routes

import (
	"github.com/betterde/template/fiber/internal/response"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

func RegisterRoutes(app *fiber.App) {
	app.Get("/health", func(ctx *fiber.Ctx) error {
		return ctx.JSON(response.Success("Success", nil))
	}).Name("Health check")

	// Swagger API specification file router
	//app.Get("/swagger/*", filesystem.New(filesystem.Config{
	//	Root:               docs.Serve(),
	//	Index:              "user.swagger.json",
	//	NotFoundFile:       "user.swagger.json",
	//	ContentTypeCharset: "UTF-8",
	//})).Name("Swagger JSON Schema")

	// Swagger UI router
	app.Get("/docs/*", swagger.New(swagger.Config{
		URL:          "/swagger/user.swagger.json",
		DeepLinking:  false,
		DocExpansion: "none",
	})).Name("Swagger UI")

	// Embed SPA static resource
	//app.Get("*", filesystem.New(filesystem.Config{
	//	Root:               spa.Serve(),
	//	Index:              "index.html",
	//	NotFoundFile:       "index.html",
	//	ContentTypeCharset: "UTF-8",
	//})).Name("SPA static resource")
}
