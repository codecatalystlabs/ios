package main

import (
	"database/sql"
	"log/slog"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	_ "github.com/lib/pq"

	"case/internal/handlers"
	"case/internal/reports"
)

func SetRoute(app *fiber.App, db *sql.DB, store *session.Store, sl *slog.Logger, config handlers.Config) {
	RouteHome(app, db, sl, store, config)
	RouteVerify(app, db, sl, store, config)

	// Add outbreak routes
	RouteOutbreaks(app, db, sl, store, config)

	// discharge verification route
	app.Use(func(c *fiber.Ctx) error {
		return c.Next()
	})

	// Main application routes
	appGroup := app.Group("/")
	appGroup.Use(AuthRequired(store)) // Apply middleware for protected routes
	{
		// Home route
		appGroup.Get("/", func(c *fiber.Ctx) error { return handlers.HandlerHome(c, db, sl, store, config) })

		// Add more routes as needed...

		api := app.Group("/api") // Group for all API routes

		enk := api.Group("/encounter")
		sym := api.Group("/sym")
		mob := api.Group("/mob")
		rus := api.Group("/rush")
		lab := api.Group("/lab")
		sta := api.Group("/status")

		emp := app.Group("/employees") // Employees
		usr := app.Group("/users")     // users
		hfs := app.Group("/secure")    // Health facilities
		cse := app.Group("/cases")

		//enc := app.Group("/encounter")
		dis := app.Group("/discharge")
		rpt := app.Group("/reports")

		// Additional routes
		RouteFacilities(hfs, db, sl, store, config)
		RouteUsers(usr, db, sl, store, config)
		RouteCases(cse, db, sl, store, config)
		RouteMorbidity(mob, db, sl, store, config)
		RouteSymptoms(sym, db, sl, store, config)
		RouteRush(rus, db, sl, store, config)
		RouteLab(lab, db, sl, store, config)

		RouteEmployees(emp, db, sl, store, config)
		RouteDischarge(dis, db, sl, store, config)

		RouteReports(rpt, db, sl, store, config)

		RouteAPIEncounter(enk, db, sl, store, config)
		RouteAPIStatus(sta, db, sl, store, config)
	}
}

func AuthRequired(store *session.Store) fiber.Handler {
	return func(c *fiber.Ctx) error {
		if c.Path() == "/discharges/verify/:i" {
			return c.Next() // Skip authentication for this route
		}

		sess, err := store.Get(c)
		if err != nil {
			return err
		}
		userID := sess.Get("user")
		if userID == nil {
			return c.Redirect("/login", 302)
		}

		// Store user ID in Fiber Locals for later use
		c.Locals("userID", userID)

		return c.Next()
	}
}

func RouteAPIEncounter(v fiber.Router, db *sql.DB, sl *slog.Logger, store *session.Store, config handlers.Config) {
	v.Get("/", func(c *fiber.Ctx) error { return handlers.HandlerAPIGetEncounter(c, db, sl, store, config) })
}

func RouteAPIStatus(v fiber.Router, db *sql.DB, sl *slog.Logger, store *session.Store, config handlers.Config) {
	v.Get("/list", func(c *fiber.Ctx) error { return handlers.HandlerAPIGetStatuses(c, db, sl, store, config) })
	v.Post("/save", func(c *fiber.Ctx) error { return handlers.HandlerAPIPostStatus(c, db, sl, store, config) })
}

func RouteDischarge(v fiber.Router, db *sql.DB, sl *slog.Logger, store *session.Store, config handlers.Config) {
	v.Get("/list", func(c *fiber.Ctx) error { return handlers.GetDischarge(c, db, sl, store, config) })
	v.Get("/certificate", func(c *fiber.Ctx) error { return handlers.Certificate(c, db, sl, store, config) })
	v.Post("/save", func(c *fiber.Ctx) error { return handlers.Discharge(c, db, sl, store, config) })
}

func RouteHome(app *fiber.App, db *sql.DB, sl *slog.Logger, store *session.Store, config handlers.Config) {
	app.Get("/login", func(c *fiber.Ctx) error { return handlers.HandlerLoginForm(c, db, sl, store, config) })
	app.Post("/login", func(c *fiber.Ctx) error { return handlers.HandlerLoginSubmit(c, db, sl, store, config) })
	app.Get("/logout", func(c *fiber.Ctx) error { return handlers.HandlerLoginOut(c, sl, store, config) })
	app.Get("/forget", func(c *fiber.Ctx) error { return handlers.HandlerLoginForgot(c, db, sl, store, config) })
	app.Get("/help", func(c *fiber.Ctx) error { return handlers.HandlerHelp(c, db, sl, store, config) })
	app.Get("/vhf-cif", func(c *fiber.Ctx) error { return handlers.GenerateHTML(c, db, nil, "vhf_cif") })
	app.Post("/vhf-cif/save", func(c *fiber.Ctx) error { return handlers.HandlerVHFCIFSubmit(c, db, sl, store, config) })
}

func RouteVerify(app *fiber.App, db *sql.DB, sl *slog.Logger, store *session.Store, config handlers.Config) {
	app.Get("/verify/discharges/:i", func(c *fiber.Ctx) error { return handlers.VerifyDischarge2(c, db, sl, store, config) })
}

func RouteFacilities(v fiber.Router, db *sql.DB, sl *slog.Logger, store *session.Store, config handlers.Config) {
	v.Get("/new/:i", func(c *fiber.Ctx) error { return handlers.HandlerFacilityForm(c, db, sl, store, config) })
	v.Post("/save", func(c *fiber.Ctx) error { return handlers.HandlerFacilitySubmit(c, db, sl, store, config) })
	v.Post("/filter", func(c *fiber.Ctx) error { return handlers.HandlerFacilityList(c, db, sl, store, config) })
	v.Get("/list", func(c *fiber.Ctx) error { return handlers.HandlerFacilityList(c, db, sl, store, config) })
	v.Get("/", func(c *fiber.Ctx) error { return handlers.HandlerFacilityList(c, db, sl, store, config) })
}

func RouteUsers(v fiber.Router, db *sql.DB, sl *slog.Logger, store *session.Store, config handlers.Config) {
	v.Get("/new/:i", func(c *fiber.Ctx) error { return handlers.HandlerUserForm(c, db, sl, store, config) })
	v.Post("/save", func(c *fiber.Ctx) error { return handlers.HandlerUserSubmit(c, db, sl, store, config) })
	v.Post("/filter", func(c *fiber.Ctx) error { return handlers.HandlerUserList(c, db, sl, store, config) })
	v.Get("/list", func(c *fiber.Ctx) error { return handlers.HandlerUserList(c, db, sl, store, config) })
	v.Get("/", func(c *fiber.Ctx) error { return handlers.HandlerUserList(c, db, sl, store, config) })
}

func RouteEmployees(v fiber.Router, db *sql.DB, sl *slog.Logger, store *session.Store, config handlers.Config) {
	v.Get("/new/:i", func(c *fiber.Ctx) error { return handlers.HandlerEmployeeForm(c, db, sl, store, config) })
	v.Post("/save", func(c *fiber.Ctx) error { return handlers.HandlerEmployeeSubmit(c, db, sl, store, config) })
	v.Post("/filter", func(c *fiber.Ctx) error { return handlers.HandlerEmployeeList(c, db, sl, store, config) })
	v.Get("/list", func(c *fiber.Ctx) error { return handlers.HandlerEmployeeList(c, db, sl, store, config) })
	v.Get("/", func(c *fiber.Ctx) error { return handlers.HandlerEmployeeList(c, db, sl, store, config) })
}

func RouteCases(v fiber.Router, db *sql.DB, sl *slog.Logger, store *session.Store, config handlers.Config) {
	v.Get("/new/:i", func(c *fiber.Ctx) error { return handlers.HandlerCasesForm(c, db, sl, store, config) })
	v.Post("/save", func(c *fiber.Ctx) error { return handlers.HandlerCasesSubmit(c, db, sl, store, config) })
	v.Post("/filter", func(c *fiber.Ctx) error { return handlers.HandlerCasesList(c, db, sl, store, config) })
	v.Get("/list", func(c *fiber.Ctx) error { return handlers.HandlerCasesList(c, db, sl, store, config) })
	v.Get("/", func(c *fiber.Ctx) error { return handlers.HandlerCasesList(c, db, sl, store, config) })

	v.Get("/encounters/list/:i", func(c *fiber.Ctx) error { return handlers.HandlerCaseEncounterForm(c, db, sl, store, config) })
	v.Get("/encounters/new/:i", func(c *fiber.Ctx) error { return handlers.HandlerCaseEncounterForm(c, db, sl, store, config) })
	v.Get("/encounters/new/:i/:j", func(c *fiber.Ctx) error { return handlers.HandlerCaseEncounterForm(c, db, sl, store, config) })
	v.Post("/encounters/save", func(c *fiber.Ctx) error { return handlers.HandlerCaseEncounterSubmit(c, db, sl, store, config) })
}

func RouteCaseDischarge(v fiber.Router, db *sql.DB, sl *slog.Logger, store *session.Store, config handlers.Config) {
	v.Get("/view/:i/:j", func(c *fiber.Ctx) error { return handlers.HandlerCasesForm(c, db, sl, store, config) })
	v.Get("/new/:i/:j", func(c *fiber.Ctx) error { return handlers.HandlerCasesForm(c, db, sl, store, config) })
	v.Post("/save/:i/:j", func(c *fiber.Ctx) error { return handlers.HandlerCasesSubmit(c, db, sl, store, config) })
}

func RouteSymptoms(v fiber.Router, db *sql.DB, sl *slog.Logger, store *session.Store, config handlers.Config) {
	v.Get("/new/:i", func(c *fiber.Ctx) error { return handlers.HandlerSymptomsForm(c, db, sl, store, config) })
	v.Post("/save", func(c *fiber.Ctx) error { return handlers.HandlerSymptomsSubmit(c, db, sl, store, config) })
	v.Post("/filter", func(c *fiber.Ctx) error { return handlers.HandlerSymptomsList(c, db, sl, store, config) })
	v.Get("/list", func(c *fiber.Ctx) error { return handlers.HandlerSymptomsList(c, db, sl, store, config) })
	v.Get("/", func(c *fiber.Ctx) error { return handlers.HandlerSymptomsList(c, db, sl, store, config) })
}

func RouteMorbidity(v fiber.Router, db *sql.DB, sl *slog.Logger, store *session.Store, config handlers.Config) {
	v.Get("/new/:i", func(c *fiber.Ctx) error { return handlers.HandlerMorbidityForm(c, db, sl, store, config) })
	v.Post("/save", func(c *fiber.Ctx) error { return handlers.HandlerMorbiditySubmit(c, db, sl, store, config) })
	v.Post("/filter", func(c *fiber.Ctx) error { return handlers.HandlerMorbidityList(c, db, sl, store, config) })
	v.Get("/list", func(c *fiber.Ctx) error { return handlers.HandlerMorbidityList(c, db, sl, store, config) })
	v.Get("/", func(c *fiber.Ctx) error { return handlers.HandlerMorbidityList(c, db, sl, store, config) })
}

func RouteRush(v fiber.Router, db *sql.DB, sl *slog.Logger, store *session.Store, config handlers.Config) {
	v.Get("/new/:i", func(c *fiber.Ctx) error { return handlers.HandlerRushForm(c, db, sl, store, config) })
	v.Post("/save", func(c *fiber.Ctx) error { return handlers.HandlerRushSubmit(c, db, sl, store, config) })
	v.Post("/filter", func(c *fiber.Ctx) error { return handlers.HandlerRushList(c, db, sl, store, config) })
	v.Get("/list", func(c *fiber.Ctx) error { return handlers.HandlerRushList(c, db, sl, store, config) })
	v.Get("/", func(c *fiber.Ctx) error { return handlers.HandlerRushList(c, db, sl, store, config) })
}

func RouteLab(v fiber.Router, db *sql.DB, sl *slog.Logger, store *session.Store, config handlers.Config) {
	v.Get("/new/:i", func(c *fiber.Ctx) error { return handlers.HandlerLabForm(c, db, sl, store, config) })
	v.Post("/save", func(c *fiber.Ctx) error { return handlers.HandlerLabSubmit(c, db, sl, store, config) })
	v.Post("/filter", func(c *fiber.Ctx) error { return handlers.HandlerLabList(c, db, sl, store, config) })
	v.Get("/list", func(c *fiber.Ctx) error { return handlers.HandlerLabList(c, db, sl, store, config) })
	v.Get("/", func(c *fiber.Ctx) error { return handlers.HandlerLabList(c, db, sl, store, config) })
}

func RouteReports(v fiber.Router, db *sql.DB, sl *slog.Logger, store *session.Store, config handlers.Config) { //+
	//+
	v.Get("/view", func(c *fiber.Ctx) error { return reports.ReportView(c, db, sl, store, config) }) //+
	v.Get("/", func(c *fiber.Ctx) error { return reports.ReportHome(c, db, sl, store, config) })
}

// Add this new function for outbreak routes
func RouteOutbreaks(app *fiber.App, db *sql.DB, sl *slog.Logger, store *session.Store, config handlers.Config) {
	// Public routes
	app.Get("/outbreaks", func(c *fiber.Ctx) error { return handlers.HandlerOutbreakList(c, db, sl, store, config) })
	app.Get("/outbreaks/new/:i", func(c *fiber.Ctx) error { return handlers.HandlerOutbreakForm(c, db, sl, store, config) })
	app.Get("/outbreaks/edit/:i", func(c *fiber.Ctx) error { return handlers.HandlerOutbreakForm(c, db, sl, store, config) })
	app.Post("/outbreaks/save", func(c *fiber.Ctx) error { return handlers.HandlerOutbreakSubmit(c, db, sl, store, config) })
	app.Get("/outbreaks/close/:i", func(c *fiber.Ctx) error { return handlers.HandlerOutbreakClose(c, db, sl, store, config) })
	app.Post("/outbreaks/select/:i", func(c *fiber.Ctx) error {
		id, err := strconv.Atoi(c.Params("i"))
		if err != nil {
			return c.Status(400).SendString("Invalid outbreak ID")
		}
		if err := handlers.SetSelectedOutbreak(c, store, id); err != nil {
			return c.Status(500).SendString("Failed to select outbreak")
		}
		return c.SendStatus(200)
	})
}
