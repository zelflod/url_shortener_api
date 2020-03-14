DROP TABLE IF EXISTS links;

CREATE TABLE IF NOT EXISTS links
(
    id        bigserial                    not null primary key,
    url       varchar(200) COLLATE "POSIX" not null,
    short_url varchar(50) COLLATE "POSIX",
    created   timestamp DEFAULT now(),
    expires   timestamp DEFAULT now() + INTERVAL '1 minute'
);