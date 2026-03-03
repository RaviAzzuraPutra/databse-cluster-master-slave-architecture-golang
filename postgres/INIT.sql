CREATE EXTENSION IF NOT EXISTS "pgcrypto";

CREATE TABLE cases (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    case_number VARCHAR(20) NOT NULL UNIQUE,
    case_title VARCHAR(150) NOT NULL,
    case_description TEXT,
    incident_date DATE NOT NULL,
    location VARCHAR(150) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
);

CREATE TABLE suspects (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    case_id UUID NOT NULL,
    id_card_number VARCHAR(16) UNIQUE,
    full_name VARCHAR(100) NOT NULL,
    address TEXT,
    alibi TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT fk_suspects_case FOREIGN KEY (case_id) 
        REFERENCES cases(id) ON DELETE CASCADE,
);

CREATE INDEX idx_suspects_case_id ON suspects(case_id);