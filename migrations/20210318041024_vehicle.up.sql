CREATE TABLE IF NOT EXISTS vehicle
(
    id              serial PRIMARY KEY,
    vehicle_type    varchar(20) NOT NULL,
    registration_no varchar(40) NOT NULL
);
