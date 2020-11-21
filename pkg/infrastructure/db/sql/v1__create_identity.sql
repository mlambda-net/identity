
CREATE SEQUENCE user_id START 1;

CREATE TABLE identities (
    id int NOT NULL DEFAULT nextval('user_id'),
    name varchar(100) not null,
    last_name varchar(100) not null,
    password varchar(255) not null,
    email varchar(255) not null UNIQUE

);

CREATE INDEX identities_email ON identities(email);

CREATE INDEX identities_name ON identities (name, last_name);