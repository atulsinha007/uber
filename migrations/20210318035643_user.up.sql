CREATE TABLE IF NOT EXISTS users
(
    user_id     serial PRIMARY KEY,
    user_type   varchar(10)              NOT NULL,
    first_name varchar(20)              NOT NULL,
    last_name  varchar(20)              NOT NULL,
    phone      varchar(10)              NOT NULL,
    current_lat float,
    current_lng float,
    created_at  timestamp with time zone NOT NULL DEFAULT now(),
    updated_at  timestamp with time zone NOT NULL DEFAULT now()
);
