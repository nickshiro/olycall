BEGIN;

CREATE TABLE app_user (
    id UUID PRIMARY KEY,
    email TEXT NOT NULL UNIQUE,
    username TEXT NOT NULL UNIQUE,
    name TEXT NOT NULL,
    avatar_url TEXT,
    created_at TIMESTAMPTZ NOT NULL
);

-- CREATE TABLE oauth2_identitiy (
--     id SERIAL PRIMARY KEY,
--     user_id UUID NOT NULL REFERENCES app_user (id) ON DELETE CASCADE,
--     provider TEXT NOT NULL,
--     subject TEXT NOT NULL,
--     created_at TIMESTAMPTZ,

--     UNIQUE (provider, subject)
-- );

CREATE TABLE chat (
    id UUID PRIMARY KEY,
    type TEXT NOT NULL CHECK (type IN ('group', 'direct'))
);

CREATE TABLE user_chat (
    user_id UUID NOT NULL REFERENCES app_user (id) ON DELETE CASCADE,
    chat_id UUID NOT NULL REFERENCES chat (id) ON DELETE CASCADE,
    pinned BOOLEAN NOT NULL,
    muted BOOLEAN NOT NULL,

    PRIMARY KEY (user_id, chat_id)
);

CREATE TABLE group_chat (
    chat_id UUID PRIMARY KEY REFERENCES chat (id) ON DELETE CASCADE,
    name TEXT NOT NULL,
    avatar_url TEXT,
    description TEXT
);

CREATE TABLE group_chat_member (
    group_chat_id UUID NOT NULL REFERENCES group_chat (
        chat_id
    ) ON DELETE CASCADE,
    user_id UUID NOT NULL REFERENCES app_user (id) ON DELETE CASCADE,
    joined_at TIMESTAMPTZ NOT NULL,

    PRIMARY KEY (group_chat_id, user_id)
);

CREATE TABLE direct_chat (
    chat_id UUID PRIMARY KEY REFERENCES chat (id) ON DELETE CASCADE,
    user1_id UUID NOT NULL REFERENCES app_user (id) ON DELETE CASCADE,
    user2_id UUID NOT NULL REFERENCES app_user (id) ON DELETE CASCADE,

    UNIQUE (user1_id, user2_id),
    CHECK (user1_id <> user2_id)
);

CREATE TABLE message (
    id UUID PRIMARY KEY,
    sender_id UUID NOT NULL REFERENCES app_user (id),
    chat_id UUID NOT NULL REFERENCES chat (id),
    created_at TIMESTAMPTZ NOT NULL,
    updated_at TIMESTAMPTZ,
    reply_to_id UUID REFERENCES message (id),
    forwarded_message_id UUID REFERENCES message (id),
    content TEXT
);

CREATE TABLE file (
    id UUID PRIMARY KEY,
    name TEXT NOT NULL,
    mime_type TEXT NOT NULL,
    size_bytes BIGINT NOT NULL,
    upload_timestamp TIMESTAMPTZ NOT NULL
);

CREATE TABLE message_file (
    message_id UUID NOT NULL REFERENCES message (id) ON DELETE CASCADE,
    file_id UUID NOT NULL REFERENCES file (id) ON DELETE CASCADE,

    PRIMARY KEY (message_id, file_id)
);

COMMIT;
