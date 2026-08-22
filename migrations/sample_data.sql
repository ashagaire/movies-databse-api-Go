-- Insert Genres (Minimum 5)
INSERT INTO genres (id, name) VALUES
(1, 'Action'), (2, 'Sci-Fi'), (3, 'Drama'), (4, 'Thriller'), (5, 'Comedy'), (6, 'Romance');

-- Insert Actors (Minimum 15)
INSERT INTO actors (id, name, birth_date) VALUES
(1, 'Keanu Reeves', '1964-09-02'), (2, 'Carrie-Anne Moss', '1967-08-21'),
(3, 'Tom Hanks', '1956-07-09'), (4, 'Leonardo DiCaprio', '1974-11-11'),
(5, 'Joseph Gordon-Levitt', '1981-02-17'), (6, 'Elliot Page', '1987-02-21'),
(7, 'Christian Bale', '1974-01-30'), (8, 'Morgan Freeman', '1937-06-01'),
(9, 'Heath Ledger', '1979-04-04'), (10, 'Michael Caine', '1933-03-14'),
(11, 'Matthew McConaughey', '1969-11-04'), (12, 'Anne Hathaway', '1982-11-12'),
(13, 'Jessica Chastain', '1977-03-24'), (14, 'Matt Damon', '1970-10-08'),
(15, 'Brad Pitt', '1963-12-18'), (16, 'Edward Norton', '1969-08-18');

-- Insert Movies (Minimum 20, spanning multiple decades)
INSERT INTO movies (id, title, release_year, duration) VALUES
(1, 'The Matrix', 1999, 136), (2, 'The Matrix Reloaded', 2003, 138),
(3, 'Forrest Gump', 1994, 142), (4, 'Cast Away', 2000, 143),
(5, 'Inception', 2010, 148), (6, 'The Dark Knight', 2008, 152),
(7, 'Batman Begins', 2005, 140), (8, 'Interstellar', 2014, 169),
(9, 'The Martian', 2015, 144), (10, 'Fight Club', 1999, 139),
(11, 'Se7en', 1995, 127), (12, 'Ocean''s Eleven', 2001, 116),
(13, 'The Truman Show', 1998, 103), (14, 'Saving Private Ryan', 1998, 169),
(15, 'The Green Mile', 1999, 189), (16, 'Catch Me If You Can', 2002, 141),
(17, 'Apollo 13', 1995, 140), (18, 'The Departed', 2006, 151),
(19, 'Gladiator', 2000, 155), (20, 'Jurassic Park', 1993, 127);

-- Insert Movie-Genre Relationships
INSERT INTO movie_genres (movie_id, genre_id) VALUES
(1, 1), (1, 2), (2, 1), (2, 2), (3, 3), (3, 6), (4, 3), (5, 1), (5, 2), (5, 4),
(6, 1), (6, 3), (6, 4), (7, 1), (7, 3), (8, 2), (8, 3), (9, 2), (9, 3), (10, 3),
(11, 3), (11, 4), (12, 4), (12, 5), (13, 3), (13, 5), (14, 1), (14, 3), (15, 3),
(16, 3), (16, 5), (17, 3), (18, 3), (18, 4), (19, 1), (19, 3), (20, 1), (20, 2);

-- Insert Movie-Actor Relationships
INSERT INTO movie_actors (movie_id, actor_id) VALUES
(1, 1), (1, 2), (2, 1), (2, 2), (3, 3), (4, 3), (5, 4), (5, 5), (5, 6), (5, 10),
(6, 7), (6, 8), (6, 9), (6, 10), (7, 7), (7, 8), (7, 10), (8, 11), (8, 12), (8, 13),
(8, 10), (9, 14), (9, 13), (10, 15), (10, 16), (11, 15), (11, 8), (12, 15), (12, 14),
(14, 3), (14, 14), (15, 3), (16, 3), (16, 4), (17, 3), (18, 4), (18, 14);
