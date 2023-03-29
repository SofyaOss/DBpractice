DROP TABLE IF EXISTS articles, authors;

CREATE TABLE authors (
                         id SERIAL PRIMARY KEY,
                         name TEXT NOT NULL
);

CREATE TABLE articles (
                       id SERIAL PRIMARY KEY,
                       author_id INTEGER REFERENCES authors(id) NOT NULL,
                       title TEXT NOT NULL,
                       content TEXT NOT NULL,
                       created_at BIGINT NOT NULL
);

INSERT INTO authors (name) VALUES ('автор1'),('автор2'),('автор3');
INSERT INTO articles (author_id, title, content, created_at)
VALUES(0, 'статья1', 'содержание1', 0),( 0, 'статья 2', 'содержание2', 10);