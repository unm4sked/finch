CREATE TABLE IF NOT EXISTS configurations (
    id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
    description TEXT,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);