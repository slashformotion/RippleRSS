-- +goose Up
CREATE TABLE feed (
    id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
    title text DEFAULT "" NOT NULL,
    description text DEFAULT "" NOT NULL,
    link text DEFAULT "",
    feed_link text DEFAULT "" NOT NULL,
    updated text DEFAULT "" NOT NULL,
    updated_parsed text,
    image_id integer UNIQUE,
    guid text NOT NULL,
    FOREIGN KEY (image_id) REFERENCES image(id) ON DELETE
    SET
        NULL
);

CREATE TABLE person (
    id integer PRIMARY KEY AUTOINCREMENT NOT NULL,
    name text NOT NULL,
    email text,
    feed_id integer UNIQUE,
    FOREIGN KEY (feed_id) REFERENCES feed(id) ON DELETE CASCADE
);

CREATE TABLE item (
    id integer PRIMARY KEY AUTOINCREMENT NOT NULL,
    title text DEFAULT "" NOT NULL,
    description text DEFAULT "" NOT NULL,
    link text DEFAULT "",
    feed_link text DEFAULT "" NOT NULL,
    updated text DEFAULT "" NOT NULL,
    updated_parsed text,
    published text DEFAULT "" NOT NULL,
    published_parsed text,
    image_id integer UNIQUE,
    guid text NOT NULL,
    FOREIGN KEY (image_id) REFERENCES image(id) ON DELETE
    SET
        NULL
);

CREATE TABLE image (
    id integer PRIMARY KEY AUTOINCREMENT NOT NULL,
    url text NOT NULL,
    title text,
);

-- +goose Down
DROP TABLE feed;

DROP TABLE person;

DROP TABLE image;

DROP TABLE item;
