package handlers

import (
	"case/internal/models"
	"case/internal/security"
	"database/sql"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

func HandlerCasesForm(c *fiber.Ctx, db *sql.DB, sl *slog.Logger, store *session.Store, config Config) error {
	DoZaLogging("INFO", "Starting Client form", nil)

	userID, userName := GetUser(c, sl, store)
	role := security.GetRoles(userID, "admin")
	id, err := strconv.Atoi(c.Params("i"))
	data := NewTemplateData(c, store)

	var client models.Client

	if err != nil || id < 1 {
		client.ID = 0
		data.IsIDPos = false
	} else {
		c, err := models.ClientByID(c.Context(), db, id)
		if err == nil {
			client = *c
		}

		data.IsIDPos = true
	}

	// Get outbreak ID from session
	sess, err := store.Get(c)
	if err != nil {
		return c.Status(400).SendString("Failed to get session")
	}
	outbreakID := sess.Get("outbreak_id")
	if outbreakID == nil {
		return c.Status(400).SendString("No outbreak selected")
	}

	// Set outbreak ID for new cases
	if client.ID == 0 {
		client.OutbreakID = sql.NullInt64{Int64: int64(outbreakID.(int)), Valid: true}
	}

	cE, err := models.ClientEncounterz(c.Context(), db, "client_id="+strconv.Itoa(id), outbreakID.(int))
	if err != nil {
		DoZaLogging("ERROR", "Failed to get encounters", err)
	}

	st, err := models.Statuses(c.Context(), db, "client_id="+strconv.Itoa(id))
	if err != nil {
		DoZaLogging("ERROR", "Failed to get statuses", err)
	}

	data.User = userName
	data.Role = role
	data.Optionz = Get_Client_Optionz()
	data.Form = client
	data.FormChild1 = cE
	data.FormChild2 = st

	DoZaLogging("INFO", "Load Client form", err)
	return GenerateHTML(c, db, data, "form_patients")
}

func HandlerCasesSubmit(c *fiber.Ctx, db *sql.DB, sl *slog.Logger, store *session.Store, config Config) error {

	id, er := strconv.Atoi(c.FormValue("id"))
	if er != nil {
		id = 0
	}

	client := models.Client{
		ID:               id,
		Firstname:        ParseNullString(c.FormValue("firstname")),
		Lastname:         ParseNullString(c.FormValue("lastname")),
		Othername:        ParseNullString(c.FormValue("othername")),
		Gender:           ParseNullInt(c.FormValue("gender")),
		DateOfBirth:      ParseNullString(c.FormValue("date_of_birth")),
		Age:              ParseNullFloat(c.FormValue("age")),
		Marital:          ParseNullInt(c.FormValue("marital")),
		Nin:              ParseNullString(c.FormValue("nin")),
		Nationality:      ParseNullInt(c.FormValue("nationality")),
		AdmDate:          ParseNullString(c.FormValue("adm_date")),
		AdmFrom:          ParseNullString(c.FormValue("adm_from")),
		LabNo:            ParseNullString(c.FormValue("lab_no")),
		CifNo:            ParseNullString(c.FormValue("cif_no")),
		EtuNo:            ParseNullString(c.FormValue("etu_no")),
		CaseNo:           ParseNullString(c.FormValue("case_no")),
		Occupation:       ParseNullInt(c.FormValue("occupation")),
		OccupationAza:    ParseNullString(c.FormValue("occupation_aza")),
		DateSymptomOnset: ParseNullString(c.FormValue("date_symptom_onset")),
		DateIsolation:    ParseNullString(c.FormValue("date_isolation")),
		Pregnant:         ParseNullInt(c.FormValue("pregnant")),
		AdmWard:          ParseNullString(c.FormValue("adm_ward")),
		Tb:               ParseNullInt(c.FormValue("tb")),
		Asplenia:         ParseNullInt(c.FormValue("asplenia")),
		Hep:              ParseNullInt(c.FormValue("hep")),
		Diabetes:         ParseNullInt(c.FormValue("diabetes")),
		Hiv:              ParseNullInt(c.FormValue("hiv")),
		Liver:            ParseNullInt(c.FormValue("liver")),
		Malignancy:       ParseNullInt(c.FormValue("malignancy")),
		Heart:            ParseNullInt(c.FormValue("heart")),
		Pulmonary:        ParseNullInt(c.FormValue("pulmonary")),
		Kidney:           ParseNullInt(c.FormValue("kidney")),
		Neurologic:       ParseNullInt(c.FormValue("neurologic")),
		Other:            ParseNullString(c.FormValue("other")),
		Transfer:         ParseNullInt(c.FormValue("transfer")),
		Site:             ParseNullInt(c.FormValue("site")),
		Status:           ParseNullString(c.FormValue("status")),

		//Status: ParseNullString(c.FormValue("status")),
	}

	//visID, _ := utilities.GetSequentialVisitID()
	userID := GetCurrentUser(c, store)

	client.EditOn.Valid = true
	client.EditBy.Valid = true

	client.EditBy.Int64 = int64(userID)
	client.EditOn.Time = time.Now()

	if client.ID == 0 {

		client.EnterOn.Valid = true
		client.EnterBy.Valid = true

		client.EnterBy.Int64 = int64(userID)
		client.EnterOn.Time = time.Now()

		client.UUID.Valid = true
		client.UUID.String = models.CreateUUID()

		//appID := models.CreateUUID()
		//client.UUID.String = appID

	}

	if client.ID == 0 {
		err := client.Insert(c.Context(), db)
		if err != nil {
			fmt.Println(err.Error())
		}
	} else {
		client.SetAsExists()
		err := client.Update(c.Context(), db)
		if err != nil {
			fmt.Println(err.Error())
		}
	}

	urlx := "/cases/new/" + strconv.Itoa(client.ID)

	return c.Redirect(urlx)
}

func HandlerCasesList(c *fiber.Ctx, db *sql.DB, sl *slog.Logger, store *session.Store, config Config) error {
	fmt.Println("starting case list")

	userID, userName := GetUser(c, sl, store)
	role := security.GetRoles(userID, "admin")

	data := NewTemplateData(c, store)
	data.User = userName
	data.Role = role

	fmt.Println("loading case list page")

	facility := GetCurrentFacility(c, db, sl, store)
	scope := GetDBInt("user_right", "function_scope", "", "user_id= "+strconv.Itoa(userID), 0)

	// Get outbreak ID from session
	sess, err := store.Get(c)
	if err != nil {
		return c.Status(400).SendString("Failed to get session")
	}
	outbreakID := sess.Get("outbreak_id")
	if outbreakID == nil {
		return c.Status(400).SendString("No outbreak selected")
	}

	filter := fmt.Sprintf("outbreak_id = %d", outbreakID.(int))
	if scope == 15 { // Full access to all facilities
		// Keep outbreak filter
	} else {
		if facility > 0 {
			filter += fmt.Sprintf(" AND site = %d", facility)
		}
	}

	clients, err := models.Clients(c.Context(), db, filter)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			fmt.Println("error loading case list: ", err.Error())
		} else {
			fmt.Println("error loading case list: ", err.Error())
		}
	}

	data.Form = clients

	return GenerateHTML(c, db, data, "list_patients")
}

func HandlerCaseEncounterForm(c *fiber.Ctx, db *sql.DB, sl *slog.Logger, store *session.Store, config Config) error {
	// Get client ID from query parameter
	clientIDStr := c.Query("client_id")
	if clientIDStr == "" {
		return c.Status(400).SendString("Client ID is required")
	}

	// Convert client ID to int
	clientID, err := strconv.Atoi(clientIDStr)
	if err != nil {
		return c.Status(400).SendString("Invalid client ID")
	}

	// Get outbreak ID from session
	sess, err := store.Get(c)
	if err != nil {
		return c.Status(400).SendString("Failed to get session")
	}
	outbreakID := sess.Get("outbreak_id")
	if outbreakID == nil {
		return c.Status(400).SendString("No outbreak selected")
	}

	// Get encounter date from query parameter
	encounterDate := c.Query("encounter_date")
	if encounterDate == "" {
		return c.Status(400).SendString("Encounter date is required")
	}

	// Get encounters for this client and date
	encounters, err := models.ClientEncounters(c.Context(), db, fmt.Sprintf("client_id = %d AND encounter_date = '%s'", clientID, encounterDate), outbreakID.(int))
	if err != nil {
		sl.Error("Failed to get encounters", "error", err)
		return c.Status(500).SendString("Failed to get encounters")
	}

	// Get client details
	client, err := models.ClientByID(c.Context(), db, clientID)
	if err != nil {
		sl.Error("Failed to get client", "error", err)
		return c.Status(500).SendString("Failed to get client")
	}

	data := fiber.Map{
		"Client":     client,
		"Encounters": encounters,
		"Date":       encounterDate,
	}

	return GenerateHTML(c, db, data, "form_case_encounter")
}

func HandlerCaseEncounterList(c *fiber.Ctx, db *sql.DB, sl *slog.Logger, store *session.Store, config Config) error {
	// Get client ID from query parameter
	clientIDStr := c.Query("client_id")
	if clientIDStr == "" {
		return c.Status(400).SendString("Client ID is required")
	}

	// Convert client ID to int
	clientID, err := strconv.Atoi(clientIDStr)
	if err != nil {
		return c.Status(400).SendString("Invalid client ID")
	}

	// Get outbreak ID from session
	sess, err := store.Get(c)
	if err != nil {
		return c.Status(400).SendString("Failed to get session")
	}
	outbreakID := sess.Get("outbreak_id")
	if outbreakID == nil {
		return c.Status(400).SendString("No outbreak selected")
	}

	// Get encounters for this client
	encounters, err := models.ClientEncounterz(c.Context(), db, fmt.Sprintf("client_id = %d", clientID), outbreakID.(int))
	if err != nil {
		sl.Error("Failed to get encounters", "error", err)
		return c.Status(500).SendString("Failed to get encounters")
	}

	// Get client details
	client, err := models.ClientByID(c.Context(), db, clientID)
	if err != nil {
		sl.Error("Failed to get client", "error", err)
		return c.Status(500).SendString("Failed to get client")
	}

	data := fiber.Map{
		"Client":     client,
		"Encounters": encounters,
	}

	return GenerateHTML(c, db, data, "list_case_encounter")
}

func saveEncounter(c *fiber.Ctx, db *sql.DB, userID int, cid, dte string) (int, int, int) {
	// Get selected outbreak from session
	store := session.New()
	sess, err := store.Get(c)
	if err != nil {
		return 0, 0, 0
	}
	outbreakID := sess.Get("selected_outbreak")
	if outbreakID == nil {
		return 0, 0, 0
	}

	var z int
	idx := []int{0, 1, 2}
	for k := 0; k < 3; k++ {
		z = k + 1
		id, er := strconv.Atoi(c.FormValue(fmt.Sprintf("id_%d", z)))
		if er != nil {
			id = 0
		}

		//encounter
		encounter := models.Encounter{
			EncounterID:   id,
			EncounterType: ParseNullInt(c.FormValue("encounter_type")),
			EncounterTime: ParseNullString(c.FormValue(fmt.Sprintf("encounter_time%d", z))),
			ClientID:      ParseNullInt(cid),
			EncounterDate: ParseNullString(dte),
			ManagedBy:     ParseNullInt(c.FormValue("managed_by")),
			ClinicalTeam:  ParseNullString(c.FormValue("clinical_team")),
			OutbreakID:    sql.NullInt64{Int64: int64(outbreakID.(int)), Valid: true},
		}

		if id == 0 {
			encounter.EnterOn.Valid = true
			encounter.EnterBy.Valid = true

			encounter.EnterBy.Int64 = int64(userID)
			encounter.EnterOn.Time = time.Now()
			err := encounter.Insert(c.Context(), db)
			if err != nil {
				fmt.Println(err.Error())
			}
		} else {
			encounter.SetAsExists()
			err := encounter.Update(c.Context(), db)
			if err != nil {
				fmt.Println(err.Error())
			}
		}
		idx[k] = encounter.EncounterID
	}

	return idx[0], idx[1], idx[2]
}

func saveVitals(c *fiber.Ctx, db *sql.DB, id1, id2, id3 int) {
	var z int
	z = 0

	idx := []int{0, 1, 2}

	idx[0] = id1
	idx[1] = id2
	idx[2] = id3

	for k := 0; k < 3; k++ {
		z = k + 1
		vitals_id, er := strconv.Atoi(c.FormValue(fmt.Sprintf("vitals_id_%d", z)))
		if er != nil {
			vitals_id = 0
		}

		vital := models.Vital{
			VitalsID:            vitals_id,
			EncounterID:         sql.NullInt64{Int64: int64(idx[k]), Valid: true},
			HeartRate:           ParseNullFloat(c.FormValue(fmt.Sprintf("heart_rate%d", z))),
			BpSystolic:          ParseNullFloat(c.FormValue(fmt.Sprintf("bp_systolic%d", z))),
			BpDiastolic:         ParseNullFloat(c.FormValue(fmt.Sprintf("bp_diastolic%d", z))),
			CapillaryRefill:     ParseNullInt(c.FormValue(fmt.Sprintf("capillary_refill%d", z))),
			RespiratoryRate:     ParseNullFloat(c.FormValue(fmt.Sprintf("respiratory_rate%d", z))),
			Saturation:          ParseNullFloat(c.FormValue(fmt.Sprintf("saturation%d", z))),
			Weight:              ParseNullFloat(c.FormValue(fmt.Sprintf("weight%d", z))),
			Height:              ParseNullFloat(c.FormValue(fmt.Sprintf("height%d", z))),
			Temperature:         ParseNullFloat(c.FormValue(fmt.Sprintf("temperature%d", z))),
			LowestConsciousness: ParseNullString(c.FormValue(fmt.Sprintf("lowest_consciousness%d", z))),
			MentalStatus:        ParseNullString(c.FormValue(fmt.Sprintf("mental_status%d", z))),
			Muac:                ParseNullFloat(c.FormValue(fmt.Sprintf("muac%d", z))),
			Bleeding:            ParseNullInt(c.FormValue(fmt.Sprintf("bleeding%d", z))),
			Shock:               ParseNullInt(c.FormValue(fmt.Sprintf("shock%d", z))),
			Meningitis:          ParseNullInt(c.FormValue(fmt.Sprintf("meningitis%d", z))),
			Confusion:           ParseNullInt(c.FormValue(fmt.Sprintf("confusion%d", z))),
			Seizure:             ParseNullInt(c.FormValue(fmt.Sprintf("seizure%d", z))),
			Coma:                ParseNullInt(c.FormValue(fmt.Sprintf("coma%d", z))),
			Bacteraemia:         ParseNullInt(c.FormValue(fmt.Sprintf("bacteraemia%d", z))),
			Hyperglycemia:       ParseNullInt(c.FormValue(fmt.Sprintf("hyperglycemia%d", z))),
			Hypoglycemia:        ParseNullInt(c.FormValue(fmt.Sprintf("hypoglycemia%d", z))),
			Other:               ParseNullString(c.FormValue(fmt.Sprintf("other%d", z))),
		}

		if vitals_id == 0 {
			err := vital.Insert(c.Context(), db)
			if err != nil {
				fmt.Println(err.Error())
			}
		} else {
			vital.SetAsExists()
			err := vital.Update(c.Context(), db)
			if err != nil {
				fmt.Println(err.Error())
			}
		}
	}

}

func getZaFormValue(c *fiber.Ctx, zname string, i int) string {
	return c.FormValue(fmt.Sprintf("%s%d", zname, i))
}

func saveClinical(c *fiber.Ctx, db *sql.DB, id1, id2, id3 int) {
	var z int

	idx := []int{0, 1, 2}

	idx[0] = id1
	idx[1] = id2
	idx[2] = id3

	for k := 0; k < 3; k++ {
		z = k + 1

		clinical_id, er := strconv.Atoi(c.FormValue(fmt.Sprintf("clinical_id_%d", z)))
		if er != nil {
			clinical_id = 0
		}

		clinical := models.Clinical{
			ClinicalID:  clinical_id,
			EncounterID: sql.NullInt64{Int64: int64(idx[k]), Valid: true},

			Fever:                ParseNullInt(getZaFormValue(c, "fever", z)),
			Fatigue:              ParseNullInt(getZaFormValue(c, "fatigue", z)),
			Weakness:             ParseNullInt(getZaFormValue(c, "weakness", z)),
			Malaise:              ParseNullInt(getZaFormValue(c, "malaise", z)),
			Myalgia:              ParseNullInt(getZaFormValue(c, "myalgia", z)),
			Anorexia:             ParseNullInt(getZaFormValue(c, "anorexia", z)),
			SoreThroat:           ParseNullInt(getZaFormValue(c, "sore_throat", z)),
			Headache:             ParseNullInt(getZaFormValue(c, "headache", z)),
			Nausea:               ParseNullInt(getZaFormValue(c, "nausea", z)),
			ChestPain:            ParseNullInt(getZaFormValue(c, "chest_pain", z)),
			JointPain:            ParseNullInt(getZaFormValue(c, "joint_pain", z)),
			Hiccups:              ParseNullInt(getZaFormValue(c, "hiccups", z)),
			Cough:                ParseNullInt(getZaFormValue(c, "cough", z)),
			DifficultyBreathing:  ParseNullInt(getZaFormValue(c, "difficulty_breathing", z)),
			DifficultySwallowing: ParseNullInt(getZaFormValue(c, "difficulty_swallowing", z)),
			AbdominalPain:        ParseNullInt(getZaFormValue(c, "abdominal_pain", z)),
			Diarrhoea:            ParseNullInt(getZaFormValue(c, "diarrhoea", z)),
			Vomiting:             ParseNullInt(getZaFormValue(c, "vomiting", z)),
			Irritability:         ParseNullInt(getZaFormValue(c, "irritability", z)),

			Dysphagia:              ParseNullInt(c.FormValue("dysphagia")),
			UnusualBleeding:        ParseNullInt(c.FormValue("unusual_bleeding")),
			Dehydration:            ParseNullInt(c.FormValue("dehydration")),
			Shock:                  ParseNullInt(c.FormValue("shock")),
			Anuria:                 ParseNullInt(c.FormValue("anuria")),
			Disorientation:         ParseNullInt(c.FormValue("disorientation")),
			Agitation:              ParseNullInt(c.FormValue("agitation")),
			Seizure:                ParseNullInt(c.FormValue("seizure")),
			Meningitis:             ParseNullInt(c.FormValue("meningitis")),
			Confusion:              ParseNullInt(c.FormValue("confusion")),
			Coma:                   ParseNullInt(c.FormValue("coma")),
			Bacteraemia:            ParseNullInt(c.FormValue("bacteraemia")),
			Hyperglycemia:          ParseNullInt(c.FormValue("hyperglycemia")),
			Hypoglycemia:           ParseNullInt(c.FormValue("hypoglycemia")),
			OtherComplications:     ParseNullInt(c.FormValue("other_complications")),
			AzaComplicationsSpecif: ParseNullString(c.FormValue("aza_complications_specif")),
			PharyngealErythema:     ParseNullInt(c.FormValue("pharyngeal_erythema")),
			PharyngealExudate:      ParseNullInt(c.FormValue("pharyngeal_exudate")),
			ConjunctivalInjection:  ParseNullInt(c.FormValue("conjunctival_injection")),
			OedemaFace:             ParseNullInt(c.FormValue("oedema_face")),
			TenderAbdomen:          ParseNullInt(c.FormValue("tender_abdomen")),
			SunkenEyes:             ParseNullInt(c.FormValue("sunken_eyes")),
			TentingSkin:            ParseNullInt(c.FormValue("tenting_skin")),
			PalpableLiver:          ParseNullInt(c.FormValue("palpable_liver")),
			PalpableSpleen:         ParseNullInt(c.FormValue("palpable_spleen")),
			Jaundice:               ParseNullInt(c.FormValue("jaundice")),
			EnlargedLymphNodes:     ParseNullInt(c.FormValue("enlarged_lymph_nodes")),
			LowerExtremityOedema:   ParseNullInt(c.FormValue("lower_extremity_oedema")),
			Bleeding:               ParseNullInt(c.FormValue("clinical_bleeding")),
			BleedingNose:           ParseNullInt(c.FormValue("bleeding_nose")),
			BleedingMouth:          ParseNullInt(c.FormValue("bleeding_mouth")),
			BleedingVagina:         ParseNullInt(c.FormValue("bleeding_vagina")),
			BleedingRectum:         ParseNullInt(c.FormValue("bleeding_rectum")),
			BleedingSputum:         ParseNullInt(c.FormValue("bleeding_sputum")),
			BleedingUrine:          ParseNullInt(c.FormValue("bleeding_urine")),
			BleedingIvSite:         ParseNullInt(c.FormValue("bleeding_iv_site")),
			BleedingOther:          ParseNullInt(c.FormValue("bleeding_other")),
			BleedingOtherSpecif:    ParseNullString(c.FormValue("bleeding_other_specif")),
		}

		if clinical_id == 0 {
			err := clinical.Insert(c.Context(), db)
			if err != nil {
				fmt.Println(err.Error())
			}
		} else {
			clinical.SetAsExists()
			err := clinical.Update(c.Context(), db)
			if err != nil {
				fmt.Println(err.Error())
			}
		}

	}

}

func HandlerCaseEncounterSubmit(c *fiber.Ctx, db *sql.DB, sl *slog.Logger, store *session.Store, config Config) error {
	userID := GetCurrentUser(c, store)
	cid := c.FormValue("cid")
	dte := c.FormValue("encounter_date")

	id1, id2, id3 := saveEncounter(c, db, userID, cid, dte)

	//vital
	saveVitals(c, db, id1, id2, id3)

	//clinical

	saveClinical(c, db, id1, id2, id3)

	//lab
	lab_id, er := strconv.Atoi(c.FormValue("lab_id"))
	if er != nil {
		lab_id = 0
	}

	lab := models.Lab{
		LabID:                 lab_id,
		EncounterID:           sql.NullInt64{Int64: int64(id1), Valid: true},
		Specimen:              ParseNullInt(c.FormValue("specimen")),
		SampleBlood:           ParseNullInt(c.FormValue("sample_blood")),
		SampleUrine:           ParseNullInt(c.FormValue("sample_urine")),
		SampleSwab:            ParseNullInt(c.FormValue("sample_swab")),
		SampleAza:             ParseNullString(c.FormValue("sample_aza")),
		EbolaRdt:              ParseNullInt(c.FormValue("ebola_rdt")),
		EbolaRdtDate:          ParseNullString(c.FormValue("ebola_rdt_date")),
		EbolaRdtResults:       ParseNullInt(c.FormValue("ebola_rdt_results")),
		EbolaPcr:              ParseNullInt(c.FormValue("ebola_pcr")),
		EbolaPcrAza:           ParseNullString(c.FormValue("ebola_pcr_aza")),
		EbolaPcrDate:          ParseNullString(c.FormValue("ebola_pcr_date")),
		EbolaPcrGp:            ParseNullInt(c.FormValue("ebola_pcr_gp")),
		EbolaPcrGpCt:          ParseNullFloat(c.FormValue("ebola_pcr_gp_ct")),
		EbolaPcrNp:            ParseNullInt(c.FormValue("ebola_pcr_np")),
		EbolaPcrNpCt:          ParseNullFloat(c.FormValue("ebola_pcr_np_ct")),
		EbolaPcrIndeterminate: ParseNullInt(c.FormValue("ebola_pcr_indeterminate")),
		MalariaRdtDate:        ParseNullString(c.FormValue("malaria_rdt_date")),
		MalariaRdtResult:      ParseNullInt(c.FormValue("malaria_rdt_result")),
		BloodCultureDate:      ParseNullString(c.FormValue("blood_culture_date")),
		BloodCultureResult:    ParseNullInt(c.FormValue("blood_culture_result")),
		TestPosInfection:      ParseNullInt(c.FormValue("test_pos_infection")),
		TestPosInfectionAza:   ParseNullString(c.FormValue("test_pos_infection_aza")),
		Haemoglobinuria:       ParseNullInt(c.FormValue("haemoglobinuria")),
		Proteinuria:           ParseNullInt(c.FormValue("proteinuria")),
		Hematuria:             ParseNullInt(c.FormValue("hematuria")),
		BloodGas:              ParseNullInt(c.FormValue("blood_gas")),
		Ph:                    ParseNullFloat(c.FormValue("ph")),
		Pco2:                  ParseNullFloat(c.FormValue("pco2")),
		Pao2:                  ParseNullFloat(c.FormValue("pao2")),
		Hco3:                  ParseNullFloat(c.FormValue("hco3")),
		OxygenTherapy:         ParseNullFloat(c.FormValue("oxygen_therapy")),
		AltSgpt:               ParseNullFloat(c.FormValue("alt_sgpt")),
		AstSgo:                ParseNullFloat(c.FormValue("ast_sgo")),
		Creatinine:            ParseNullFloat(c.FormValue("creatinine")),
		Potassium:             ParseNullFloat(c.FormValue("potassium")),
		Urea:                  ParseNullFloat(c.FormValue("urea")),
		CreatinineKinase:      ParseNullFloat(c.FormValue("creatinine_kinase")),
		Calcium:               ParseNullFloat(c.FormValue("calcium")),
		Sodium:                ParseNullFloat(c.FormValue("sodium")),
		AltSgptNd:             ParseNullInt(c.FormValue("alt_sgpt_nd")),
		AstSgoNd:              ParseNullInt(c.FormValue("ast_sgo_nd")),
		CreatinineNd:          ParseNullInt(c.FormValue("creatinine_nd")),
		PotassiumNd:           ParseNullInt(c.FormValue("potassium_nd")),
		UreaNd:                ParseNullInt(c.FormValue("urea_nd")),
		CreatinineKinaseNd:    ParseNullInt(c.FormValue("creatinine_kinase_nd")),
		CalciumNd:             ParseNullInt(c.FormValue("calcium_nd")),
		SodiumNd:              ParseNullInt(c.FormValue("sodium_nd")),
		Glucose:               ParseNullFloat(c.FormValue("glucose")),
		Lactate:               ParseNullFloat(c.FormValue("lactate")),
		Haemoglobin:           ParseNullFloat(c.FormValue("haemoglobin")),
		TotalBilirubin:        ParseNullFloat(c.FormValue("total_bilirubin")),
		WbcCount:              ParseNullFloat(c.FormValue("wbc_count")),
		Platelets:             ParseNullFloat(c.FormValue("platelets")),
		Pt:                    ParseNullFloat(c.FormValue("pt")),
		Aptt:                  ParseNullFloat(c.FormValue("aptt")),
		GlucoseNd:             ParseNullInt(c.FormValue("glucose_nd")),
		LactateNd:             ParseNullInt(c.FormValue("lactate_nd")),
		HaemoglobinNd:         ParseNullInt(c.FormValue("haemoglobin_nd")),
		TotalBilirubinNd:      ParseNullInt(c.FormValue("total_bilirubin_nd")),
		WbcCountNd:            ParseNullInt(c.FormValue("wbc_count_nd")),
		PlateletsNd:           ParseNullInt(c.FormValue("platelets_nd")),
		PtNd:                  ParseNullInt(c.FormValue("pt_nd")),
		ApttNd:                ParseNullInt(c.FormValue("aptt_nd")),
	}

	fmt.Println(lab)

	if lab_id == 0 {
		err := lab.Insert(c.Context(), db)
		if err != nil {
			fmt.Println(err.Error())
		}
	} else {
		lab.SetAsExists()
		err := lab.Update(c.Context(), db)
		if err != nil {
			fmt.Println(err.Error())
		}
	}

	//treat

	treat_id, er := strconv.Atoi(c.FormValue("treat_id"))
	if er != nil {
		treat_id = 0
	}

	treat := &models.Treatment{
		TreatmentID:            treat_id,
		EncounterID:            sql.NullInt64{Int64: int64(id1), Valid: true},
		Antibacterial:          ParseNullInt2(c.FormValue("antibacterial")),
		Amoxicillin:            ParseNullInt2(c.FormValue("amoxicillin")),
		Ceftriaxone:            ParseNullInt2(c.FormValue("ceftriaxone")),
		Cefixime:               ParseNullInt2(c.FormValue("cefixime")),
		AntibacterialOther:     ParseNullString2(c.FormValue("antibacterial_other")),
		AntibacterialDose:      ParseNullString2(c.FormValue("antibacterial_dose")),
		AntibacterialRoute:     ParseNullString2(c.FormValue("antibacterial_route")),
		AntibacterialFreq:      ParseNullString2(c.FormValue("antibacterial_freq")),
		Antimalarial:           ParseNullInt2(c.FormValue("antimalarial")),
		AntimalarialArtesunate: ParseNullInt2(c.FormValue("antimalarial_artesunate")),
		AntimalarialArthemeter: ParseNullInt2(c.FormValue("antimalarial_arthemeter")),
		AntimalarialAl:         ParseNullInt2(c.FormValue("antimalarial_al")),
		AntimalarialAa:         ParseNullInt2(c.FormValue("antimalarial_aa")),
		AntimalarialDose:       ParseNullString2(c.FormValue("antimalarial_dose")),
		AntimalarialRoute:      ParseNullString2(c.FormValue("antimalarial_route")),
		AntimalarialFreq:       ParseNullString2(c.FormValue("antimalarial_freq")),
		OtherMedsSpecify:       ParseNullString2(c.FormValue("other_meds_specify")),
		OtherMedsDose:          ParseNullString2(c.FormValue("other_meds_dose")),
		OtherMedsRoute:         ParseNullString2(c.FormValue("other_meds_route")),
		OtherMedsFreq:          ParseNullString2(c.FormValue("other_meds_freq")),
		EbolaExperimental:      ParseNullInt2(c.FormValue("ebola_experimental")),
		EbolaExperimentalIf:    ParseNullString2(c.FormValue("ebola_experimental_if")),
		Oral:                   ParseNullInt2(c.FormValue("oral")),
		OralOrs:                ParseNullInt2(c.FormValue("oral_ors")),
		OralOrsQty:             ParseNullFloat(c.FormValue("oral_ors_qty")),
		OralWater:              ParseNullInt2(c.FormValue("oral_water")),
		OralWaterQty:           ParseNullFloat(c.FormValue("oral_water_qty")),
		OralOther:              ParseNullInt2(c.FormValue("oral_other")),
		OralOtherQty:           ParseNullFloat(c.FormValue("oral_other_qty")),
		Iv:                     ParseNullInt2(c.FormValue("iv")),
		IvQty:                  ParseNullString2(c.FormValue("iv_qty")),
		IvUsing:                ParseNullString2(c.FormValue("iv_using")),
		IvAza:                  ParseNullString2(c.FormValue("iv_aza")),
		AccessType:             ParseNullInt2(c.FormValue("access_type")),
		BloodTrans:             ParseNullInt2(c.FormValue("blood_trans")),
		OxygenTherapy:          ParseNullInt2(c.FormValue("oxygen_therapy")),
		OxygenQty:              ParseNullFloat(c.FormValue("oxygen_qty")),
		OxygenWith:             ParseNullString2(c.FormValue("oxygen_with")),
		Vasopressors:           ParseNullInt2(c.FormValue("vasopressors")),
		Renal:                  ParseNullInt2(c.FormValue("renal")),
		Invasive:               ParseNullInt2(c.FormValue("invasive")),

		EbolaExperimentalIfZmap:     ParseNullInt2(c.FormValue("ebola_experimental_if_zmap")),
		EbolaExperimentalIfRemd:     ParseNullInt2(c.FormValue("ebola_experimental_if_remd")),
		EbolaExperimentalIfRegn:     ParseNullInt2(c.FormValue("ebola_experimental_if_regn")),
		EbolaExperimentalIfFavi:     ParseNullInt2(c.FormValue("ebola_experimental_if_favi")),
		EbolaExperimentalIfMab:      ParseNullInt2(c.FormValue("ebola_experimental_if_mab")),
		OralOtherAza:                ParseNullString2(c.FormValue("oral_other_aza")),
		AntibacterialAza:            ParseNullInt2(c.FormValue("antibacterial_aza")),
		AntimalarialArtesunateDose:  ParseNullString2(c.FormValue("antimalarial_artesunate_dose")),
		AntimalarialArtesunateRoute: ParseNullString2(c.FormValue("antimalarial_artesunate_route")),
		AntimalarialArtesunateFreq:  ParseNullString2(c.FormValue("antimalarial_artesunate_freq")),
		AntimalarialArthemeterDose:  ParseNullString2(c.FormValue("antimalarial_arthemeter_dose")),
		AntimalarialArthemeterRoute: ParseNullString2(c.FormValue("antimalarial_arthemeter_route")),
		AntimalarialArthemeterFreq:  ParseNullString2(c.FormValue("antimalarial_arthemeter_freq")),
		AntimalarialAlDose:          ParseNullString2(c.FormValue("antimalarial_al_dose")),
		AntimalarialAlRoute:         ParseNullString2(c.FormValue("antimalarial_al_route")),
		AntimalarialAlFreq:          ParseNullString2(c.FormValue("antimalarial_al_freq")),
		AntimalarialAaDose:          ParseNullString2(c.FormValue("antimalarial_aa_dose")),
		AntimalarialAaRoute:         ParseNullString2(c.FormValue("antimalarial_aa_route")),
		AntimalarialAaFreq:          ParseNullString2(c.FormValue("antimalarial_aa_freq")),
	}

	if treat_id == 0 {
		err := treat.Insert(c.Context(), db)
		if err != nil {
			fmt.Println(err.Error())
		}
	} else {
		treat.SetAsExists()
		err := treat.Update(c.Context(), db)
		if err != nil {
			fmt.Println(err.Error())
		}
	}

	// re route

	urlx := "/cases/encounters/new/" + cid + "/?dte=" + dte

	return c.Redirect(urlx)
}

func HandlerAPIGetEncounter(c *fiber.Ctx, db *sql.DB, sl *slog.Logger, store *session.Store, config Config) error {
	// Get ID from the query parameter

	id := c.Query("id")

	if id == "" {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "",
		})
	}

	encounter_id, err := strconv.Atoi(id)
	if err != nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "",
		})
	}

	var clinical = &models.Clinical{}
	var vital = &models.Vital{}

	clinical, _ = models.ClinicalByEncounterID(c.Context(), db, encounter_id)
	vital, _ = models.VitalByEncounterID(c.Context(), db, encounter_id)

	rtnStr := ` Vitals<br />
				<table class="full-width" border="1">
					<tr>
						<td>Weight: ` + fmt.Sprintf("%.2f", vital.Weight.Float64) + `</td>
						<td>Height: ` + fmt.Sprintf("%.2f", vital.Height.Float64) + `</td>
					</tr>
				</table>
				Symptomms<br/>
				<table class="full-width" border="1">
					<tr>
						<td valign="top">
							Fever: ` + strconv.Itoa(int(clinical.Fever.Int64)) + `<br/>
							Fatigue:<br/>
							Weakness:<br/>
							Malaise:<br/>
							Myalgia:<br/>
							Anorexia:<br/>
							Sore throat
						</td>
						<td valign="top">
							Headache:<br/> 
							Nausea:<br/> 
							Chest pain:<br/> 
							Joint Pain:<br/> 
							Hiccups:<br/>
							Cough:<br/>
						</td>
						<td valign="top">
							Chest pain:<br/>
							Difficulty breathing:<br/>
							Difficulty swallowing:<br/> 
							Abdominal pain:<br/> 
							Diarrhoea:<br/>
							Vomiting:<br/>
							Irritability / Confusion:<br/> 
						</td>
					</tr>
				</table>

				<br/>
				Signs<br/>
				<table class="full-width" border="1">
					<tr>
						<td valign="top">
							Pharyngeal erythema:<br/>  
							Pharyngeal exudate:<br/>  
							Conjunctival injection/bleeding:<br/>  
							Oedema of face/neck:<br/> 
							Tender abdomen:<br/> 
							Sunken eyes or fontanelle:<br/>  
							Tenting on skin pinch:<br/>  
							Palpable liver:<br/> 
							Palpable spleen Rash:<br/> 
							Jaundice:<br/> 

						</td>
						<td valign="top">
							Enlarged lymph nodes:<br/>
							Lower extremity oedema :<br/> 
							Bleeding:<br/> 
						</td>
					</tr>
				</table>
				<br/>
				Specimen <br/>
				<table class="full-width" border="1">
					<tr>
						<td valign="top">
						</td>
					</tr>
				</table>
				<br/>
				Lab Results <br/>
				<table class="full-width" border="1">
					<tr>
						<td valign="top">
						</td>
					</tr>
				</table>
				<br/>
				Medications <br/>
				<table class="full-width" border="1">
					<tr>
						<td valign="top">
						</td>
					</tr>
				</table>`

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": rtnStr,
	})

}

func HandlerAPIGetStatuses(c *fiber.Ctx, db *sql.DB, sl *slog.Logger, store *session.Store, config Config) error {
	userID := GetCurrentUser(c, store)

	// Check if user is logged in
	if userID == 0 {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
	}

	clientID := c.Query("client_id")
	if clientID == "" {
		clientID = "0"
	}

	statuses, er := models.Statusez(c.Context(), db, " client_id = "+clientID)
	if er != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error fetching statuses",
		})
	}

	return c.JSON(statuses)

}

func HandlerAPIPostStatus(c *fiber.Ctx, db *sql.DB, sl *slog.Logger, store *session.Store, config Config) error {

	//=================

	userID := GetCurrentUser(c, store)
	// Check if user is logged in
	if userID == 0 {
		fmt.Println("unauthorized")
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
	}

	//=============================

	var formData map[string]interface{}

	if err := c.BodyParser(&formData); err != nil {
		fmt.Println("JSON parsing failed:", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	var s models.Status

	s.StatusID = int(ParseNullInt2(formData["status_id"]).Int64)
	s.ClientID = ParseNullInt2(formData["client_id"])
	s.StatusDate = ParseNullString2(formData["status_date"])
	s.Status = ParseNullString2(formData["status"])
	s.StatusNotes = ParseNullString2(formData["status_notes"])

	s.UpdatedBy.Valid = true
	s.UpdatedBy.Int64 = int64(userID)

	s.UpdatedOn.Valid = true
	currentTime := time.Now()
	formattedTime := currentTime.Format("2006-01-02")
	s.UpdatedOn.String = formattedTime

	// Check if it's a new record (StatusID == 0)
	if s.StatusID > 0 {
		s.SetAsExists()
		err := s.Update(c.Context(), db)
		if err != nil {
			fmt.Println("update fail:", err.Error())
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}
	} else {

		err := s.Insert(c.Context(), db)
		if err != nil {
			fmt.Println("insert fail:", err.Error())
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "success",
	})
}
