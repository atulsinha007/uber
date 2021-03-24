CREATE TABLE IF NOT EXISTS driver_profile
(
    id           serial PRIMARY KEY,
    driver_id    int REFERENCES users (user_id) ON DELETE CASCADE UNIQUE,
    is_available bool                     NOT NULL DEFAULT true,
    vehicle_id   int                      NOT NULL,
    created_at   timestamp with time zone NOT NULL DEFAULT now(),
    updated_at   timestamp with time zone NOT NULL DEFAULT now()
);

CREATE INDEX IF NOT EXISTS driver_profile_vehicle_id_idx on driver_profile (vehicle_id);
CREATE INDEX IF NOT EXISTS driver_profile_driver_id_idx on driver_profile (driver_id);
