package models

import (
	"database/sql"
	"time"
)

// VHFCIF represents a VHF Case Investigation Form
type VHFCIF struct {
	ID             int64          `json:"id"`
	PatientName    string         `json:"patient_name"`
	Age            int            `json:"age"`
	Gender         string         `json:"gender"`
	District       string         `json:"district"`
	HealthFacility string         `json:"health_facility"`
	DateOfOnset    time.Time      `json:"date_of_onset"`
	Symptoms       string         `json:"symptoms"`
	TravelHistory  sql.NullString `json:"travel_history"`
	CreatedAt      time.Time      `json:"created_at"`
	UpdatedAt      time.Time      `json:"updated_at"`
}

// SaveVHFCIF saves a new VHF CIF record to the database
func SaveVHFCIF(db *sql.DB, cif *VHFCIF) error {
	query := `
		INSERT INTO vhf_cif (
			patient_name, age, gender, district, health_facility,
			date_of_onset, symptoms, travel_history, created_at, updated_at
		) VALUES (
			$1, $2, $3, $4, $5, $6, $7, $8, $9, $10
		) RETURNING id`

	now := time.Now()
	cif.CreatedAt = now
	cif.UpdatedAt = now

	err := db.QueryRow(
		query,
		cif.PatientName,
		cif.Age,
		cif.Gender,
		cif.District,
		cif.HealthFacility,
		cif.DateOfOnset,
		cif.Symptoms,
		cif.TravelHistory,
		cif.CreatedAt,
		cif.UpdatedAt,
	).Scan(&cif.ID)

	return err
}

// GetVHFCIF retrieves a VHF CIF record by ID
func GetVHFCIF(db *sql.DB, id int64) (*VHFCIF, error) {
	cif := &VHFCIF{}
	query := `
		SELECT id, patient_name, age, gender, district, health_facility,
			date_of_onset, symptoms, travel_history, created_at, updated_at
		FROM vhf_cif
		WHERE id = $1`

	err := db.QueryRow(query, id).Scan(
		&cif.ID,
		&cif.PatientName,
		&cif.Age,
		&cif.Gender,
		&cif.District,
		&cif.HealthFacility,
		&cif.DateOfOnset,
		&cif.Symptoms,
		&cif.TravelHistory,
		&cif.CreatedAt,
		&cif.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return cif, nil
}

// ListVHFCIF retrieves all VHF CIF records
func ListVHFCIF(db *sql.DB) ([]*VHFCIF, error) {
	query := `
		SELECT id, patient_name, age, gender, district, health_facility,
			date_of_onset, symptoms, travel_history, created_at, updated_at
		FROM vhf_cif
		ORDER BY created_at DESC`

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var cifs []*VHFCIF
	for rows.Next() {
		cif := &VHFCIF{}
		err := rows.Scan(
			&cif.ID,
			&cif.PatientName,
			&cif.Age,
			&cif.Gender,
			&cif.District,
			&cif.HealthFacility,
			&cif.DateOfOnset,
			&cif.Symptoms,
			&cif.TravelHistory,
			&cif.CreatedAt,
			&cif.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		cifs = append(cifs, cif)
	}

	return cifs, nil
}
