DROP TABLE IF EXISTS url_shortener.urls;

CREATE TABLE url_shortener.urls (
    hash text PRIMARY KEY,
    long_url text
);