CREATE TABLE IF NOT EXISTS users (
     id INTEGER PRIMARY KEY AUTO_INCREMENT,
     first_name varchar(255) NOT NULL,
     last_name varchar(255) NOT NULL,
     email varchar(255) NOT NULL,
     password varchar(255) NOT NULL,
     access_level int,
     created_at datetime,
     updated_at datetime
);

CREATE TABLE IF NOT EXISTS records (
     record_id INTEGER PRIMARY KEY AUTO_INCREMENT,
     title varchar(255) NOT NULL,
     label varchar(255) NOT NULL,
     year varchar(255) NOT NULL,
     cover varchar(255) NOT NULL,
     created_at datetime,
     updated_at datetime
);

CREATE TABLE IF NOT EXISTS genres (
    genre_id int PRIMARY KEY AUTO_INCREMENT,
    genre_name varchar(255),
    created_at DATETIME,
    updated_at DATETIME
);


CREATE TABLE IF NOT EXISTS artists (
     artist_id int PRIMARY KEY AUTO_INCREMENT,
     first_name varchar(255),
     last_name varchar(255),
     created_at datetime,
     updated_at datetime
);


CREATE TABLE IF NOT EXISTS record_genres (
   id INTEGER PRIMARY KEY AUTO_INCREMENT,
   record_id INTEGER,
   genre_id INTEGER,
   created_at datetime,
   updated_at datetime,
   FOREIGN KEY (record_id) REFERENCES records(record_id) ON DELETE CASCADE,
   FOREIGN KEY (genre_id) REFERENCES genres(genre_id) ON DELETE CASCADE

);


CREATE TABLE IF NOT EXISTS record_artists (
    id int PRIMARY KEY AUTO_INCREMENT,
    artist_id int,
    record_id int,
    created_at datetime,
    updated_at datetime,
    FOREIGN KEY (artist_id) REFERENCES artists(artist_id) ON DELETE CASCADE,
    FOREIGN KEY (record_id) REFERENCES records(record_id) ON DELETE CASCADE

);



