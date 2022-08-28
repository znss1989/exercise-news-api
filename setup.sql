DROP TABLE IF EXISTS channels;

CREATE TABLE channels(
    id INT AUTO_INCREMENT NOT NULL,
    channel_name VARCHAR(100) NOT NULL,
    PRIMARY KEY(id)
);

DROP TABLE IF EXISTS articles;

CREATE TABLE articles(
    id INT AUTO_INCREMENT NOT NULL,
    channel_id INT NOT NULL,
    url VARCHAR NOT NULL,
    wc INT NOT NULL,
    PRIMARY KEY(id)
    FOREIGN KEY(channel_id) REFERENCES channels
);