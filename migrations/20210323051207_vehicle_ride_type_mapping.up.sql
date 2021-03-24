CREATE TABLE IF NOT EXISTS vehicle_ride_type_mapping
(
    id         serial PRIMARY KEY,
    vehicle_id int REFERENCES vehicle (id) ON DELETE CASCADE NOT NULL,
    ride_type  varchar(40)                                   NOT NULL
);

CREATE UNIQUE INDEX unique_idx1 ON vehicle_ride_type_mapping (vehicle_id, ride_type);
