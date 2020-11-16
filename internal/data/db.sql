CREATE TABLE users(
    id SERIAL PRIMARY KEY NOT NULL,
    email VARCHAR(254)  NOT NULL,
    apiKey CHAR(64) NOT NULL,
    created_at TIMESTAMPTZ NOT NULL,
    updated_at TIMESTAMPTZ NOT NULL,
    CONSTRAINT email_uniq UNIQUE (email)
);

CREATE INDEX idx_users_apiKey ON users(apiKey);


INSERT INTO users(email, apiKey, created_at, updated_at) values('ask@me.com',
    '61ea4113f56ba6e8cdb313eeb4d953be6ba29796614046c5281d5e7b468a2765',
    'NOW()', 'NOW()');