package models

// Code generated by xo. DO NOT EDIT.

import (
	"context"
	"database/sql"
)

// Clinical represents a row from 'public.clinical'.
type Clinical struct {
	ClinicalID              int            `json:"clinical_id"`               // clinical_id
	EncounterID             sql.NullInt64  `json:"encounter_id"`              // encounter_id
	Fever                   sql.NullInt64  `json:"fever"`                     // fever
	Fatigue                 sql.NullInt64  `json:"fatigue"`                   // fatigue
	Weakness                sql.NullInt64  `json:"weakness"`                  // weakness
	Malaise                 sql.NullInt64  `json:"malaise"`                   // malaise
	Myalgia                 sql.NullInt64  `json:"myalgia"`                   // myalgia
	Anorexia                sql.NullInt64  `json:"anorexia"`                  // anorexia
	SoreThroat              sql.NullInt64  `json:"sore_throat"`               // sore_throat
	Headache                sql.NullInt64  `json:"headache"`                  // headache
	Nausea                  sql.NullInt64  `json:"nausea"`                    // nausea
	ChestPain               sql.NullInt64  `json:"chest_pain"`                // chest_pain
	JointPain               sql.NullInt64  `json:"joint_pain"`                // joint_pain
	Hiccups                 sql.NullInt64  `json:"hiccups"`                   // hiccups
	Cough                   sql.NullInt64  `json:"cough"`                     // cough
	DifficultyBreathing     sql.NullInt64  `json:"difficulty_breathing"`      // difficulty_breathing
	DifficultySwallowing    sql.NullInt64  `json:"difficulty_swallowing"`     // difficulty_swallowing
	AbdominalPain           sql.NullInt64  `json:"abdominal_pain"`            // abdominal_pain
	Diarrhoea               sql.NullInt64  `json:"diarrhoea"`                 // diarrhoea
	Vomiting                sql.NullInt64  `json:"vomiting"`                  // vomiting
	Irritability            sql.NullInt64  `json:"irritability"`              // irritability
	Dysphagia               sql.NullInt64  `json:"dysphagia"`                 // dysphagia
	UnusualBleeding         sql.NullInt64  `json:"unusual_bleeding"`          // unusual_bleeding
	Dehydration             sql.NullInt64  `json:"dehydration"`               // dehydration
	Shock                   sql.NullInt64  `json:"shock"`                     // shock
	Anuria                  sql.NullInt64  `json:"anuria"`                    // anuria
	Disorientation          sql.NullInt64  `json:"disorientation"`            // disorientation
	Agitation               sql.NullInt64  `json:"agitation"`                 // agitation
	Seizure                 sql.NullInt64  `json:"seizure"`                   // seizure
	Meningitis              sql.NullInt64  `json:"meningitis"`                // meningitis
	Confusion               sql.NullInt64  `json:"confusion"`                 // confusion
	Coma                    sql.NullInt64  `json:"coma"`                      // coma
	Bacteraemia             sql.NullInt64  `json:"bacteraemia"`               // bacteraemia
	Hyperglycemia           sql.NullInt64  `json:"hyperglycemia"`             // hyperglycemia
	Hypoglycemia            sql.NullInt64  `json:"hypoglycemia"`              // hypoglycemia
	OtherComplications      sql.NullInt64  `json:"other_complications"`       // other_complications
	AzaComplicationsSpecif  sql.NullString `json:"aza_complications_specif"`  // aza_complications_specif
	PharyngealErythema      sql.NullInt64  `json:"pharyngeal_erythema"`       // pharyngeal_erythema
	PharyngealExudate       sql.NullInt64  `json:"pharyngeal_exudate"`        // pharyngeal_exudate
	ConjunctivalInjection   sql.NullInt64  `json:"conjunctival_injection"`    // conjunctival_injection
	OedemaFace              sql.NullInt64  `json:"oedema_face"`               // oedema_face
	TenderAbdomen           sql.NullInt64  `json:"tender_abdomen"`            // tender_abdomen
	SunkenEyes              sql.NullInt64  `json:"sunken_eyes"`               // sunken_eyes
	TentingSkin             sql.NullInt64  `json:"tenting_skin"`              // tenting_skin
	PalpableLiver           sql.NullInt64  `json:"palpable_liver"`            // palpable_liver
	PalpableSpleen          sql.NullInt64  `json:"palpable_spleen"`           // palpable_spleen
	Jaundice                sql.NullInt64  `json:"jaundice"`                  // jaundice
	EnlargedLymphNodes      sql.NullInt64  `json:"enlarged_lymph_nodes"`      // enlarged_lymph_nodes
	LowerExtremityOedema    sql.NullInt64  `json:"lower_extremity_oedema"`    // lower_extremity_oedema
	Bleeding                sql.NullInt64  `json:"bleeding"`                  // bleeding
	BleedingNose            sql.NullInt64  `json:"bleeding_nose"`             // bleeding_nose
	BleedingMouth           sql.NullInt64  `json:"bleeding_mouth"`            // bleeding_mouth
	BleedingVagina          sql.NullInt64  `json:"bleeding_vagina"`           // bleeding_vagina
	BleedingRectum          sql.NullInt64  `json:"bleeding_rectum"`           // bleeding_rectum
	BleedingSputum          sql.NullInt64  `json:"bleeding_sputum"`           // bleeding_sputum
	BleedingUrine           sql.NullInt64  `json:"bleeding_urine"`            // bleeding_urine
	BleedingIvSite          sql.NullInt64  `json:"bleeding_iv_site"`          // bleeding_iv_site
	BleedingOther           sql.NullInt64  `json:"bleeding_other"`            // bleeding_other
	BleedingOtherSpecif     sql.NullString `json:"bleeding_other_specif"`     // bleeding_other_specif
	FinalDiagnosis          sql.NullInt64  `json:"final_diagnosis"`           // final_diagnosis
	FinalDiagnosisAza       sql.NullString `json:"final_diagnosis_aza"`       // final_diagnosis_aza
	OutcomeDischarge        sql.NullInt64  `json:"outcome_discharge"`         // outcome_discharge
	OutcomeDischargeIfHear  sql.NullInt64  `json:"outcome_discharge_if_hear"` // outcome_discharge_if_hear
	OutcomeDischargeIfArth  sql.NullInt64  `json:"outcome_discharge_if_arth"` // outcome_discharge_if_arth
	OutcomeDischargeIfAbor  sql.NullInt64  `json:"outcome_discharge_if_abor"` // outcome_discharge_if_abor
	OutcomeDischargeIfNeur  sql.NullInt64  `json:"outcome_discharge_if_neur"` // outcome_discharge_if_neur
	OutcomeDischargeIfOcul  sql.NullInt64  `json:"outcome_discharge_if_ocul"` // outcome_discharge_if_ocul
	OutcomeDischargeIfExtr  sql.NullInt64  `json:"outcome_discharge_if_extr"` // outcome_discharge_if_extr
	OutcomeDischargeIfOthe  sql.NullInt64  `json:"outcome_discharge_if_othe"` // outcome_discharge_if_othe
	OutcomeDischargeIfAza   sql.NullString `json:"outcome_discharge_if_aza"`  // outcome_discharge_if_aza
	OutcomeReferredFacility sql.NullString `json:"outcome_referred_facility"` // outcome_referred_facility
	DischargeDate           sql.NullString `json:"discharge_date"`            // discharge_date
	SurvivorCounselling     sql.NullInt64  `json:"survivor_counselling"`      // survivor_counselling
	// xo fields
	_exists, _deleted bool
}

// Exists returns true when the [Clinical] exists in the database.
func (c *Clinical) Exists() bool {
	return c._exists
}

// Deleted returns true when the [Clinical] has been marked for deletion
// from the database.
func (c *Clinical) Deleted() bool {
	return c._deleted
}

// Insert inserts the [Clinical] to the database.
func (c *Clinical) Insert(ctx context.Context, db DB) error {
	switch {
	case c._exists: // already exists
		return logerror(&ErrInsertFailed{ErrAlreadyExists})
	case c._deleted: // deleted
		return logerror(&ErrInsertFailed{ErrMarkedForDeletion})
	}
	// insert (primary key generated and returned by database)
	const sqlstr = `INSERT INTO public.clinical (` +
		`encounter_id, fever, fatigue, weakness, malaise, myalgia, anorexia, sore_throat, headache, nausea, chest_pain, joint_pain, hiccups, cough, difficulty_breathing, difficulty_swallowing, abdominal_pain, diarrhoea, vomiting, irritability, dysphagia, unusual_bleeding, dehydration, shock, anuria, disorientation, agitation, seizure, meningitis, confusion, coma, bacteraemia, hyperglycemia, hypoglycemia, other_complications, aza_complications_specif, pharyngeal_erythema, pharyngeal_exudate, conjunctival_injection, oedema_face, tender_abdomen, sunken_eyes, tenting_skin, palpable_liver, palpable_spleen, jaundice, enlarged_lymph_nodes, lower_extremity_oedema, bleeding, bleeding_nose, bleeding_mouth, bleeding_vagina, bleeding_rectum, bleeding_sputum, bleeding_urine, bleeding_iv_site, bleeding_other, bleeding_other_specif, final_diagnosis, final_diagnosis_aza, outcome_discharge, outcome_discharge_if_hear, outcome_discharge_if_arth, outcome_discharge_if_abor, outcome_discharge_if_neur, outcome_discharge_if_ocul, outcome_discharge_if_extr, outcome_discharge_if_othe, outcome_discharge_if_aza, outcome_referred_facility, discharge_date, survivor_counselling` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25, $26, $27, $28, $29, $30, $31, $32, $33, $34, $35, $36, $37, $38, $39, $40, $41, $42, $43, $44, $45, $46, $47, $48, $49, $50, $51, $52, $53, $54, $55, $56, $57, $58, $59, $60, $61, $62, $63, $64, $65, $66, $67, $68, $69, $70, $71, $72` +
		`) RETURNING clinical_id`
	// run
	logf(sqlstr, c.EncounterID, c.Fever, c.Fatigue, c.Weakness, c.Malaise, c.Myalgia, c.Anorexia, c.SoreThroat, c.Headache, c.Nausea, c.ChestPain, c.JointPain, c.Hiccups, c.Cough, c.DifficultyBreathing, c.DifficultySwallowing, c.AbdominalPain, c.Diarrhoea, c.Vomiting, c.Irritability, c.Dysphagia, c.UnusualBleeding, c.Dehydration, c.Shock, c.Anuria, c.Disorientation, c.Agitation, c.Seizure, c.Meningitis, c.Confusion, c.Coma, c.Bacteraemia, c.Hyperglycemia, c.Hypoglycemia, c.OtherComplications, c.AzaComplicationsSpecif, c.PharyngealErythema, c.PharyngealExudate, c.ConjunctivalInjection, c.OedemaFace, c.TenderAbdomen, c.SunkenEyes, c.TentingSkin, c.PalpableLiver, c.PalpableSpleen, c.Jaundice, c.EnlargedLymphNodes, c.LowerExtremityOedema, c.Bleeding, c.BleedingNose, c.BleedingMouth, c.BleedingVagina, c.BleedingRectum, c.BleedingSputum, c.BleedingUrine, c.BleedingIvSite, c.BleedingOther, c.BleedingOtherSpecif, c.FinalDiagnosis, c.FinalDiagnosisAza, c.OutcomeDischarge, c.OutcomeDischargeIfHear, c.OutcomeDischargeIfArth, c.OutcomeDischargeIfAbor, c.OutcomeDischargeIfNeur, c.OutcomeDischargeIfOcul, c.OutcomeDischargeIfExtr, c.OutcomeDischargeIfOthe, c.OutcomeDischargeIfAza, c.OutcomeReferredFacility, c.DischargeDate, c.SurvivorCounselling)
	if err := db.QueryRowContext(ctx, sqlstr, c.EncounterID, c.Fever, c.Fatigue, c.Weakness, c.Malaise, c.Myalgia, c.Anorexia, c.SoreThroat, c.Headache, c.Nausea, c.ChestPain, c.JointPain, c.Hiccups, c.Cough, c.DifficultyBreathing, c.DifficultySwallowing, c.AbdominalPain, c.Diarrhoea, c.Vomiting, c.Irritability, c.Dysphagia, c.UnusualBleeding, c.Dehydration, c.Shock, c.Anuria, c.Disorientation, c.Agitation, c.Seizure, c.Meningitis, c.Confusion, c.Coma, c.Bacteraemia, c.Hyperglycemia, c.Hypoglycemia, c.OtherComplications, c.AzaComplicationsSpecif, c.PharyngealErythema, c.PharyngealExudate, c.ConjunctivalInjection, c.OedemaFace, c.TenderAbdomen, c.SunkenEyes, c.TentingSkin, c.PalpableLiver, c.PalpableSpleen, c.Jaundice, c.EnlargedLymphNodes, c.LowerExtremityOedema, c.Bleeding, c.BleedingNose, c.BleedingMouth, c.BleedingVagina, c.BleedingRectum, c.BleedingSputum, c.BleedingUrine, c.BleedingIvSite, c.BleedingOther, c.BleedingOtherSpecif, c.FinalDiagnosis, c.FinalDiagnosisAza, c.OutcomeDischarge, c.OutcomeDischargeIfHear, c.OutcomeDischargeIfArth, c.OutcomeDischargeIfAbor, c.OutcomeDischargeIfNeur, c.OutcomeDischargeIfOcul, c.OutcomeDischargeIfExtr, c.OutcomeDischargeIfOthe, c.OutcomeDischargeIfAza, c.OutcomeReferredFacility, c.DischargeDate, c.SurvivorCounselling).Scan(&c.ClinicalID); err != nil {
		return logerror(err)
	}
	// set exists
	c._exists = true
	return nil
}

// Update updates a [Clinical] in the database.
func (c *Clinical) Update(ctx context.Context, db DB) error {
	switch {
	case !c._exists: // doesn't exist
		return logerror(&ErrUpdateFailed{ErrDoesNotExist})
	case c._deleted: // deleted
		return logerror(&ErrUpdateFailed{ErrMarkedForDeletion})
	}
	// update with composite primary key
	const sqlstr = `UPDATE public.clinical SET ` +
		`encounter_id = $1, fever = $2, fatigue = $3, weakness = $4, malaise = $5, myalgia = $6, anorexia = $7, sore_throat = $8, headache = $9, nausea = $10, chest_pain = $11, joint_pain = $12, hiccups = $13, cough = $14, difficulty_breathing = $15, difficulty_swallowing = $16, abdominal_pain = $17, diarrhoea = $18, vomiting = $19, irritability = $20, dysphagia = $21, unusual_bleeding = $22, dehydration = $23, shock = $24, anuria = $25, disorientation = $26, agitation = $27, seizure = $28, meningitis = $29, confusion = $30, coma = $31, bacteraemia = $32, hyperglycemia = $33, hypoglycemia = $34, other_complications = $35, aza_complications_specif = $36, pharyngeal_erythema = $37, pharyngeal_exudate = $38, conjunctival_injection = $39, oedema_face = $40, tender_abdomen = $41, sunken_eyes = $42, tenting_skin = $43, palpable_liver = $44, palpable_spleen = $45, jaundice = $46, enlarged_lymph_nodes = $47, lower_extremity_oedema = $48, bleeding = $49, bleeding_nose = $50, bleeding_mouth = $51, bleeding_vagina = $52, bleeding_rectum = $53, bleeding_sputum = $54, bleeding_urine = $55, bleeding_iv_site = $56, bleeding_other = $57, bleeding_other_specif = $58, final_diagnosis = $59, final_diagnosis_aza = $60, outcome_discharge = $61, outcome_discharge_if_hear = $62, outcome_discharge_if_arth = $63, outcome_discharge_if_abor = $64, outcome_discharge_if_neur = $65, outcome_discharge_if_ocul = $66, outcome_discharge_if_extr = $67, outcome_discharge_if_othe = $68, outcome_discharge_if_aza = $69, outcome_referred_facility = $70, discharge_date = $71, survivor_counselling = $72 ` +
		`WHERE clinical_id = $73`
	// run
	logf(sqlstr, c.EncounterID, c.Fever, c.Fatigue, c.Weakness, c.Malaise, c.Myalgia, c.Anorexia, c.SoreThroat, c.Headache, c.Nausea, c.ChestPain, c.JointPain, c.Hiccups, c.Cough, c.DifficultyBreathing, c.DifficultySwallowing, c.AbdominalPain, c.Diarrhoea, c.Vomiting, c.Irritability, c.Dysphagia, c.UnusualBleeding, c.Dehydration, c.Shock, c.Anuria, c.Disorientation, c.Agitation, c.Seizure, c.Meningitis, c.Confusion, c.Coma, c.Bacteraemia, c.Hyperglycemia, c.Hypoglycemia, c.OtherComplications, c.AzaComplicationsSpecif, c.PharyngealErythema, c.PharyngealExudate, c.ConjunctivalInjection, c.OedemaFace, c.TenderAbdomen, c.SunkenEyes, c.TentingSkin, c.PalpableLiver, c.PalpableSpleen, c.Jaundice, c.EnlargedLymphNodes, c.LowerExtremityOedema, c.Bleeding, c.BleedingNose, c.BleedingMouth, c.BleedingVagina, c.BleedingRectum, c.BleedingSputum, c.BleedingUrine, c.BleedingIvSite, c.BleedingOther, c.BleedingOtherSpecif, c.FinalDiagnosis, c.FinalDiagnosisAza, c.OutcomeDischarge, c.OutcomeDischargeIfHear, c.OutcomeDischargeIfArth, c.OutcomeDischargeIfAbor, c.OutcomeDischargeIfNeur, c.OutcomeDischargeIfOcul, c.OutcomeDischargeIfExtr, c.OutcomeDischargeIfOthe, c.OutcomeDischargeIfAza, c.OutcomeReferredFacility, c.DischargeDate, c.SurvivorCounselling, c.ClinicalID)
	if _, err := db.ExecContext(ctx, sqlstr, c.EncounterID, c.Fever, c.Fatigue, c.Weakness, c.Malaise, c.Myalgia, c.Anorexia, c.SoreThroat, c.Headache, c.Nausea, c.ChestPain, c.JointPain, c.Hiccups, c.Cough, c.DifficultyBreathing, c.DifficultySwallowing, c.AbdominalPain, c.Diarrhoea, c.Vomiting, c.Irritability, c.Dysphagia, c.UnusualBleeding, c.Dehydration, c.Shock, c.Anuria, c.Disorientation, c.Agitation, c.Seizure, c.Meningitis, c.Confusion, c.Coma, c.Bacteraemia, c.Hyperglycemia, c.Hypoglycemia, c.OtherComplications, c.AzaComplicationsSpecif, c.PharyngealErythema, c.PharyngealExudate, c.ConjunctivalInjection, c.OedemaFace, c.TenderAbdomen, c.SunkenEyes, c.TentingSkin, c.PalpableLiver, c.PalpableSpleen, c.Jaundice, c.EnlargedLymphNodes, c.LowerExtremityOedema, c.Bleeding, c.BleedingNose, c.BleedingMouth, c.BleedingVagina, c.BleedingRectum, c.BleedingSputum, c.BleedingUrine, c.BleedingIvSite, c.BleedingOther, c.BleedingOtherSpecif, c.FinalDiagnosis, c.FinalDiagnosisAza, c.OutcomeDischarge, c.OutcomeDischargeIfHear, c.OutcomeDischargeIfArth, c.OutcomeDischargeIfAbor, c.OutcomeDischargeIfNeur, c.OutcomeDischargeIfOcul, c.OutcomeDischargeIfExtr, c.OutcomeDischargeIfOthe, c.OutcomeDischargeIfAza, c.OutcomeReferredFacility, c.DischargeDate, c.SurvivorCounselling, c.ClinicalID); err != nil {
		return logerror(err)
	}
	return nil
}

// Save saves the [Clinical] to the database.
func (c *Clinical) Save(ctx context.Context, db DB) error {
	if c.Exists() {
		return c.Update(ctx, db)
	}
	return c.Insert(ctx, db)
}

// Upsert performs an upsert for [Clinical].
func (c *Clinical) Upsert(ctx context.Context, db DB) error {
	switch {
	case c._deleted: // deleted
		return logerror(&ErrUpsertFailed{ErrMarkedForDeletion})
	}
	// upsert
	const sqlstr = `INSERT INTO public.clinical (` +
		`clinical_id, encounter_id, fever, fatigue, weakness, malaise, myalgia, anorexia, sore_throat, headache, nausea, chest_pain, joint_pain, hiccups, cough, difficulty_breathing, difficulty_swallowing, abdominal_pain, diarrhoea, vomiting, irritability, dysphagia, unusual_bleeding, dehydration, shock, anuria, disorientation, agitation, seizure, meningitis, confusion, coma, bacteraemia, hyperglycemia, hypoglycemia, other_complications, aza_complications_specif, pharyngeal_erythema, pharyngeal_exudate, conjunctival_injection, oedema_face, tender_abdomen, sunken_eyes, tenting_skin, palpable_liver, palpable_spleen, jaundice, enlarged_lymph_nodes, lower_extremity_oedema, bleeding, bleeding_nose, bleeding_mouth, bleeding_vagina, bleeding_rectum, bleeding_sputum, bleeding_urine, bleeding_iv_site, bleeding_other, bleeding_other_specif, final_diagnosis, final_diagnosis_aza, outcome_discharge, outcome_discharge_if_hear, outcome_discharge_if_arth, outcome_discharge_if_abor, outcome_discharge_if_neur, outcome_discharge_if_ocul, outcome_discharge_if_extr, outcome_discharge_if_othe, outcome_discharge_if_aza, outcome_referred_facility, discharge_date, survivor_counselling` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25, $26, $27, $28, $29, $30, $31, $32, $33, $34, $35, $36, $37, $38, $39, $40, $41, $42, $43, $44, $45, $46, $47, $48, $49, $50, $51, $52, $53, $54, $55, $56, $57, $58, $59, $60, $61, $62, $63, $64, $65, $66, $67, $68, $69, $70, $71, $72, $73` +
		`)` +
		` ON CONFLICT (clinical_id) DO ` +
		`UPDATE SET ` +
		`encounter_id = EXCLUDED.encounter_id, fever = EXCLUDED.fever, fatigue = EXCLUDED.fatigue, weakness = EXCLUDED.weakness, malaise = EXCLUDED.malaise, myalgia = EXCLUDED.myalgia, anorexia = EXCLUDED.anorexia, sore_throat = EXCLUDED.sore_throat, headache = EXCLUDED.headache, nausea = EXCLUDED.nausea, chest_pain = EXCLUDED.chest_pain, joint_pain = EXCLUDED.joint_pain, hiccups = EXCLUDED.hiccups, cough = EXCLUDED.cough, difficulty_breathing = EXCLUDED.difficulty_breathing, difficulty_swallowing = EXCLUDED.difficulty_swallowing, abdominal_pain = EXCLUDED.abdominal_pain, diarrhoea = EXCLUDED.diarrhoea, vomiting = EXCLUDED.vomiting, irritability = EXCLUDED.irritability, dysphagia = EXCLUDED.dysphagia, unusual_bleeding = EXCLUDED.unusual_bleeding, dehydration = EXCLUDED.dehydration, shock = EXCLUDED.shock, anuria = EXCLUDED.anuria, disorientation = EXCLUDED.disorientation, agitation = EXCLUDED.agitation, seizure = EXCLUDED.seizure, meningitis = EXCLUDED.meningitis, confusion = EXCLUDED.confusion, coma = EXCLUDED.coma, bacteraemia = EXCLUDED.bacteraemia, hyperglycemia = EXCLUDED.hyperglycemia, hypoglycemia = EXCLUDED.hypoglycemia, other_complications = EXCLUDED.other_complications, aza_complications_specif = EXCLUDED.aza_complications_specif, pharyngeal_erythema = EXCLUDED.pharyngeal_erythema, pharyngeal_exudate = EXCLUDED.pharyngeal_exudate, conjunctival_injection = EXCLUDED.conjunctival_injection, oedema_face = EXCLUDED.oedema_face, tender_abdomen = EXCLUDED.tender_abdomen, sunken_eyes = EXCLUDED.sunken_eyes, tenting_skin = EXCLUDED.tenting_skin, palpable_liver = EXCLUDED.palpable_liver, palpable_spleen = EXCLUDED.palpable_spleen, jaundice = EXCLUDED.jaundice, enlarged_lymph_nodes = EXCLUDED.enlarged_lymph_nodes, lower_extremity_oedema = EXCLUDED.lower_extremity_oedema, bleeding = EXCLUDED.bleeding, bleeding_nose = EXCLUDED.bleeding_nose, bleeding_mouth = EXCLUDED.bleeding_mouth, bleeding_vagina = EXCLUDED.bleeding_vagina, bleeding_rectum = EXCLUDED.bleeding_rectum, bleeding_sputum = EXCLUDED.bleeding_sputum, bleeding_urine = EXCLUDED.bleeding_urine, bleeding_iv_site = EXCLUDED.bleeding_iv_site, bleeding_other = EXCLUDED.bleeding_other, bleeding_other_specif = EXCLUDED.bleeding_other_specif, final_diagnosis = EXCLUDED.final_diagnosis, final_diagnosis_aza = EXCLUDED.final_diagnosis_aza, outcome_discharge = EXCLUDED.outcome_discharge, outcome_discharge_if_hear = EXCLUDED.outcome_discharge_if_hear, outcome_discharge_if_arth = EXCLUDED.outcome_discharge_if_arth, outcome_discharge_if_abor = EXCLUDED.outcome_discharge_if_abor, outcome_discharge_if_neur = EXCLUDED.outcome_discharge_if_neur, outcome_discharge_if_ocul = EXCLUDED.outcome_discharge_if_ocul, outcome_discharge_if_extr = EXCLUDED.outcome_discharge_if_extr, outcome_discharge_if_othe = EXCLUDED.outcome_discharge_if_othe, outcome_discharge_if_aza = EXCLUDED.outcome_discharge_if_aza, outcome_referred_facility = EXCLUDED.outcome_referred_facility, discharge_date = EXCLUDED.discharge_date, survivor_counselling = EXCLUDED.survivor_counselling `
	// run
	logf(sqlstr, c.ClinicalID, c.EncounterID, c.Fever, c.Fatigue, c.Weakness, c.Malaise, c.Myalgia, c.Anorexia, c.SoreThroat, c.Headache, c.Nausea, c.ChestPain, c.JointPain, c.Hiccups, c.Cough, c.DifficultyBreathing, c.DifficultySwallowing, c.AbdominalPain, c.Diarrhoea, c.Vomiting, c.Irritability, c.Dysphagia, c.UnusualBleeding, c.Dehydration, c.Shock, c.Anuria, c.Disorientation, c.Agitation, c.Seizure, c.Meningitis, c.Confusion, c.Coma, c.Bacteraemia, c.Hyperglycemia, c.Hypoglycemia, c.OtherComplications, c.AzaComplicationsSpecif, c.PharyngealErythema, c.PharyngealExudate, c.ConjunctivalInjection, c.OedemaFace, c.TenderAbdomen, c.SunkenEyes, c.TentingSkin, c.PalpableLiver, c.PalpableSpleen, c.Jaundice, c.EnlargedLymphNodes, c.LowerExtremityOedema, c.Bleeding, c.BleedingNose, c.BleedingMouth, c.BleedingVagina, c.BleedingRectum, c.BleedingSputum, c.BleedingUrine, c.BleedingIvSite, c.BleedingOther, c.BleedingOtherSpecif, c.FinalDiagnosis, c.FinalDiagnosisAza, c.OutcomeDischarge, c.OutcomeDischargeIfHear, c.OutcomeDischargeIfArth, c.OutcomeDischargeIfAbor, c.OutcomeDischargeIfNeur, c.OutcomeDischargeIfOcul, c.OutcomeDischargeIfExtr, c.OutcomeDischargeIfOthe, c.OutcomeDischargeIfAza, c.OutcomeReferredFacility, c.DischargeDate, c.SurvivorCounselling)
	if _, err := db.ExecContext(ctx, sqlstr, c.ClinicalID, c.EncounterID, c.Fever, c.Fatigue, c.Weakness, c.Malaise, c.Myalgia, c.Anorexia, c.SoreThroat, c.Headache, c.Nausea, c.ChestPain, c.JointPain, c.Hiccups, c.Cough, c.DifficultyBreathing, c.DifficultySwallowing, c.AbdominalPain, c.Diarrhoea, c.Vomiting, c.Irritability, c.Dysphagia, c.UnusualBleeding, c.Dehydration, c.Shock, c.Anuria, c.Disorientation, c.Agitation, c.Seizure, c.Meningitis, c.Confusion, c.Coma, c.Bacteraemia, c.Hyperglycemia, c.Hypoglycemia, c.OtherComplications, c.AzaComplicationsSpecif, c.PharyngealErythema, c.PharyngealExudate, c.ConjunctivalInjection, c.OedemaFace, c.TenderAbdomen, c.SunkenEyes, c.TentingSkin, c.PalpableLiver, c.PalpableSpleen, c.Jaundice, c.EnlargedLymphNodes, c.LowerExtremityOedema, c.Bleeding, c.BleedingNose, c.BleedingMouth, c.BleedingVagina, c.BleedingRectum, c.BleedingSputum, c.BleedingUrine, c.BleedingIvSite, c.BleedingOther, c.BleedingOtherSpecif, c.FinalDiagnosis, c.FinalDiagnosisAza, c.OutcomeDischarge, c.OutcomeDischargeIfHear, c.OutcomeDischargeIfArth, c.OutcomeDischargeIfAbor, c.OutcomeDischargeIfNeur, c.OutcomeDischargeIfOcul, c.OutcomeDischargeIfExtr, c.OutcomeDischargeIfOthe, c.OutcomeDischargeIfAza, c.OutcomeReferredFacility, c.DischargeDate, c.SurvivorCounselling); err != nil {
		return logerror(err)
	}
	// set exists
	c._exists = true
	return nil
}

// Delete deletes the [Clinical] from the database.
func (c *Clinical) Delete(ctx context.Context, db DB) error {
	switch {
	case !c._exists: // doesn't exist
		return nil
	case c._deleted: // deleted
		return nil
	}
	// delete with single primary key
	const sqlstr = `DELETE FROM public.clinical ` +
		`WHERE clinical_id = $1`
	// run
	logf(sqlstr, c.ClinicalID)
	if _, err := db.ExecContext(ctx, sqlstr, c.ClinicalID); err != nil {
		return logerror(err)
	}
	// set deleted
	c._deleted = true
	return nil
}

// ClinicalByClinicalID retrieves a row from 'public.clinical' as a [Clinical].
//
// Generated from index 'clinical_pkey'.
func ClinicalByClinicalID(ctx context.Context, db DB, clinicalID int) (*Clinical, error) {
	// query
	const sqlstr = `SELECT ` +
		`clinical_id, encounter_id, fever, fatigue, weakness, malaise, myalgia, anorexia, sore_throat, headache, nausea, chest_pain, joint_pain, hiccups, cough, difficulty_breathing, difficulty_swallowing, abdominal_pain, diarrhoea, vomiting, irritability, dysphagia, unusual_bleeding, dehydration, shock, anuria, disorientation, agitation, seizure, meningitis, confusion, coma, bacteraemia, hyperglycemia, hypoglycemia, other_complications, aza_complications_specif, pharyngeal_erythema, pharyngeal_exudate, conjunctival_injection, oedema_face, tender_abdomen, sunken_eyes, tenting_skin, palpable_liver, palpable_spleen, jaundice, enlarged_lymph_nodes, lower_extremity_oedema, bleeding, bleeding_nose, bleeding_mouth, bleeding_vagina, bleeding_rectum, bleeding_sputum, bleeding_urine, bleeding_iv_site, bleeding_other, bleeding_other_specif, final_diagnosis, final_diagnosis_aza, outcome_discharge, outcome_discharge_if_hear, outcome_discharge_if_arth, outcome_discharge_if_abor, outcome_discharge_if_neur, outcome_discharge_if_ocul, outcome_discharge_if_extr, outcome_discharge_if_othe, outcome_discharge_if_aza, outcome_referred_facility, discharge_date, survivor_counselling ` +
		`FROM public.clinical ` +
		`WHERE clinical_id = $1`
	// run
	logf(sqlstr, clinicalID)
	c := Clinical{
		_exists: true,
	}
	if err := db.QueryRowContext(ctx, sqlstr, clinicalID).Scan(&c.ClinicalID, &c.EncounterID, &c.Fever, &c.Fatigue, &c.Weakness, &c.Malaise, &c.Myalgia, &c.Anorexia, &c.SoreThroat, &c.Headache, &c.Nausea, &c.ChestPain, &c.JointPain, &c.Hiccups, &c.Cough, &c.DifficultyBreathing, &c.DifficultySwallowing, &c.AbdominalPain, &c.Diarrhoea, &c.Vomiting, &c.Irritability, &c.Dysphagia, &c.UnusualBleeding, &c.Dehydration, &c.Shock, &c.Anuria, &c.Disorientation, &c.Agitation, &c.Seizure, &c.Meningitis, &c.Confusion, &c.Coma, &c.Bacteraemia, &c.Hyperglycemia, &c.Hypoglycemia, &c.OtherComplications, &c.AzaComplicationsSpecif, &c.PharyngealErythema, &c.PharyngealExudate, &c.ConjunctivalInjection, &c.OedemaFace, &c.TenderAbdomen, &c.SunkenEyes, &c.TentingSkin, &c.PalpableLiver, &c.PalpableSpleen, &c.Jaundice, &c.EnlargedLymphNodes, &c.LowerExtremityOedema, &c.Bleeding, &c.BleedingNose, &c.BleedingMouth, &c.BleedingVagina, &c.BleedingRectum, &c.BleedingSputum, &c.BleedingUrine, &c.BleedingIvSite, &c.BleedingOther, &c.BleedingOtherSpecif, &c.FinalDiagnosis, &c.FinalDiagnosisAza, &c.OutcomeDischarge, &c.OutcomeDischargeIfHear, &c.OutcomeDischargeIfArth, &c.OutcomeDischargeIfAbor, &c.OutcomeDischargeIfNeur, &c.OutcomeDischargeIfOcul, &c.OutcomeDischargeIfExtr, &c.OutcomeDischargeIfOthe, &c.OutcomeDischargeIfAza, &c.OutcomeReferredFacility, &c.DischargeDate, &c.SurvivorCounselling); err != nil {
		return nil, logerror(err)
	}
	return &c, nil
}
