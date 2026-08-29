-- +goose Up
CREATE TABLE genres ( 
id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
name TEXT
);

CREATE TABLE actors (
id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
name TEXT NOT NULL,
birth_date TEXT
);

CREATE TABLE movies (
id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
title TEXT NOT NULL,
release_year TEXT NOT NULL,
duration INTEGER NOT NULL
);

CREATE TABLE movie_genres (
movie_id INTEGER NOT NULL,
genre_id INTEGER NOT NULL,
PRIMARY KEY (movie_id, genre_id),
FOREIGN KEY (movie_id) REFERENCES movies(id) ON DELETE CASCADE,
FOREIGN KEY (genre_id) REFERENCES genres(id) ON DELETE CASCADE
);

CREATE TABLE movie_actors (
movie_id INTEGER NOT NULL,
actor_id INTEGER NOT NULL,
PRIMARY KEY (movie_id, actor_id),
FOREIGN KEY (movie_id) REFERENCES movies(id) ON DELETE CASCADE,
FOREIGN KEY (actor_id) REFERENCES actors(id) ON DELETE CASCADE
);

-- +goose Down

DROP TABLE movie_actors;
DROP TABLE movie_genres;
DROP TABLE movies;
DROP TABLE actors;
DROP TABLE genres;
