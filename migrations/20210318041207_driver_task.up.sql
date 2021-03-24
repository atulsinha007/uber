CREATE TABLE IF NOT EXISTS driver_task
(
    id               serial PRIMARY KEY,
    customer_task_id int REFERENCES customer_task (id) ON DELETE CASCADE NOT NULL,
    driver_id        int REFERENCES users (user_id) ON DELETE CASCADE    NOT NULL,
    payable_amount   float                                               NOT NULL,
    ride_type        varchar(10)                                         NOT NULL,
    status           varchar(10)                                         NOT NULL,
    distance         float                                               NOT NULL,
    rating           int,
    created_at       timestamp with time zone                            NOT NULL DEFAULT now(),
    updated_at       timestamp with time zone                            NOT NULL DEFAULT now()
);

CREATE INDEX IF NOT EXISTS driver_task_driver_id_status_idx on driver_task (driver_id, status);
CREATE INDEX IF NOT EXISTS driver_task_customer_task_id_idx on driver_task (customer_task_id);