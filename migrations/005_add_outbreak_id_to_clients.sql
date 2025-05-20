-- Add outbreak_id column to clients table
ALTER TABLE public.clients
ADD COLUMN IF NOT EXISTS outbreak_id INTEGER REFERENCES public.outbreaks(id); 