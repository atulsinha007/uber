CREATE TABLE IF NOT EXISTS user
(
    user_id    serial PRIMARY KEY,
    user_type  varchar(10)              NOT NULL,
    person_id  int REFERENCES person (id) ON DELETE CASCADE,
    created_at timestamp with time zone NOT NULL DEFAULT now(),
    updated_at timestamp with time zone NOT NULL
);
