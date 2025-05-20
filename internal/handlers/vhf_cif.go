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

// HandlerVHFCIFSubmit handles the submission of the VHF CIF form
func HandlerVHFCIFSubmit(c *fiber.Ctx, db *sql.DB, sl *slog.Logger, store *session.Store, config Config) error {
	// Create patient record
	patient := &models.VHFPatient{
		Surname:            c.FormValue("surname"),
		OtherNames:         c.FormValue("other_names"),
		Gender:             c.FormValue("gender"),
		PatientPhone:       c.FormValue("patient_phone"),
		PhoneOwner:         c.FormValue("phone_owner"),
		NextOfKin:          c.FormValue("next_of_kin"),
		NextOfKinPhone:     c.FormValue("next_of_kin_phone"),
		Status:             c.FormValue("status"),
		HeadOfHousehold:    c.FormValue("head_of_household"),
		VillageTown:        c.FormValue("village_town"),
		Parish:             c.FormValue("parish"),
		Subcounty:          c.FormValue("subcounty"),
		District:           c.FormValue("district"),
		CountryOfResidence: c.FormValue("country_of_residence"),
		Occupation:         c.FormValue("occupation"),
		IllVillageTown:     c.FormValue("ill_village_town"),
		IllSubcounty:       c.FormValue("ill_subcounty"),
		IllDistrict:        c.FormValue("ill_district"),
	}

	// Parse date fields
	if dob := c.FormValue("dob"); dob != "" {
		if t, err := time.Parse("2006-01-02", dob); err == nil {
			patient.DateOfBirth = sql.NullTime{Time: t, Valid: true}
		}
	}

	if dod := c.FormValue("date_of_death"); dod != "" {
		if t, err := time.Parse("2006-01-02", dod); err == nil {
			patient.DateOfDeath = sql.NullTime{Time: t, Valid: true}
		}
	}

	// Parse age fields
	if ageYears := c.FormValue("age_years"); ageYears != "" {
		if y, err := strconv.ParseInt(ageYears, 10, 32); err == nil {
			patient.AgeYears = sql.NullInt32{Int32: int32(y), Valid: true}
		}
	}

	if ageMonths := c.FormValue("age_months"); ageMonths != "" {
		if m, err := strconv.ParseInt(ageMonths, 10, 32); err == nil {
			patient.AgeMonths = sql.NullInt32{Int32: int32(m), Valid: true}
		}
	}

	// Parse location fields
	if lat := c.FormValue("latitude"); lat != "" {
		if l, err := strconv.ParseFloat(lat, 64); err == nil {
			patient.Latitude = sql.NullFloat64{Float64: l, Valid: true}
		}
	}

	if lng := c.FormValue("longitude"); lng != "" {
		if l, err := strconv.ParseFloat(lng, 64); err == nil {
			patient.Longitude = sql.NullFloat64{Float64: l, Valid: true}
		}
	}

	// Parse date range fields
	if from := c.FormValue("date_residing_from"); from != "" {
		if t, err := time.Parse("2006-01-02", from); err == nil {
			patient.DateResidingFrom = sql.NullTime{Time: t, Valid: true}
		}
	}

	if to := c.FormValue("date_residing_to"); to != "" {
		if t, err := time.Parse("2006-01-02", to); err == nil {
			patient.DateResidingTo = sql.NullTime{Time: t, Valid: true}
		}
	}

	// Save patient to database
	if err := models.SaveVHFPatient(db, patient); err != nil {
		sl.Error("Failed to save patient", "error", err)
		return c.Status(500).SendString("Failed to save patient information")
	}

	// Create clinical signs record
	signs := &models.VHFClinicalSigns{
		PatientID: patient.ID,
	}

	// Parse clinical signs fields
	if onset := c.FormValue("date_initial_onset"); onset != "" {
		if t, err := time.Parse("2006-01-02", onset); err == nil {
			signs.DateInitialOnset = sql.NullTime{Time: t, Valid: true}
		}
	}

	// Parse boolean fields for symptoms
	signs.Fever = sql.NullBool{Bool: c.FormValue("fever") == "Yes", Valid: true}
	signs.Vomiting = sql.NullBool{Bool: c.FormValue("vomiting") == "Yes", Valid: true}
	signs.Nausea = sql.NullBool{Bool: c.FormValue("nausea") == "Yes", Valid: true}
	signs.Diarrhea = sql.NullBool{Bool: c.FormValue("diarrhea") == "Yes", Valid: true}
	signs.IntenseFatigueGeneralWeakness = sql.NullBool{Bool: c.FormValue("intense_fatigue_general_weakness") == "Yes", Valid: true}
	signs.EpigastricPain = sql.NullBool{Bool: c.FormValue("epigastric_pain") == "Yes", Valid: true}
	signs.LowerAbdominalPain = sql.NullBool{Bool: c.FormValue("lower_abdominal_pain") == "Yes", Valid: true}
	signs.ChestPain = sql.NullBool{Bool: c.FormValue("chest_pain") == "Yes", Valid: true}
	signs.MusclePain = sql.NullBool{Bool: c.FormValue("muscle_pain") == "Yes", Valid: true}
	signs.JointPain = sql.NullBool{Bool: c.FormValue("joint_pain") == "Yes", Valid: true}
	signs.Headache = sql.NullBool{Bool: c.FormValue("headache") == "Yes", Valid: true}
	signs.Cough = sql.NullBool{Bool: c.FormValue("cough") == "Yes", Valid: true}
	signs.DifficultyBreathing = sql.NullBool{Bool: c.FormValue("difficulty_breathing") == "Yes", Valid: true}
	signs.DifficultySwallowing = sql.NullBool{Bool: c.FormValue("difficulty_swallowing") == "Yes", Valid: true}
	signs.SoreThroat = sql.NullBool{Bool: c.FormValue("sore_throat") == "Yes", Valid: true}
	signs.Jaundice = sql.NullBool{Bool: c.FormValue("jaundice") == "Yes", Valid: true}
	signs.Conjunctivitis = sql.NullBool{Bool: c.FormValue("conjunctivitis") == "Yes", Valid: true}
	signs.SkinRash = sql.NullBool{Bool: c.FormValue("skin_rash") == "Yes", Valid: true}
	signs.Hiccups = sql.NullBool{Bool: c.FormValue("hiccups") == "Yes", Valid: true}
	signs.PainBehindEyes = sql.NullBool{Bool: c.FormValue("pain_behind_eyes") == "Yes", Valid: true}
	signs.SensitiveToLight = sql.NullBool{Bool: c.FormValue("sensitive_to_light") == "Yes", Valid: true}
	signs.ComaUnconscious = sql.NullBool{Bool: c.FormValue("coma_unconscious") == "Yes", Valid: true}
	signs.ConfusedOrDisoriented = sql.NullBool{Bool: c.FormValue("confused_or_disoriented") == "Yes", Valid: true}
	signs.Convulsions = sql.NullBool{Bool: c.FormValue("convulsions") == "Yes", Valid: true}
	signs.UnexplainedBleeding = sql.NullBool{Bool: c.FormValue("unexplained_bleeding") == "Yes", Valid: true}
	signs.BleedingOfTheGums = sql.NullBool{Bool: c.FormValue("bleeding_of_the_gums") == "Yes", Valid: true}
	signs.BleedingFromInjectionSite = sql.NullBool{Bool: c.FormValue("bleeding_from_injection_site") == "Yes", Valid: true}
	signs.NoseBleedEpistaxis = sql.NullBool{Bool: c.FormValue("nose_bleed_epistaxis") == "Yes", Valid: true}
	signs.BloodyStool = sql.NullBool{Bool: c.FormValue("bloody_stool") == "Yes", Valid: true}
	signs.BloodInVomit = sql.NullBool{Bool: c.FormValue("blood_in_vomit") == "Yes", Valid: true}
	signs.CoughingUpBloodHemoptysis = sql.NullBool{Bool: c.FormValue("coughing_up_blood_hemoptysis") == "Yes", Valid: true}
	signs.BleedingFromVagina = sql.NullBool{Bool: c.FormValue("bleeding_from_vagina") == "Yes", Valid: true}
	signs.BruisingOfTheSkin = sql.NullBool{Bool: c.FormValue("bruising_of_the_skin") == "Yes", Valid: true}
	signs.BloodInUrine = sql.NullBool{Bool: c.FormValue("blood_in_urine") == "Yes", Valid: true}
	signs.OtherHemorrhagicSymptoms = sql.NullBool{Bool: c.FormValue("other_hemorrhagic_symptoms") == "Yes", Valid: true}

	// Save clinical signs to database
	if err := models.SaveVHFClinicalSigns(db, signs); err != nil {
		sl.Error("Failed to save clinical signs", "error", err)
		return c.Status(500).SendString("Failed to save clinical signs")
	}

	// Create hospitalization record
	hospitalization := &models.VHFHospitalization{
		PatientID:          patient.ID,
		Hospitalized:       c.FormValue("hospitalized") == "Yes",
		HealthFacilityName: c.FormValue("health_facility_name"),
		InIsolation:        c.FormValue("isolation") == "Yes",
	}

	// Parse date fields
	if admission := c.FormValue("admission_date"); admission != "" {
		if t, err := time.Parse("2006-01-02", admission); err == nil {
			hospitalization.AdmissionDate = sql.NullTime{Time: t, Valid: true}
		}
	}

	if isolation := c.FormValue("isolation_date"); isolation != "" {
		if t, err := time.Parse("2006-01-02", isolation); err == nil {
			hospitalization.IsolationDate = sql.NullTime{Time: t, Valid: true}
		}
	}

	// Save hospitalization to database
	if err := models.SaveVHFHospitalization(db, hospitalization); err != nil {
		sl.Error("Failed to save hospitalization", "error", err)
		return c.Status(500).SendString("Failed to save hospitalization information")
	}

	// Create risk factors record
	riskFactors := &models.VHFRiskFactors{
		PatientID:       patient.ID,
		ContactName:     c.FormValue("contact_name"),
		ContactRelation: c.FormValue("contact_relation"),
		ContactDates:    c.FormValue("contact_dates"),
		ContactVillage:  c.FormValue("contact_village"),
		ContactDistrict: c.FormValue("contact_district"),
		ContactStatus:   c.FormValue("contact_status"),
		ContactTypes:    c.FormValue("contact_types"),
	}

	// Parse boolean fields
	riskFactors.ContactWithCase = sql.NullBool{Bool: c.FormValue("contactWithCase") == "Yes", Valid: true}

	// Parse date fields
	if deathDate := c.FormValue("contact_death_date"); deathDate != "" {
		if t, err := time.Parse("2006-01-02", deathDate); err == nil {
			riskFactors.ContactDeathDate = sql.NullTime{Time: t, Valid: true}
		}
	}

	// Save risk factors to database
	if err := models.SaveVHFRiskFactors(db, riskFactors); err != nil {
		sl.Error("Failed to save risk factors", "error", err)
		return c.Status(500).SendString("Failed to save risk factors")
	}

	// Create laboratory record
	laboratory := &models.VHFLaboratory{
		PatientID:       patient.ID,
		SampleType:      c.FormValue("sample_type"),
		OtherSampleType: c.FormValue("other_sample_type"),
		RequestedTest:   c.FormValue("requested_test"),
		Serology:        c.FormValue("serology"),
		MalariaRDT:      c.FormValue("malaria_rdt"),
		HIVRDT:          c.FormValue("hiv_rdt"),
	}

	// Parse date and time fields
	if collectionDate := c.FormValue("sample_collection_date"); collectionDate != "" {
		if t, err := time.Parse("2006-01-02", collectionDate); err == nil {
			laboratory.SampleCollectionDate = sql.NullTime{Time: t, Valid: true}
		}
	}

	if collectionTime := c.FormValue("sample_collection_time"); collectionTime != "" {
		laboratory.SampleCollectionTime = sql.NullString{String: collectionTime, Valid: true}
	}

	// Save laboratory to database
	if err := models.SaveVHFLaboratory(db, laboratory); err != nil {
		sl.Error("Failed to save laboratory", "error", err)
		return c.Status(500).SendString("Failed to save laboratory information")
	}

	// Create investigator record
	investigator := &models.VHFInvestigator{
		PatientID:         patient.ID,
		InvestigatorName:  c.FormValue("investigator_name"),
		Phone:             c.FormValue("phone"),
		Email:             c.FormValue("email"),
		Position:          c.FormValue("position"),
		District:          c.FormValue("district"),
		HealthFacility:    c.FormValue("health_facility"),
		InformationSource: c.FormValue("information_source"),
		ProxyName:         c.FormValue("proxy_name"),
		ProxyRelation:     c.FormValue("proxy_relation"),
	}

	// Save investigator to database
	if err := models.SaveVHFInvestigator(db, investigator); err != nil {
		sl.Error("Failed to save investigator", "error", err)
		return c.Status(500).SendString("Failed to save investigator information")
	}

	// Redirect to success page
	return c.Redirect("/vhf-cif/success")
}
