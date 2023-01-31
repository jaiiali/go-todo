CREATE TABLE IF NOT EXISTS todos(
    id VARCHAR(26),
    title VARCHAR(255) NOT NULL,
    description TEXt NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL
    PRIMARY KEY(id)
);

CREATE INDEX todos_created_at_index ON todos (created_at DESC);
