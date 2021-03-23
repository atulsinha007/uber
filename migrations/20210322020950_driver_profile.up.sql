CREATE TABLE IF NOT EXISTS driver_profile
(
    id           serial PRIMARY KEY,
    driver_id    int REFERENCES users (user_id) ON DELETE CASCADE,
    is_available bool                     NOT NULL DEFAULT true,
    vehicle_id   int                      NOT NULL,
    created_at   timestamp with time zone NOT NULL DEFAULT now(),
    updated_at   timestamp with time zone NOT NULL DEFAULT now()
);
