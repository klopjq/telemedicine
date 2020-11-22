CREATE TABLE IF NOT EXISTS users(
    id SERIAL PRIMARY KEY NOT NULL,
    email VARCHAR(254)  NOT NULL,
    apiKey CHAR(64) NOT NULL,
    created_at TIMESTAMPTZ NOT NULL,
    updated_at TIMESTAMPTZ NOT NULL,
    CONSTRAINT email_uniq UNIQUE (email)
);

CREATE INDEX IF NOT EXISTS idx_users_apiKey ON users(apiKey);


INSERT INTO users(email, apiKey, created_at, updated_at) values('ask@me.com',
    '61ea4113f56ba6e8cdb313eeb4d953be6ba29796614046c5281d5e7b468a2765',
    'NOW()', 'NOW()')
ON CONFLICT DO NOTHING	;

CREATE TABLE IF NOT EXISTS users_medical_data(
    id INTEGER PRIMARY KEY NOT NULL,
    data jsonb
);

CREATE INDEX IF NOT EXISTS idx_users_medical_data ON users_medical_data USING GIN (data jsonb_path_ops);

INSERT INTO users_medical_data(id, data)
values(1, jsonb_build_array(jsonb_build_object('ts',NOW(),'bodyTemperature',36.7,'heartRate',56)))
ON CONFLICT (id) DO
UPDATE SET id=excluded.id, data=users_medical_data.data||excluded.data;

INSERT INTO users_medical_data(id, data)
values(1, jsonb_build_array(jsonb_build_object(
	'ts',NOW()+(RANDOM() * (NOW()+'7 days' - NOW())) + '2 days',
	'bodyTemperature',36.7+(FLOOR(RANDOM()*10+1)::INTEGER),
	'heartRate',56+(FLOOR(RANDOM()*10+1)::INTEGER))))
ON CONFLICT (id) DO
UPDATE SET id=excluded.id, data=users_medical_data.data||excluded.data;