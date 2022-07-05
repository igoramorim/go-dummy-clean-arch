DROP TABLE IF EXISTS artist;
CREATE TABLE artist (
                       id       INT AUTO_INCREMENT NOT NULL,
                       name     VARCHAR(255) NOT NULL,
                       PRIMARY KEY (`id`)
);

INSERT INTO artist
    (name)
VALUES
    ('Led Zeppelin'),
    ('Pink Floyd');

COMMIT;