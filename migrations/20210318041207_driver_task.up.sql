CREATE TABLE IF NOT EXISTS driver_task
(
    id               serial PRIMARY KEY,
    customer_task_id int REFERENCES customer_task (id) ON DELETE CASCADE NOT NULL,
    driver_id        int REFERENCES user (id) ON DELETE CASCADE          NOT NULL,
    payable_amount   float                                               NOT NULL,
    ride_type        varchar(10)                                         NOT NULL,
    status           varchar(10)                                         NOT NULL,
    distance         float                                               NOT NULL,
    rating           int,
    created_at       timestamp with time zone                            NOT NULL DEFAULT now(),
    updated_at       timestamp with time zone                            NOT NULL
);
