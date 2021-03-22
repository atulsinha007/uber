CREATE TABLE IF NOT EXISTS driver_profile
(
    id            serial PRIMARY KEY,
    driver_id     int REFERENCES users (user_id) ON DELETE CASCADE,
    vehicle_id    varchar(20)              NOT NULL,
    created_at    timestamp with time zone NOT NULL DEFAULT now(),
    updated_at    timestamp with time zone NOT NULL DEFAULT now()
);
