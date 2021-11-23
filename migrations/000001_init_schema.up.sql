CREATE TABLE records (
     record_id INTEGER PRIMARY KEY AUTO_INCREMENT,
     title varchar(255) NOT NULL,
     label varchar(255) NOT NULL,
     year varchar(255) NOT NULL,
     created_at datetime,
     updated_at datetime
);

CREATE TABLE genres (
    genre_id int PRIMARY KEY AUTO_INCREMENT,
    genre_name varchar(255),
    created_at DATETIME,
    updated_at DATETIME
);


CREATE TABLE artists (
     artist_id int PRIMARY KEY AUTO_INCREMENT,
     first_name varchar(255),
     last_name varchar(255),
     created_at datetime,
     updated_at datetime
);


CREATE TABLE record_genres (
   id INTEGER PRIMARY KEY AUTO_INCREMENT,
   record_id INTEGER,
   genre_id INTEGER,
   created_at datetime,
   updated_at datetime,
   FOREIGN KEY (record_id) REFERENCES records(record_id) ON DELETE CASCADE,
   FOREIGN KEY (genre_id) REFERENCES genres(genre_id) ON DELETE CASCADE

);


CREATE TABLE record_artists (
    id int PRIMARY KEY AUTO_INCREMENT,
    artist_id int,
    record_id int,
    created_at datetime,
    updated_at datetime,
    FOREIGN KEY (artist_id) REFERENCES artists(artist_id) ON DELETE CASCADE,
    FOREIGN KEY (record_id) REFERENCES records(record_id) ON DELETE CASCADE

);
