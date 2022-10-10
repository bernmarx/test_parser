CREATE TABLE IF NOT EXISTS tasks (
    id SERIAL NOT NULL PRIMARY KEY,
    question VARCHAR(256) NOT NULL UNIQUE,
    answer VARCHAR(256) NOT NULL
);
CREATE TABLE IF NOT EXISTS options (
    id SERIAL NOT NULL PRIMARY KEY,
    question_id INTEGER REFERENCES tasks (id) ON DELETE CASCADE,
    option VARCHAR(256) NOT NULL
);