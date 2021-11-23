alter table record_genres drop foreign key record_genres_ibfk_1;
alter table record_genres drop foreign key record_genres_ibfk_2;
alter table record_artists drop foreign key record_artists_ibfk_1;
alter table record_artists drop foreign key record_artists_ibfk_2;

DROP TABLE IF EXISTS records;
DROP TABLE IF EXISTS genres;
DROP TABLE IF EXISTS artists;
DROP TABLE IF EXISTS record_artists;
DROP TABLE IF EXISTS record_genres;



