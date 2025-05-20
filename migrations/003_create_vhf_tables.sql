-- Create VHF patients table
CREATE TABLE IF NOT EXISTS public.vhf_patients (
    id SERIAL PRIMARY KEY,
    surname VARCHAR(255) NOT NULL,
    other_names VARCHAR(255),
    date_of_birth TIMESTAMP,
    age_years INTEGER,
    age_months INTEGER,
    gender VARCHAR(50),
    patient_phone VARCHAR(50),
    phone_owner VARCHAR(255),
    next_of_kin VARCHAR(255),
    next_of_kin_phone VARCHAR(50),
    status VARCHAR(50),
    date_of_death TIMESTAMP,
    head_of_household VARCHAR(255),
    village_town VARCHAR(255),
    parish VARCHAR(255),
    subcounty VARCHAR(255),
    district VARCHAR(255),
    country_of_residence VARCHAR(255),
    occupation VARCHAR(255),
    ill_village_town VARCHAR(255),
    ill_subcounty VARCHAR(255),
    ill_district VARCHAR(255),
    latitude DOUBLE PRECISION,
    longitude DOUBLE PRECISION,
    date_residing_from TIMESTAMP,
    date_residing_to TIMESTAMP,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- Create VHF clinical signs table
CREATE TABLE IF NOT EXISTS public.vhf_clinical_signs (
    id SERIAL PRIMARY KEY,
    patient_id INTEGER NOT NULL REFERENCES public.vhf_patients(id),
    date_initial_onset TIMESTAMP,
    fever BOOLEAN,
    vomiting BOOLEAN,
    nausea BOOLEAN,
    diarrhea BOOLEAN,
    intense_fatigue_general_weakness BOOLEAN,
    epigastric_pain BOOLEAN,
    lower_abdominal_pain BOOLEAN,
    chest_pain BOOLEAN,
    muscle_pain BOOLEAN,
    joint_pain BOOLEAN,
    headache BOOLEAN,
    cough BOOLEAN,
    difficulty_breathing BOOLEAN,
    difficulty_swallowing BOOLEAN,
    sore_throat BOOLEAN,
    jaundice BOOLEAN,
    conjunctivitis BOOLEAN,
    skin_rash BOOLEAN,
    hiccups BOOLEAN,
    pain_behind_eyes BOOLEAN,
    sensitive_to_light BOOLEAN,
    coma_unconscious BOOLEAN,
    confused_or_disoriented BOOLEAN,
    convulsions BOOLEAN,
    unexplained_bleeding BOOLEAN,
    bleeding_of_the_gums BOOLEAN,
    bleeding_from_injection_site BOOLEAN,
    nose_bleed_epistaxis BOOLEAN,
    bloody_stool BOOLEAN,
    blood_in_vomit BOOLEAN,
    coughing_up_blood_hemoptysis BOOLEAN,
    bleeding_from_vagina BOOLEAN,
    bruising_of_the_skin BOOLEAN,
    blood_in_urine BOOLEAN,
    other_hemorrhagic_symptoms BOOLEAN,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- Create VHF hospitalization table
CREATE TABLE IF NOT EXISTS public.vhf_hospitalization (
    id SERIAL PRIMARY KEY,
    patient_id INTEGER NOT NULL REFERENCES public.vhf_patients(id),
    hospitalized BOOLEAN NOT NULL,
    admission_date TIMESTAMP,
    health_facility_name VARCHAR(255),
    in_isolation BOOLEAN NOT NULL,
    isolation_date TIMESTAMP,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- Create VHF risk factors table
CREATE TABLE IF NOT EXISTS public.vhf_risk_factors (
    id SERIAL PRIMARY KEY,
    patient_id INTEGER NOT NULL REFERENCES public.vhf_patients(id),
    contact_with_case BOOLEAN,
    contact_name VARCHAR(255),
    contact_relation VARCHAR(255),
    contact_dates VARCHAR(255),
    contact_village VARCHAR(255),
    contact_district VARCHAR(255),
    contact_status VARCHAR(50),
    contact_death_date TIMESTAMP,
    contact_types VARCHAR(255),
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- Create VHF laboratory table
CREATE TABLE IF NOT EXISTS public.vhf_laboratory (
    id SERIAL PRIMARY KEY,
    patient_id INTEGER NOT NULL REFERENCES public.vhf_patients(id),
    sample_collection_date TIMESTAMP,
    sample_collection_time VARCHAR(50),
    sample_type VARCHAR(255),
    other_sample_type VARCHAR(255),
    requested_test VARCHAR(255),
    serology VARCHAR(255),
    malaria_rdt VARCHAR(255),
    hiv_rdt VARCHAR(255),
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- Create VHF investigator table
CREATE TABLE IF NOT EXISTS public.vhf_investigator (
    id SERIAL PRIMARY KEY,
    patient_id INTEGER NOT NULL REFERENCES public.vhf_patients(id),
    investigator_name VARCHAR(255),
    phone VARCHAR(50),
    email VARCHAR(255),
    position VARCHAR(255),
    district VARCHAR(255),
    health_facility VARCHAR(255),
    information_source VARCHAR(255),
    proxy_name VARCHAR(255),
    proxy_relation VARCHAR(255),
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- Create indexes for better query performance
CREATE INDEX IF NOT EXISTS idx_vhf_clinical_signs_patient ON public.vhf_clinical_signs(patient_id);
CREATE INDEX IF NOT EXISTS idx_vhf_hospitalization_patient ON public.vhf_hospitalization(patient_id);
CREATE INDEX IF NOT EXISTS idx_vhf_risk_factors_patient ON public.vhf_risk_factors(patient_id);
CREATE INDEX IF NOT EXISTS idx_vhf_laboratory_patient ON public.vhf_laboratory(patient_id);
CREATE INDEX IF NOT EXISTS idx_vhf_investigator_patient ON public.vhf_investigator(patient_id); 