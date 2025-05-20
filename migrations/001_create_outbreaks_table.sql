-- Create outbreaks table
CREATE TABLE IF NOT EXISTS public.outbreaks (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    start_date TIMESTAMP NOT NULL,
    end_date TIMESTAMP,
    status VARCHAR(50) NOT NULL DEFAULT 'active',
    enter_on TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    enter_by INTEGER REFERENCES public.users(user_id),
    edit_on TIMESTAMP,
    edit_by INTEGER REFERENCES public.users(user_id)
);

-- Create index on status for faster queries
CREATE INDEX IF NOT EXISTS idx_outbreaks_status ON public.outbreaks(status);

-- Insert default Ebola 2025 outbreak
INSERT INTO public.outbreaks (name, description, start_date, status, enter_on)
VALUES (
    'Ebola 2025',
    'Ebola outbreak in 2025',
    CURRENT_TIMESTAMP,
    'active',
    CURRENT_TIMESTAMP
)
ON CONFLICT (name) DO NOTHING; 