CREATE TABLE IF NOT EXISTS tenants(
   id serial PRIMARY KEY,
   name VARCHAR (250) NOT NULL,
   active BOOLEAN NOT NULL DEFAULT TRUE,
   created_at timestamp with time zone NOT NULL DEFAULT (current_timestamp AT TIME ZONE 'UTC'),
   update_at timestamp with time zone,
   expires_in JSONB
);
