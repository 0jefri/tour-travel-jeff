create database travel_db;

CREATE TABLE Tour (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    place_id INTEGER NOT NULL,
    date DATE,
    FOREIGN KEY (place_id) REFERENCES Place(id)
);

CREATE TABLE Place (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    photo VARCHAR(255),
    price FLOAT
);

CREATE TABLE Transaction (
    id SERIAL PRIMARY KEY,
    tour_id INTEGER NOT NULL,
    status VARCHAR(50),
    FOREIGN KEY (tour_id) REFERENCES Tour(id)
);

CREATE TABLE Review (
    id SERIAL PRIMARY KEY,
    transaction_id INTEGER NOT NULL,
    rating FLOAT CHECK (rating >= 0 AND rating <= 5),
    FOREIGN KEY (transaction_id) REFERENCES Transaction(id)
);