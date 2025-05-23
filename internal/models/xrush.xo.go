package models

// Code generated by xo. DO NOT EDIT.

import (
	"context"
	"database/sql"
)

// XRush represents a row from 'public.x_rush'.
type XRush struct {
	ID                   int            `json:"id"`                     // id
	EncounterID          sql.NullInt64  `json:"encounter_id"`           // encounter_id
	LesionsBody          sql.NullInt64  `json:"lesions_body"`           // lesions_body
	LesionsRightLeg      sql.NullInt64  `json:"lesions_right_leg"`      // lesions_right_leg
	LesionsRightArm      sql.NullInt64  `json:"lesions_right_arm"`      // lesions_right_arm
	LesionsLeftLeg       sql.NullInt64  `json:"lesions_left_leg"`       // lesions_left_leg
	LesionsLeftArm       sql.NullInt64  `json:"lesions_left_arm"`       // lesions_left_arm
	LesionsGenitals      sql.NullInt64  `json:"lesions_genitals"`       // lesions_genitals
	LesionsOral          sql.NullInt64  `json:"lesions_oral"`           // lesions_oral
	LesionsPerianal1     sql.NullInt64  `json:"lesions_perianal_1"`     // lesions_perianal_1
	Face                 sql.NullInt64  `json:"face"`                   // face
	Nares                sql.NullInt64  `json:"nares"`                  // nares
	Mouth                sql.NullInt64  `json:"mouth"`                  // mouth
	Chest                sql.NullInt64  `json:"chest"`                  // chest
	Abdomen              sql.NullInt64  `json:"abdomen"`                // abdomen
	Back                 sql.NullInt64  `json:"back"`                   // back
	Perianal             sql.NullInt64  `json:"perianal"`               // perianal
	Genitals             sql.NullInt64  `json:"genitals"`               // genitals
	Palms                sql.NullInt64  `json:"palms"`                  // palms
	Arms                 sql.NullInt64  `json:"arms"`                   // arms
	Forearms             sql.NullInt64  `json:"forearms"`               // forearms
	Thighs               sql.NullInt64  `json:"thighs"`                 // thighs
	Legs                 sql.NullInt64  `json:"legs"`                   // legs
	Soles                sql.NullInt64  `json:"soles"`                  // soles
	Other                sql.NullInt64  `json:"other"`                  // other
	OtherSpecify         sql.NullString `json:"other_specify"`          // other_specify
	Macule               sql.NullInt64  `json:"macule"`                 // macule
	Papule               sql.NullInt64  `json:"papule"`                 // papule
	EarlyVesicle         sql.NullInt64  `json:"early_vesicle"`          // early_vesicle
	SmallPustule         sql.NullInt64  `json:"small_pustule"`          // small_pustule
	UmbilicatedPustule   sql.NullInt64  `json:"umbilicated_pustule"`    // umbilicated_pustule
	UlceratedLesion      sql.NullInt64  `json:"ulcerated_lesion"`       // ulcerated_lesion
	CrustingMature       sql.NullInt64  `json:"crusting_mature"`        // crusting_mature
	PartiallyRemoved     sql.NullInt64  `json:"partially_removed"`      // partially_removed
	PainAtSite           sql.NullInt64  `json:"pain_at_site"`           // pain_at_site
	PainScore            sql.NullInt64  `json:"pain_score"`             // pain_score
	LesionsPerianal2     sql.NullInt64  `json:"lesions_perianal_2"`     // lesions_perianal_2
	ShadeLocationLesions sql.NullInt64  `json:"shade_location_lesions"` // shade_location_lesions
	EnterBy              sql.NullInt64  `json:"enter_by"`               // enter_by
	EnterOn              sql.NullTime   `json:"enter_on"`               // enter_on
	// xo fields
	_exists, _deleted bool
}

// Exists returns true when the [XRush] exists in the database.
func (xr *XRush) Exists() bool {
	return xr._exists
}

// Deleted returns true when the [XRush] has been marked for deletion
// from the database.
func (xr *XRush) Deleted() bool {
	return xr._deleted
}

// Insert inserts the [XRush] to the database.
func (xr *XRush) Insert(ctx context.Context, db DB) error {
	switch {
	case xr._exists: // already exists
		return logerror(&ErrInsertFailed{ErrAlreadyExists})
	case xr._deleted: // deleted
		return logerror(&ErrInsertFailed{ErrMarkedForDeletion})
	}
	// insert (primary key generated and returned by database)
	const sqlstr = `INSERT INTO public.x_rush (` +
		`encounter_id, lesions_body, lesions_right_leg, lesions_right_arm, lesions_left_leg, lesions_left_arm, lesions_genitals, lesions_oral, lesions_perianal_1, face, nares, mouth, chest, abdomen, back, perianal, genitals, palms, arms, forearms, thighs, legs, soles, other, other_specify, macule, papule, early_vesicle, small_pustule, umbilicated_pustule, ulcerated_lesion, crusting_mature, partially_removed, pain_at_site, pain_score, lesions_perianal_2, shade_location_lesions, enter_by, enter_on` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25, $26, $27, $28, $29, $30, $31, $32, $33, $34, $35, $36, $37, $38, $39` +
		`) RETURNING id`
	// run
	logf(sqlstr, xr.EncounterID, xr.LesionsBody, xr.LesionsRightLeg, xr.LesionsRightArm, xr.LesionsLeftLeg, xr.LesionsLeftArm, xr.LesionsGenitals, xr.LesionsOral, xr.LesionsPerianal1, xr.Face, xr.Nares, xr.Mouth, xr.Chest, xr.Abdomen, xr.Back, xr.Perianal, xr.Genitals, xr.Palms, xr.Arms, xr.Forearms, xr.Thighs, xr.Legs, xr.Soles, xr.Other, xr.OtherSpecify, xr.Macule, xr.Papule, xr.EarlyVesicle, xr.SmallPustule, xr.UmbilicatedPustule, xr.UlceratedLesion, xr.CrustingMature, xr.PartiallyRemoved, xr.PainAtSite, xr.PainScore, xr.LesionsPerianal2, xr.ShadeLocationLesions, xr.EnterBy, xr.EnterOn)
	if err := db.QueryRowContext(ctx, sqlstr, xr.EncounterID, xr.LesionsBody, xr.LesionsRightLeg, xr.LesionsRightArm, xr.LesionsLeftLeg, xr.LesionsLeftArm, xr.LesionsGenitals, xr.LesionsOral, xr.LesionsPerianal1, xr.Face, xr.Nares, xr.Mouth, xr.Chest, xr.Abdomen, xr.Back, xr.Perianal, xr.Genitals, xr.Palms, xr.Arms, xr.Forearms, xr.Thighs, xr.Legs, xr.Soles, xr.Other, xr.OtherSpecify, xr.Macule, xr.Papule, xr.EarlyVesicle, xr.SmallPustule, xr.UmbilicatedPustule, xr.UlceratedLesion, xr.CrustingMature, xr.PartiallyRemoved, xr.PainAtSite, xr.PainScore, xr.LesionsPerianal2, xr.ShadeLocationLesions, xr.EnterBy, xr.EnterOn).Scan(&xr.ID); err != nil {
		return logerror(err)
	}
	// set exists
	xr._exists = true
	return nil
}

// Update updates a [XRush] in the database.
func (xr *XRush) Update(ctx context.Context, db DB) error {
	switch {
	case !xr._exists: // doesn't exist
		return logerror(&ErrUpdateFailed{ErrDoesNotExist})
	case xr._deleted: // deleted
		return logerror(&ErrUpdateFailed{ErrMarkedForDeletion})
	}
	// update with composite primary key
	const sqlstr = `UPDATE public.x_rush SET ` +
		`encounter_id = $1, lesions_body = $2, lesions_right_leg = $3, lesions_right_arm = $4, lesions_left_leg = $5, lesions_left_arm = $6, lesions_genitals = $7, lesions_oral = $8, lesions_perianal_1 = $9, face = $10, nares = $11, mouth = $12, chest = $13, abdomen = $14, back = $15, perianal = $16, genitals = $17, palms = $18, arms = $19, forearms = $20, thighs = $21, legs = $22, soles = $23, other = $24, other_specify = $25, macule = $26, papule = $27, early_vesicle = $28, small_pustule = $29, umbilicated_pustule = $30, ulcerated_lesion = $31, crusting_mature = $32, partially_removed = $33, pain_at_site = $34, pain_score = $35, lesions_perianal_2 = $36, shade_location_lesions = $37, enter_by = $38, enter_on = $39 ` +
		`WHERE id = $40`
	// run
	logf(sqlstr, xr.EncounterID, xr.LesionsBody, xr.LesionsRightLeg, xr.LesionsRightArm, xr.LesionsLeftLeg, xr.LesionsLeftArm, xr.LesionsGenitals, xr.LesionsOral, xr.LesionsPerianal1, xr.Face, xr.Nares, xr.Mouth, xr.Chest, xr.Abdomen, xr.Back, xr.Perianal, xr.Genitals, xr.Palms, xr.Arms, xr.Forearms, xr.Thighs, xr.Legs, xr.Soles, xr.Other, xr.OtherSpecify, xr.Macule, xr.Papule, xr.EarlyVesicle, xr.SmallPustule, xr.UmbilicatedPustule, xr.UlceratedLesion, xr.CrustingMature, xr.PartiallyRemoved, xr.PainAtSite, xr.PainScore, xr.LesionsPerianal2, xr.ShadeLocationLesions, xr.EnterBy, xr.EnterOn, xr.ID)
	if _, err := db.ExecContext(ctx, sqlstr, xr.EncounterID, xr.LesionsBody, xr.LesionsRightLeg, xr.LesionsRightArm, xr.LesionsLeftLeg, xr.LesionsLeftArm, xr.LesionsGenitals, xr.LesionsOral, xr.LesionsPerianal1, xr.Face, xr.Nares, xr.Mouth, xr.Chest, xr.Abdomen, xr.Back, xr.Perianal, xr.Genitals, xr.Palms, xr.Arms, xr.Forearms, xr.Thighs, xr.Legs, xr.Soles, xr.Other, xr.OtherSpecify, xr.Macule, xr.Papule, xr.EarlyVesicle, xr.SmallPustule, xr.UmbilicatedPustule, xr.UlceratedLesion, xr.CrustingMature, xr.PartiallyRemoved, xr.PainAtSite, xr.PainScore, xr.LesionsPerianal2, xr.ShadeLocationLesions, xr.EnterBy, xr.EnterOn, xr.ID); err != nil {
		return logerror(err)
	}
	return nil
}

// Save saves the [XRush] to the database.
func (xr *XRush) Save(ctx context.Context, db DB) error {
	if xr.Exists() {
		return xr.Update(ctx, db)
	}
	return xr.Insert(ctx, db)
}

// Upsert performs an upsert for [XRush].
func (xr *XRush) Upsert(ctx context.Context, db DB) error {
	switch {
	case xr._deleted: // deleted
		return logerror(&ErrUpsertFailed{ErrMarkedForDeletion})
	}
	// upsert
	const sqlstr = `INSERT INTO public.x_rush (` +
		`id, encounter_id, lesions_body, lesions_right_leg, lesions_right_arm, lesions_left_leg, lesions_left_arm, lesions_genitals, lesions_oral, lesions_perianal_1, face, nares, mouth, chest, abdomen, back, perianal, genitals, palms, arms, forearms, thighs, legs, soles, other, other_specify, macule, papule, early_vesicle, small_pustule, umbilicated_pustule, ulcerated_lesion, crusting_mature, partially_removed, pain_at_site, pain_score, lesions_perianal_2, shade_location_lesions, enter_by, enter_on` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25, $26, $27, $28, $29, $30, $31, $32, $33, $34, $35, $36, $37, $38, $39, $40` +
		`)` +
		` ON CONFLICT (id) DO ` +
		`UPDATE SET ` +
		`encounter_id = EXCLUDED.encounter_id, lesions_body = EXCLUDED.lesions_body, lesions_right_leg = EXCLUDED.lesions_right_leg, lesions_right_arm = EXCLUDED.lesions_right_arm, lesions_left_leg = EXCLUDED.lesions_left_leg, lesions_left_arm = EXCLUDED.lesions_left_arm, lesions_genitals = EXCLUDED.lesions_genitals, lesions_oral = EXCLUDED.lesions_oral, lesions_perianal_1 = EXCLUDED.lesions_perianal_1, face = EXCLUDED.face, nares = EXCLUDED.nares, mouth = EXCLUDED.mouth, chest = EXCLUDED.chest, abdomen = EXCLUDED.abdomen, back = EXCLUDED.back, perianal = EXCLUDED.perianal, genitals = EXCLUDED.genitals, palms = EXCLUDED.palms, arms = EXCLUDED.arms, forearms = EXCLUDED.forearms, thighs = EXCLUDED.thighs, legs = EXCLUDED.legs, soles = EXCLUDED.soles, other = EXCLUDED.other, other_specify = EXCLUDED.other_specify, macule = EXCLUDED.macule, papule = EXCLUDED.papule, early_vesicle = EXCLUDED.early_vesicle, small_pustule = EXCLUDED.small_pustule, umbilicated_pustule = EXCLUDED.umbilicated_pustule, ulcerated_lesion = EXCLUDED.ulcerated_lesion, crusting_mature = EXCLUDED.crusting_mature, partially_removed = EXCLUDED.partially_removed, pain_at_site = EXCLUDED.pain_at_site, pain_score = EXCLUDED.pain_score, lesions_perianal_2 = EXCLUDED.lesions_perianal_2, shade_location_lesions = EXCLUDED.shade_location_lesions, enter_by = EXCLUDED.enter_by, enter_on = EXCLUDED.enter_on `
	// run
	logf(sqlstr, xr.ID, xr.EncounterID, xr.LesionsBody, xr.LesionsRightLeg, xr.LesionsRightArm, xr.LesionsLeftLeg, xr.LesionsLeftArm, xr.LesionsGenitals, xr.LesionsOral, xr.LesionsPerianal1, xr.Face, xr.Nares, xr.Mouth, xr.Chest, xr.Abdomen, xr.Back, xr.Perianal, xr.Genitals, xr.Palms, xr.Arms, xr.Forearms, xr.Thighs, xr.Legs, xr.Soles, xr.Other, xr.OtherSpecify, xr.Macule, xr.Papule, xr.EarlyVesicle, xr.SmallPustule, xr.UmbilicatedPustule, xr.UlceratedLesion, xr.CrustingMature, xr.PartiallyRemoved, xr.PainAtSite, xr.PainScore, xr.LesionsPerianal2, xr.ShadeLocationLesions, xr.EnterBy, xr.EnterOn)
	if _, err := db.ExecContext(ctx, sqlstr, xr.ID, xr.EncounterID, xr.LesionsBody, xr.LesionsRightLeg, xr.LesionsRightArm, xr.LesionsLeftLeg, xr.LesionsLeftArm, xr.LesionsGenitals, xr.LesionsOral, xr.LesionsPerianal1, xr.Face, xr.Nares, xr.Mouth, xr.Chest, xr.Abdomen, xr.Back, xr.Perianal, xr.Genitals, xr.Palms, xr.Arms, xr.Forearms, xr.Thighs, xr.Legs, xr.Soles, xr.Other, xr.OtherSpecify, xr.Macule, xr.Papule, xr.EarlyVesicle, xr.SmallPustule, xr.UmbilicatedPustule, xr.UlceratedLesion, xr.CrustingMature, xr.PartiallyRemoved, xr.PainAtSite, xr.PainScore, xr.LesionsPerianal2, xr.ShadeLocationLesions, xr.EnterBy, xr.EnterOn); err != nil {
		return logerror(err)
	}
	// set exists
	xr._exists = true
	return nil
}

// Delete deletes the [XRush] from the database.
func (xr *XRush) Delete(ctx context.Context, db DB) error {
	switch {
	case !xr._exists: // doesn't exist
		return nil
	case xr._deleted: // deleted
		return nil
	}
	// delete with single primary key
	const sqlstr = `DELETE FROM public.x_rush ` +
		`WHERE id = $1`
	// run
	logf(sqlstr, xr.ID)
	if _, err := db.ExecContext(ctx, sqlstr, xr.ID); err != nil {
		return logerror(err)
	}
	// set deleted
	xr._deleted = true
	return nil
}

// XRushByID retrieves a row from 'public.x_rush' as a [XRush].
//
// Generated from index 'rush_pkey'.
func XRushByID(ctx context.Context, db DB, id int) (*XRush, error) {
	// query
	const sqlstr = `SELECT ` +
		`id, encounter_id, lesions_body, lesions_right_leg, lesions_right_arm, lesions_left_leg, lesions_left_arm, lesions_genitals, lesions_oral, lesions_perianal_1, face, nares, mouth, chest, abdomen, back, perianal, genitals, palms, arms, forearms, thighs, legs, soles, other, other_specify, macule, papule, early_vesicle, small_pustule, umbilicated_pustule, ulcerated_lesion, crusting_mature, partially_removed, pain_at_site, pain_score, lesions_perianal_2, shade_location_lesions, enter_by, enter_on ` +
		`FROM public.x_rush ` +
		`WHERE id = $1`
	// run
	logf(sqlstr, id)
	xr := XRush{
		_exists: true,
	}
	if err := db.QueryRowContext(ctx, sqlstr, id).Scan(&xr.ID, &xr.EncounterID, &xr.LesionsBody, &xr.LesionsRightLeg, &xr.LesionsRightArm, &xr.LesionsLeftLeg, &xr.LesionsLeftArm, &xr.LesionsGenitals, &xr.LesionsOral, &xr.LesionsPerianal1, &xr.Face, &xr.Nares, &xr.Mouth, &xr.Chest, &xr.Abdomen, &xr.Back, &xr.Perianal, &xr.Genitals, &xr.Palms, &xr.Arms, &xr.Forearms, &xr.Thighs, &xr.Legs, &xr.Soles, &xr.Other, &xr.OtherSpecify, &xr.Macule, &xr.Papule, &xr.EarlyVesicle, &xr.SmallPustule, &xr.UmbilicatedPustule, &xr.UlceratedLesion, &xr.CrustingMature, &xr.PartiallyRemoved, &xr.PainAtSite, &xr.PainScore, &xr.LesionsPerianal2, &xr.ShadeLocationLesions, &xr.EnterBy, &xr.EnterOn); err != nil {
		return nil, logerror(err)
	}
	return &xr, nil
}
