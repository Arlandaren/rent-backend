CREATE TABLE booking (
   id BIGSERIAL PRIMARY KEY,
   apartment_id BIGINT NOT NULL,
   date_start TIMESTAMP,
   date_end TIMESTAMP,
   price BIGINT NOT NULL,
   customer_id BIGINT NOT NULL,
   status VARCHAR(20) DEFAULT 'pending',
   comment VARCHAR(255),
   date_created TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
