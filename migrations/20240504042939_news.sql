CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TYPE news_status_enum AS ENUM ('draft', 'published', 'hidden');

CREATE TABLE news (
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4 (),
  title VARCHAR(255) NOT NULL,
  status news_status_enum DEFAULT 'draft',
  description VARCHAR(512),
  content TEXT NOT NULL,
  author_id UUID NOT NULL References users(id) ON DELETE CASCADE,
  slug VARCHAR(255) NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  deleted_at TIMESTAMP
);
