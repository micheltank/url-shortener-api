DROP KEYSPACE IF EXISTS url_shortener;

CREATE KEYSPACE url_shortener WITH replication = {'class': 'SimpleStrategy', 'replication_factor' : 1};