ALTER TABLE record_genres DROP FOREIGN KEY record_genres_ibfk_1;
ALTER TABLE record_genres DROP FOREIGN KEY record_genres_ibfk_2;
ALTER TABLE record_artists DROP FOREIGN KEY record_artists_ibfk_1;
ALTER TABLE record_artists DROP FOREIGN KEY record_artists_ibfk_2;

DROP TABLE IF EXISTS artists;
DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS records;
DROP TABLE IF EXISTS genres;
DROP TABLE IF EXISTS record_artists;
DROP TABLE IF EXISTS record_genres;



