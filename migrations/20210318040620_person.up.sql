CREATE TABLE IF NOT EXISTS person
(
    id         serial PRIMARY KEY,
    first_name varchar(20)              NOT NULL,
    last_name  varchar(20)              NOT NULL,
    phone      varchar(10)              NOT NULL,
    created_at timestamp with time zone NOT NULL DEFAULT now(),
    updated_at timestamp with time zone NOT NULL
);
