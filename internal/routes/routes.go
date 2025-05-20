package routes

import (
	"database/sql"
	"log/slog"

	"case/internal/handlers"
	"case/internal/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

// SetupRoutes configures all the routes for the application
func SetupRoutes(app *fiber.App, db *sql.DB, sl *slog.Logger, store *session.Store, config handlers.Config) {
	// Public routes
	app.Get("/", func(c *fiber.Ctx) error {
		return handlers.HandlerHome(c, db, sl, store, config)
	})
	app.Get("/login", func(c *fiber.Ctx) error {
		return handlers.HandlerLogin(c, db, sl, store, config)
	})
	app.Post("/login", func(c *fiber.Ctx) error {
		return handlers.HandlerLoginSubmit(c, db, sl, store, config)
	})

	// Protected routes
	api := app.Group("/", middleware.AuthRequired)

	// VHF CIF routes
	api.Get("/vhf-cif/patient", func(c *fiber.Ctx) error {
		return handlers.HandlerVHFPatient(c, db, sl, store, config)
	})
	api.Post("/vhf-cif/patient", func(c *fiber.Ctx) error {
		return handlers.HandlerVHFPatientSubmit(c, db, sl, store, config)
	})
	api.Get("/vhf-cif/clinical-signs/:id", func(c *fiber.Ctx) error {
		return handlers.HandlerVHFClinicalSigns(c, db, sl, store, config)
	})
	api.Post("/vhf-cif/clinical-signs/:id", func(c *fiber.Ctx) error {
		return handlers.HandlerVHFClinicalSignsSubmit(c, db, sl, store, config)
	})
	api.Get("/vhf-cif/hospitalization/:id", func(c *fiber.Ctx) error {
		return handlers.HandlerVHFHospitalization(c, db, sl, store, config)
	})
	api.Post("/vhf-cif/hospitalization/:id", func(c *fiber.Ctx) error {
		return handlers.HandlerVHFHospitalizationSubmit(c, db, sl, store, config)
	})
	api.Get("/vhf-cif/risk-factors/:id", func(c *fiber.Ctx) error {
		return handlers.HandlerVHFRiskFactors(c, db, sl, store, config)
	})
	api.Post("/vhf-cif/risk-factors/:id", func(c *fiber.Ctx) error {
		return handlers.HandlerVHFRiskFactorsSubmit(c, db, sl, store, config)
	})
	api.Get("/vhf-cif/laboratory/:id", func(c *fiber.Ctx) error {
		return handlers.HandlerVHFLaboratory(c, db, sl, store, config)
	})
	api.Post("/vhf-cif/laboratory/:id", func(c *fiber.Ctx) error {
		return handlers.HandlerVHFLaboratorySubmit(c, db, sl, store, config)
	})
	api.Get("/vhf-cif/investigator/:id", func(c *fiber.Ctx) error {
		return handlers.HandlerVHFInvestigator(c, db, sl, store, config)
	})
	api.Post("/vhf-cif/investigator/:id", func(c *fiber.Ctx) error {
		return handlers.HandlerVHFInvestigatorSubmit(c, db, sl, store, config)
	})
	api.Get("/vhf-cif/success", func(c *fiber.Ctx) error {
		return handlers.HandlerVHFSuccess(c, db, sl, store, config)
	})
	api.Get("/vhf-cif/list", func(c *fiber.Ctx) error {
		return handlers.HandlerVHFList(c, db, sl, store, config)
	})
	api.Get("/vhf-cif/view/:id", func(c *fiber.Ctx) error {
		return handlers.HandlerVHFView(c, db, sl, store, config)
	})

	// Other routes...
}
