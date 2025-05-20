-- Add missing columns to vhf_clinical_signs table
ALTER TABLE public.vhf_clinical_signs
ADD COLUMN IF NOT EXISTS temp_source VARCHAR(255),
ADD COLUMN IF NOT EXISTS temperature DOUBLE PRECISION; 