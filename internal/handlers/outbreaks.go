package handlers

import (
	"case/internal/models"
	"database/sql"
	"log/slog"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

// HandlerOutbreakList handles the outbreak list page
func HandlerOutbreakList(c *fiber.Ctx, db *sql.DB, sl *slog.Logger, store *session.Store, config Config) error {
	// Get active outbreaks
	outbreaks, err := models.GetActiveOutbreaks(c.Context(), db)
	if err != nil {
		sl.Error("Failed to get outbreaks: " + err.Error())
		return c.Status(500).SendString("Failed to get outbreaks")
	}

	// Get default outbreak
	defaultOutbreak, err := models.GetDefaultOutbreak(c.Context(), db)
	if err != nil && err != sql.ErrNoRows {
		sl.Error("Failed to get default outbreak: " + err.Error())
		return c.Status(500).SendString("Failed to get default outbreak")
	}

	// If no default outbreak exists, create it
	if err == sql.ErrNoRows {
		defaultOutbreak = &models.Outbreak{
			Name:        sql.NullString{String: "Ebola 2025", Valid: true},
			Description: sql.NullString{String: "Ebola outbreak in 2025", Valid: true},
			StartDate:   sql.NullTime{Time: time.Now(), Valid: true},
			Status:      sql.NullString{String: "active", Valid: true},
			EnterOn:     sql.NullTime{Time: time.Now(), Valid: true},
		}
		if err := defaultOutbreak.Insert(c.Context(), db); err != nil {
			sl.Error("Failed to create default outbreak: " + err.Error())
			return c.Status(500).SendString("Failed to create default outbreak")
		}
	}

	// Convert outbreaks to interface slice
	items := make([]interface{}, len(outbreaks))
	for i, v := range outbreaks {
		items[i] = v
	}

	data := NewTemplateData(c, store)
	data.Items = items
	data.Form = defaultOutbreak

	return GenerateHTML(c, db, data, "outbreaks")
}

// HandlerOutbreakForm handles the outbreak form page
func HandlerOutbreakForm(c *fiber.Ctx, db *sql.DB, sl *slog.Logger, store *session.Store, config Config) error {
	var outbreak *models.Outbreak
	id := c.Params("i")
	if id != "0" {
		idInt, err := strconv.Atoi(id)
		if err != nil {
			sl.Error("Invalid outbreak ID: " + err.Error())
			return c.Status(400).SendString("Invalid outbreak ID")
		}
		outbreak, err = models.OutbreakByID(c.Context(), db, idInt)
		if err != nil {
			sl.Error("Failed to get outbreak: " + err.Error())
			return c.Status(500).SendString("Failed to get outbreak")
		}
	} else {
		outbreak = &models.Outbreak{}
	}

	data := NewTemplateData(c, store)
	data.Form = outbreak

	return GenerateHTML(c, db, data, "form_outbreak")
}

// HandlerOutbreakSubmit handles the outbreak form submission
func HandlerOutbreakSubmit(c *fiber.Ctx, db *sql.DB, sl *slog.Logger, store *session.Store, config Config) error {
	var outbreak models.Outbreak
	if err := DecodeFormData(c, &outbreak); err != nil {
		sl.Error("Failed to decode form data: " + err.Error())
		return c.Status(400).SendString("Invalid form data")
	}

	// Parse start date
	startDate, err := time.Parse("2006-01-02", c.FormValue("start_date"))
	if err != nil {
		sl.Error("Failed to parse start date: " + err.Error())
		return c.Status(400).SendString("Invalid start date format")
	}
	outbreak.StartDate = sql.NullTime{Time: startDate, Valid: true}

	// Set audit fields
	userID := GetCurrentUser(c, store)
	now := time.Now()
	if !outbreak.Exists() {
		outbreak.EnterOn = sql.NullTime{Time: now, Valid: true}
		outbreak.EnterBy = sql.NullInt64{Int64: int64(userID), Valid: true}
	}
	outbreak.EditOn = sql.NullTime{Time: now, Valid: true}
	outbreak.EditBy = sql.NullInt64{Int64: int64(userID), Valid: true}

	// Save outbreak
	if err := outbreak.Save(c.Context(), db); err != nil {
		sl.Error("Failed to save outbreak: " + err.Error())
		return c.Status(500).SendString("Failed to save outbreak")
	}

	return c.Redirect("/outbreaks")
}

// HandlerOutbreakClose handles closing an outbreak
func HandlerOutbreakClose(c *fiber.Ctx, db *sql.DB, sl *slog.Logger, store *session.Store, config Config) error {
	id := c.Params("i")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		sl.Error("Invalid outbreak ID: " + err.Error())
		return c.Status(400).SendString("Invalid outbreak ID")
	}

	outbreak, err := models.OutbreakByID(c.Context(), db, idInt)
	if err != nil {
		sl.Error("Failed to get outbreak: " + err.Error())
		return c.Status(500).SendString("Failed to get outbreak")
	}

	outbreak.Status = sql.NullString{String: "closed", Valid: true}
	outbreak.EndDate = sql.NullTime{Time: time.Now(), Valid: true}
	outbreak.EditOn = sql.NullTime{Time: time.Now(), Valid: true}
	outbreak.EditBy = sql.NullInt64{Int64: int64(GetCurrentUser(c, store)), Valid: true}

	if err := outbreak.Update(c.Context(), db); err != nil {
		sl.Error("Failed to close outbreak: " + err.Error())
		return c.Status(500).SendString("Failed to close outbreak")
	}

	return c.Redirect("/outbreaks")
}

// SetSelectedOutbreak sets the selected outbreak in the session
func SetSelectedOutbreak(c *fiber.Ctx, store *session.Store, outbreakID int) error {
	sess, err := store.Get(c)
	if err != nil {
		return err
	}
	sess.Set("selected_outbreak", outbreakID)
	sess.Set("outbreak_id", outbreakID)
	return sess.Save()
}

// GetSelectedOutbreak gets the selected outbreak from the session
func GetSelectedOutbreak(c *fiber.Ctx, store *session.Store) int {
	sess, err := store.Get(c)
	if err != nil {
		return 0
	}
	outbreakID := sess.Get("selected_outbreak")
	if outbreakID == nil {
		return 0
	}
	return outbreakID.(int)
}
