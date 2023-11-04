CREATE TABLE if not exists urls (
    "short_url" VARCHAR (10) PRIMARY KEY,
    "url" VARCHAR(2048) NOT NULL
);

select * FROM urls;

INSERT INTO urls (short_url, url) values ("ddfasdwetu", "youtube.com");
INSERT INTO urls (short_url, url) values ("ssfasd2etu", "youtube1.com");
INSERT INTO urls (short_url, url) values ("ddfasdsetu", "youtube2.com");


select * FROM urls;

INSERT INTO urls (short_url, url)
VALUES (
    'ddfassdetu',
    'youtube.com'
  );

  SELECT url FROM urls WHERE short_url='ddfassdetu'