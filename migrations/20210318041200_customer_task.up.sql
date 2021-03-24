CREATE TABLE IF NOT EXISTS customer_task
(
    id             serial PRIMARY KEY,
    customer_id    int REFERENCES users (user_id) ON DELETE CASCADE NOT NULL,
    payable_amount float                                            NOT NULL,
    status         varchar(10)                                      NOT NULL,
    ride_type      varchar(10)                                      NOT NULL,
    created_at     timestamp with time zone                         NOT NULL DEFAULT now(),
    updated_at     timestamp with time zone                         NOT NULL DEFAULT now()
);

CREATE INDEX IF NOT EXISTS customer_task_customer_id_status_idx on customer_task (customer_id, status);
