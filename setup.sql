DROP TABLE IF EXISTS channels;

CREATE TABLE channels(
    id INTEGER PRIMARY KEY,
    title VARCHAR(100) NOT NULL
);

DROP TABLE IF EXISTS articles;

CREATE TABLE articles(
    id INTEGER PRIMARY KEY,
    channel_id INT NOT NULL,
    url VARCHAR NOT NULL,
    wc INT NOT NULL,
    FOREIGN KEY(channel_id) REFERENCES channels (id)
);