CREATE TABLE IF NOT EXISTS address
(
    id          serial PRIMARY KEY,
    lat         float                    NOT NULL,
    lng         float                    NOT NULL,
    house_name  varchar(20)              NOT NULL,
    street_name varchar(20)              NOT NULL,
    landmark    varchar(20),
    city        varchar(20)              NOT NULL,
    country     varchar(20)              NOT NULL,
    created_at  timestamp with time zone NOT NULL DEFAULT now(),
    updated_at  timestamp with time zone NOT NULL DEFAULT now()
);
