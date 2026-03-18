package routes

import (
	"github.com/betterde/template/fiber/api/handler"
	"github.com/betterde/template/fiber/docs"
	"github.com/betterde/template/fiber/spa"
	swagger "github.com/gofiber/contrib/v3/swaggo"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/static"
)

func RegisterRoutes(app *fiber.App) {
	api := app.Group("/api")
	api.Get("/health", handler.HealthCheck).Name("api.health.check")

	// Swagger API specification file router
	app.Get("/swagger/*", static.New("", static.Config{
		FS:         docs.Serve(),
		Browse:     false,
		IndexNames: []string{"user.swagger.json"},
	})).Name("Swagger JSON Schema")

	// Swagger UI router
	app.Get("/docs/*", swagger.New(swagger.Config{
		URL:          "/swagger/user.swagger.json",
		DeepLinking:  false,
		DocExpansion: "none",
	})).Name("web.docs")

	// Embed SPA static resource
	app.Get("*", static.New("", static.Config{
		FS:         spa.Serve(),
		IndexNames: []string{"index.html"},
	})).Name("web.spa")
}
