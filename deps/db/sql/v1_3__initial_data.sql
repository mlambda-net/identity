
INSERT INTO app(id, name, description)
VALUES (gen_random_uuid(), 'Identity', 'The security application.');

INSERT INTO roles( id, app, name, description)
VALUES (
           gen_random_uuid(),
           (SELECT id FROM app WHERE name = 'Identity'),
           'Admin',
           'The admin role for identity'
       );

INSERT INTO roles( id, app, name, description)
VALUES (
           gen_random_uuid(),
           (SELECT id FROM app WHERE name = 'Identity'),
           'Read',
           'The read role for identity'
       );
