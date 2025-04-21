CREATE TABLE products (
    id UUID PRIMARY KEY,
    reception_time TIMESTAMP NOT NULL,
    type TEXT NOT NULL,
    reception_id UUID NOT NULL,
    previous_product_id UUID
);
