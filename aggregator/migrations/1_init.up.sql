CREATE TABLE apartment (
   id BIGSERIAL PRIMARY KEY,
   title VARCHAR NOT NULL,
   expenses INTEGER NOT NULL ,
   status VARCHAR(20) DEFAULT 'inactive',
   created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);
CREATE TABLE customers (
   id BIGSERIAL PRIMARY KEY,
   name VARCHAR(255) NOT NULL,
   phone VARCHAR(50) NOT NULL,
   passport VARCHAR,
   created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);
CREATE TABLE booking (
   id BIGSERIAL PRIMARY KEY,
   apartment_id BIGINT REFERENCES apartment(id) ON DELETE CASCADE ,
   date_start TIMESTAMP,
   date_end TIMESTAMP,
   price BIGINT NOT NULL,
   customer_id BIGINT REFERENCES customers(id) ON DELETE CASCADE ,
   status VARCHAR(20) DEFAULT 'pending',
   comment VARCHAR(255),
   date_created TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE UNIQUE INDEX booking_unique_active_apartment
ON booking(apartment_id)
WHERE status = 'active';