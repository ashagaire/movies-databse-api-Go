package internal


 const schema = `CREATE TABLE IF NOT EXISTS genres ( 
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		name TEXT
		);
		
		CREATE TABLE IF NOT EXISTS actors (
    	id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
    	birth_date TEXT
		);

		CREATE TABLE IF NOT EXISTS movies (
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    	title TEXT NOT NULL,
    	release_year TEXT NOT NULL,
		duration INTEGER NOT NULL
		);

		CREATE TABLE IF NOT EXISTS movie_genres (
		movie_id INTEGER NOT NULL,
		genre_id INTEGER NOT NULL,
		PRIMARY KEY (movie_id, genre_id),
		FOREIGN KEY (movie_id) REFERENCES movies(id) ON DELETE CASCADE,
		FOREIGN KEY (genre_id) REFERENCES genres(id) ON DELETE CASCADE
		);

		CREATE TABLE IF NOT EXISTS movie_actors (
		movie_id INTEGER NOT NULL,
		actor_id INTEGER NOT NULL,
		PRIMARY KEY (movie_id, actor_id),
		FOREIGN KEY (movie_id) REFERENCES movies(id) ON DELETE CASCADE,
		FOREIGN KEY (actor_id) REFERENCES actors(id) ON DELETE CASCADE
		);
		`

// for id to be immutable once set AUTOINCREMENT is used
