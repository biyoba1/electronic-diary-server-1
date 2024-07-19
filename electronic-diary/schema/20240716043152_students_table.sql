-- +goose Up
CREATE TABLE students (
                          id SERIAL PRIMARY KEY,
                          name VARCHAR(20),
                          last_name VARCHAR(20),
                          patronymic VARCHAR(20)
);

-- +goose Down
DROP TABLE students;

