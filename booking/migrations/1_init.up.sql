CREATE TABLE booking (
   id BIGSERIAL PRIMARY KEY,
   apartment_id BIGINT NOT NULL,
   date_start TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
   date_end TIMESTAMP,
   price INTEGER,
   customer_id BIGINT NOT NULL,
   status VARCHAR(20) DEFAULT 'inactive'
);
