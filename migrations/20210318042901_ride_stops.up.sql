CREATE TABLE IF NOT EXISTS ride_stops
(
    id            serial PRIMARY KEY,
    location      int REFERENCES address (id) ON DELETE CASCADE NOT NULL,
    prev_location int REFERENCES address (id) ON DELETE CASCADE NOT NULL,
    next_location int REFERENCES address (id) ON DELETE CASCADE NOT NULL,
    created_at    timestamp with time zone                      NOT NULL DEFAULT now(),
    updated_at    timestamp with time zone                      NOT NULL
);
