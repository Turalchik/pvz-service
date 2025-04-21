CREATE TABLE receptions (
    id UUID PRIMARY KEY,
    start_time TIMESTAMP NOT NULL,
    end_time TIMESTAMP,
    pvz_id UUID NOT NULL,
    status TEXT NOT NULL,
    last_product_id UUID,

    CONSTRAINT fk_receptions_pvz
        FOREIGN KEY (pvz_id)
            REFERENCES pvzs(id)
            ON DELETE CASCADE
);
