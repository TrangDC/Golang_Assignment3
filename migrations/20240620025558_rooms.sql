CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE rooms (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(255) NOT NULL,
    color VARCHAR(7) NOT NULL,
    office_id UUID NOT NULL,
    description VARCHAR,
    image_url VARCHAR(255)
);