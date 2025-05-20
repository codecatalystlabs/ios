-- Add outbreak_id column to encounter table
ALTER TABLE public.encounter ADD COLUMN outbreak_id INTEGER REFERENCES public.outbreaks(id);

-- Add outbreak_id column to clinical table
ALTER TABLE public.clinical ADD COLUMN outbreak_id INTEGER REFERENCES public.outbreaks(id);

-- Add outbreak_id column to vitals table
ALTER TABLE public.vitals ADD COLUMN outbreak_id INTEGER REFERENCES public.outbreaks(id);

-- Add outbreak_id column to treatment table
ALTER TABLE public.treatment ADD COLUMN outbreak_id INTEGER REFERENCES public.outbreaks(id);

-- Add outbreak_id column to lab table
ALTER TABLE public.lab ADD COLUMN outbreak_id INTEGER REFERENCES public.outbreaks(id);

-- Add outbreak_id column to rush table
ALTER TABLE public.rush ADD COLUMN outbreak_id INTEGER REFERENCES public.outbreaks(id);

-- Add outbreak_id column to symptoms table
ALTER TABLE public.symptoms ADD COLUMN outbreak_id INTEGER REFERENCES public.outbreaks(id);

-- Add outbreak_id column to comorbidities table
ALTER TABLE public.comorbidities ADD COLUMN outbreak_id INTEGER REFERENCES public.outbreaks(id);

-- Add outbreak_id column to discharge table
ALTER TABLE public.discharge ADD COLUMN outbreak_id INTEGER REFERENCES public.outbreaks(id);

-- Add outbreak_id column to enrollment table
ALTER TABLE public.enrollment ADD COLUMN outbreak_id INTEGER REFERENCES public.outbreaks(id);

-- Create indexes for better query performance
CREATE INDEX IF NOT EXISTS idx_encounter_outbreak ON public.encounter(outbreak_id);
CREATE INDEX IF NOT EXISTS idx_clinical_outbreak ON public.clinical(outbreak_id);
CREATE INDEX IF NOT EXISTS idx_vitals_outbreak ON public.vitals(outbreak_id);
CREATE INDEX IF NOT EXISTS idx_treatment_outbreak ON public.treatment(outbreak_id);
CREATE INDEX IF NOT EXISTS idx_lab_outbreak ON public.lab(outbreak_id);
CREATE INDEX IF NOT EXISTS idx_rush_outbreak ON public.rush(outbreak_id);
CREATE INDEX IF NOT EXISTS idx_symptoms_outbreak ON public.symptoms(outbreak_id);
CREATE INDEX IF NOT EXISTS idx_comorbidities_outbreak ON public.comorbidities(outbreak_id);
CREATE INDEX IF NOT EXISTS idx_discharge_outbreak ON public.discharge(outbreak_id);
CREATE INDEX IF NOT EXISTS idx_enrollment_outbreak ON public.enrollment(outbreak_id); 