package routes

import (
	"github.com/betterde/template/fiber/api/handler"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

func RegisterRoutes(app *fiber.App) {
	api := app.Group("/api")
	api.Get("/health", handler.HealthCheck).Name("api.health.check")

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
	})).Name("web.docs")

	// Embed SPA static resource
	//app.Get("*", filesystem.New(filesystem.Config{
	//	Root:               spa.Serve(),
	//	Index:              "index.html",
	//	NotFoundFile:       "index.html",
	//	ContentTypeCharset: "UTF-8",
	//})).Name("web.spa")
}
