CREATE TABLE IF NOT EXISTS vehicle
(
    id              serial PRIMARY KEY,
    model           varchar(20) NOT NULL,
    registration_no varchar(40) NOT NULL UNIQUE
);
