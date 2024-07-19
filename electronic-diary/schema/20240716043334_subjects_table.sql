-- +goose Up
CREATE TABLE subjects (
                          id SERIAL PRIMARY KEY,
                          user_id INT REFERENCES students(id) on delete cascade not null,
                          subject VARCHAR(50),
                          mark int,
                          date TIMESTAMP
);

-- +goose Down
DROP TABLE subjects;
