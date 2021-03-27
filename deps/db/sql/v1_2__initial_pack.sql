
CREATE PROCEDURE USER_ADD
(
    _id uuid,
    _name varchar (100),
    _last_name varchar(100),
    _email varchar(255),
    _password varchar(255)
) AS $$

    INSERT INTO identities(id, name, last_name, email, password)
    VALUES (_id, _name, _last_name, _email, _password);

$$ LANGUAGE SQL;


CREATE PROCEDURE USER_UPDATE
(
    _id uuid,
    _name varchar (100),
    _last_name varchar(100)
) AS $$

    UPDATE identities SET name = _name, last_name = _last_name
    WHERE id = _id;

$$ LANGUAGE SQL;

CREATE PROCEDURE USER_UPDATE_PASSWORD
(
    _id uuid,
    _password varchar(255)
) AS $$

    UPDATE identities SET password = _password
    WHERE id = _id;

$$ LANGUAGE SQL;



CREATE PROCEDURE RIGHTS_ADD
(
    _user_id uuid,
    _role_id uuid
) AS $$

    INSERT INTO rights(userid, roleid)
    VALUES (_user_id, _role_id);

$$ LANGUAGE SQL;

CREATE PROCEDURE RIGHTS_PURGE
(
    _user_id uuid
) AS $$

    DELETE FROM rights
    WHERE userid = _user_id;

$$ LANGUAGE SQL;



CREATE PROCEDURE APP_ADD
(
    _id uuid,
    _name varchar (100),
    _description varchar(500)
) AS $$
    INSERT INTO App(id, name, description)
    VALUES (_id, _name, _description)
$$ LANGUAGE SQL;

CREATE PROCEDURE APP_UPDATE
(
    _id uuid,
    _description varchar(500)
) AS $$

    UPDATE App
    SET description = _description
    WHERE id = _id

$$ LANGUAGE SQL;


CREATE PROCEDURE ROLE_ADD
(
    _id uuid,
    _app_id uuid,
    _name varchar (100),
    _description varchar(500)
) AS $$

    INSERT INTO Roles(id, app, name, description)
    VALUES (_id, _app_id, _name, _description)

$$ LANGUAGE SQL;


CREATE PROCEDURE ROLE_UPDATE
(
    _id uuid,
    _description varchar(500)
) AS $$

    UPDATE Roles SET description = _description
    WHERE id = _id

$$ LANGUAGE SQL;


CREATE PROCEDURE ROLE_DELETE
(
    _id uuid
) AS $$

    DELETE FROM Roles WHERE id = _id

$$ LANGUAGE SQL;
