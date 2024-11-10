CREATE TABLE apartment (
   id BIGINT PRIMARY KEY,
   title VARCHAR NOT NULL,
   expenses INTEGER NOT NULL ,
   status VARCHAR(20),
   created_at TIMESTAMP
);
CREATE TABLE customers (
   id BIGINT PRIMARY KEY,
   name VARCHAR(255) NOT NULL,
   phone VARCHAR(50) NOT NULL,
   passport VARCHAR,
   created_at TIMESTAMP
);
CREATE TABLE booking (
     id BIGINT PRIMARY KEY,
     apartment_id BIGINT REFERENCES apartment(id) ON DELETE CASCADE,
     date_start TIMESTAMP,
     date_end TIMESTAMP,
     price BIGINT NOT NULL,
     customer_id BIGINT REFERENCES customers(id) ON DELETE CASCADE,
     status VARCHAR(20),
     comment VARCHAR(255),
     date_created TIMESTAMP,
     active_apartment_id BIGINT GENERATED ALWAYS AS (
         CASE WHEN status = 'active' OR status = 'pending' THEN apartment_id ELSE NULL END
         ) STORED
);

CREATE UNIQUE INDEX booking_unique_active_apartment
ON booking(active_apartment_id);