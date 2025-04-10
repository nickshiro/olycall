BEGIN;

CREATE TABLE app_user (
    id UUID PRIMARY KEY,
    email TEXT NOT NULL UNIQUE,
    username TEXT NOT NULL UNIQUE,
    created_at TIMESTAMPTZ
);

-- CREATE TABLE oauth2_identitiy (
--     id SERIAL PRIMARY KEY,
--     user_id UUID NOT NULL REFERENCES app_user (id) ON DELETE CASCADE,
--     provider TEXT NOT NULL,
--     subject TEXT NOT NULL,
--     created_at TIMESTAMPTZ,

--     UNIQUE (provider, subject)
-- );

COMMIT;
