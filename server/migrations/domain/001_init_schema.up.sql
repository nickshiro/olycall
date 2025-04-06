BEGIN;

CREATE TABLE app_user (
    id UUID PRIMARY KEY,
    email TEXT NOT NULL UNIQUE,
    username TEXT NOT NULL UNIQUE,
    created_at TIMESTAMPTZ,
    updated_at TIMESTAMPTZ
);

CREATE TABLE oauth2_provider (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL UNIQUE
);

CREATE TABLE oauth2_identitiy (
    id SERIAL PRIMARY KEY,
    user_id UUID NOT NULL REFERENCES app_user (id) ON DELETE CASCADE,
    provider_id INTEGER NOT NULL REFERENCES oauth2_provider (id),
    subject TEXT NOT NULL,
    created_at TIMESTAMPTZ,

    UNIQUE (provider_id, subject)
);

COMMIT;
