-- +goose Up
-- Disable foreign key checks during batch insert
PRAGMA foreign_keys = OFF;

-- Clean existing data
DELETE FROM movie_actors;
DELETE FROM movie_genres;
DELETE FROM actors;
DELETE FROM movies;
DELETE FROM genres;

-- Re-enable foreign key checks
PRAGMA foreign_keys = ON;

INSERT INTO genres (id, name) VALUES
(1, 'Action'),
(2, 'Sci-Fi'),
(3, 'Drama'),
(4, 'Thriller'),
(5, 'Crime');


INSERT INTO actors (id, name, birth_date) VALUES
(1,  'Leonardo DiCaprio',  '1974-11-11'),
(2,  'Elliot Page',         '1987-02-21'),
(3,  'Tom Hardy',          '1977-09-15'),
(4,  'Keanu Reeves',       '1964-09-02'),
(5,  'Laurence Fishburne', '1961-07-30'),
(6,  'Carrie-Anne Moss',   '1967-08-21'),
(7,  'Christian Bale',     '1974-01-30'),
(8,  'Heath Ledger',       '1979-04-04'),
(9,  'Morgan Freeman',     '1937-06-01'),
(10, 'Brad Pitt',          '1963-12-18'),
(11, 'Edward Norton',      '1969-08-18'),
(12, 'Matthew McConaughey','1969-11-04'),
(13, 'Anne Hathaway',      '1982-11-12'),
(14, 'Al Pacino',          '1940-04-25'),
(15, 'Robert De Niro',     '1943-08-17');


INSERT INTO movies (id, title, release_year, duration) VALUES
-- 1990s (6 Movies)
(1,  'The Matrix',         1999, 136),
(2,  'Fight Club',          1999, 139),
(3,  'Se7en',               1995, 127),
(4,  'Heat',                1995, 170),
(5,  'The Devil''s Advocate', 1997, 144),
(6,  'The Matrix Reloaded', 2003, 138), -- Note: 2003, grouping below

-- 2000s (7 Movies)
(7,  'The Dark Knight',     2008, 152),
(8,  'Batman Begins',       2005, 140),
(9,  'The Prestige',        2006, 130),
(10, 'Gangs of New York',   2002, 167),
(11, 'Catch Me If You Can', 2002, 141),
(12, 'The Departed',        2006, 151),
(13, 'Blood Diamond',       2006, 143),

-- 2010s (7 Movies)
(14, 'Inception',           2010, 148),
(15, 'Interstellar',        2014, 169),
(16, 'The Dark Knight Rises', 2012, 165),
(17, 'Mad Max: Fury Road',  2015, 120),
(18, 'The Revenant',        2015, 156),
(19, 'Dunkirk',             2017, 106),
(20, 'Once Upon a Time in Hollywood', 2019, 161);


INSERT INTO movie_genres (movie_id, genre_id) VALUES
(1, 1), (1, 2),
(2, 3), (2, 4),
(3, 5), (3, 4), (3, 3),
(4, 1), (4, 5), (4, 4),
(5, 3), (5, 4),
(6, 1), (6, 2),
(7, 1), (7, 5), (7, 4),
(8, 1), (8, 5),
(9, 3), (9, 2), (9, 4),
(11, 5), (11, 3),
(12, 5), (12, 3), (12, 4),
(13, 1), (13, 3), (13, 4),
(14, 1), (14, 2), (14, 4),
(15, 2), (15, 3),
(16, 1), (16, 4),
(17, 1), (17, 2),
(18, 3), (18, 1),
(19, 1), (19, 3),
(20, 3), (20, 5);


INSERT INTO movie_actors (movie_id, actor_id) VALUES
(14, 1), (14, 2), (14, 3),
(1, 4), (1, 5), (1, 6),
(6, 4), (6, 5), (6, 6),
(7, 7), (7, 8), (7, 9),
(8, 7), (8, 9),
(16, 7), (16, 3), (16, 13), (16, 9),
(2, 10), (2, 11),
(3, 10), (3, 9),
(4, 14), (4, 15),
(5, 4), (5, 14),
(9, 7),
(10, 1),
(11, 1),
(12, 1),
(13, 1),
(15, 12), (15, 13),
(17, 3),
(18, 1), (18, 3),
(19, 3),
(20, 1), (20, 10);

-- +goose Down

DELETE FROM movie_actors;
DELETE FROM movie_genres;
DELETE FROM movies;
DELETE FROM actors;
DELETE FROM genres;