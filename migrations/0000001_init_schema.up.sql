CREATE TABLE IF NOT EXISTS genres (

    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL UNIQUE

);

CREATE TABLE IF NOT EXISTS movies (

    id INTEGER PRIMARY KEY AUTOINCREMENT,
    title TEXT NOT NULL,
    release_year INTEGER NOT NULL,
    duration INTEGER NOT NULL
    
);

CREATE TABLE IF NOT EXISTS actors (
    
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    birth_date TEXT NOT NULL

);

CREATE TABLE IF NOT EXISTS movie_genres (
    
    movie_id INTEGER NOT NULL,
    genre_id INTEGER NOT NULL,
    PRIMARY KEY (movie_id, genre_id),
    FOREIGN KEY (movie_id) REFERENCES movies(id),
    FOREIGN KEY (genre_id) REFERENCES genres(id)

);


CREATE TABLE IF NOT EXISTS movie_actors (
    
    movie_id INTEGER NOT NULL,
    actor_id INTEGER NOT NULL,
    PRIMARY KEY (movie_id, actor_id),
    FOREIGN KEY (movie_id) REFERENCES movies(id),
    FOREIGN KEY (movie_id) REFERENCES actors(id)

);