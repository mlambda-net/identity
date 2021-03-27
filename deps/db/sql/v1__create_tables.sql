
CREATE TABLE identities (
    id uuid NOT NULL,
    name varchar(100) not null,
    last_name varchar(100) not null,
    password varchar(255) not null,
    email varchar(255) not null UNIQUE
);

ALTER TABLE identities
ADD CONSTRAINT PK_identities
PRIMARY KEY (id);

CREATE INDEX identities_email ON identities(email);

CREATE INDEX identities_name ON identities (name, last_name);

CREATE TABLE App
(
    id uuid NOT NULL,
    name varchar(100),
    description varchar(500)
);

ALTER TABLE App
ADD CONSTRAINT PK_Applications
PRIMARY KEY (id);

CREATE UNIQUE INDEX app_unique_name
ON app(name);


CREATE TABLE Roles
(
    id uuid NOT NULL,
    app uuid NOT NULL,
    name varchar(100),
    description varchar(200)
);

ALTER TABLE Roles
ADD CONSTRAINT PK_Roles
PRIMARY KEY (id);

ALTER TABLE Roles
ADD CONSTRAINT FK_Roles_App
FOREIGN KEY (app)
REFERENCES App(id)
ON DELETE CASCADE;

CREATE TABLE Rights
(
    UserId  uuid NOT NULL,
    RoleId uuid NOT NULL
);

ALTER TABLE Rights
ADD CONSTRAINT FK_Identity_Rights
FOREIGN KEY (UserId)
REFERENCES identities(id);


ALTER TABLE Rights
ADD CONSTRAINT FK_Roles_Rights
FOREIGN KEY (RoleId)
REFERENCES Roles(id);

